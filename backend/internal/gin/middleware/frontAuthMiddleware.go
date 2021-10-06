package middleware

import (
	"fmt"
	"strings"

	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/SubGame-Network/SubGameModuleService/internal/gin/handler"
	SubGameHelp "github.com/SubGame-Network/SubGameModuleService/internal/md/service"
	"github.com/SubGame-Network/SubGameModuleService/internal/provider"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	frontExceptAuthUri = []string{}
)

func FrontAuthMiddleware(c *gin.Context) {
	address := c.GetHeader("Address")

	if address == "" {
		handler.Failed(c, domain.ErrorNotSubGameAddress, domain.ErrorNotSubGameAddress.Message)
		c.Abort()
		return
	}

	pubkey, err := SubGameHelp.SubGameAddressToPubkey(address)
	if err != nil {
		handler.Failed(c, domain.ErrorNotSubGameAddress, domain.ErrorNotSubGameAddress.Message)
		c.Abort()
		return
	}
	pubkey = strings.TrimPrefix(pubkey, "0x")

	c.Set("userAddress", pubkey)
	c.Next()
}

func checkFrontToken(ctx *gin.Context) bool {
	authHeader := ctx.GetHeader("Authorization")
	tokenString, err := FromFrontAuthHeader(authHeader)
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

	isAdmin := false
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

	ctx.SetCookie(domain.CookieNameFront, newToken, 60*60*24, "/", "", false, false)

	ctx.Set("userAccount", claims["account"])
	return true
}

func FromFrontAuthHeader(authHeader string) (string, error) {
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
