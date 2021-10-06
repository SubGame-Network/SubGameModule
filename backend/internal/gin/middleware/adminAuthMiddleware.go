package middleware

import (
	"net/http"
	"regexp"

	"go.uber.org/zap"

	"fmt"
	"strings"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
)

var (
	exceptAuthUri = []string{
		`/login`,
	}
)

// AdminAuthMiddleware validate token before handler
func AdminAuthMiddleware(ctx *gin.Context) {
	checkExceptAuth := false
	for _, uri := range exceptAuthUri {
		pathRegexp := regexp.MustCompile(uri)
		if pathRegexp.MatchString(ctx.FullPath()) {
			checkExceptAuth = true
		}
	}
	if checkExceptAuth {
		ctx.Next()
		return
	}

	if ok := checkToken(ctx); !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    domain.ErrorAuthTokenExpired.Code,
			"message": domain.ErrorAuthTokenExpired.Message,
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}

func checkToken(ctx *gin.Context) bool {
	authHeader := ctx.GetHeader("Authorization")
	tokenString, err := FromAuthHeader(authHeader)
	if err != nil {
		zap.S().Info(err)
		return false
	}

	config := provider.NewConfig()

	jwtUtil, _ := provider.NewJwt()
	claims, err := jwtUtil.Parse(tokenString)
	if err != nil {
		return false
	}

	redis, err := provider.NewRedis()
	if err != nil {
		zap.S().Warn(err)
		return false
	}

	isAdmin := true
	key := jwtUtil.GetAccessTokenKey(claims["account"], isAdmin)
	infoJsonStr, _ := redis.Get(key)
	if ok := jwtUtil.CheckTokenJti(claims, infoJsonStr); !ok {
		// zap.S().Infof("Redis token 不匹配, claims = %v, infoJsonStr = %v", claims, infoJsonStr)
		return false
	}

	newToken, err := jwtUtil.RenewToken(claims)
	if err != nil {
		zap.S().Warn(err)
		return false
	}

	ttl := config.Jwt.Refresh_token_exp
	if err = redis.SetEx(key, infoJsonStr, ttl); err != nil {
		zap.S().Warn(err)
		return false
	}

	ctx.SetCookie(domain.CookieName, newToken, 60*60*24, "/", "", false, false)

	ctx.Set("AdminAccount", claims["account"])
	ctx.Set("AdminUUID", claims["UUID"])
	return true
}

func FromAuthHeader(authHeader string) (string, error) {
	if authHeader == "" {
		// No error, just no token
		return "", nil
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}
