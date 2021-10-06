package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/SubGame-Network/SubGameModuleService/config"
	"github.com/SubGame-Network/SubGameModuleService/domain"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

type jwtUtil struct {
	config *config.Config
}

func NewJwt(config *config.Config) domain.JwtUtil {
	return &jwtUtil{
		config: config,
	}
}

type claims struct {
	Account string `json:"account"`
	UUID    string `json:"UUID"`
	jwt.StandardClaims
}

const adminAccessTokenKeyPrefix = "SubgameModuleAdmin:AccessToken"
const frontAccessTokenKeyPrefix = "SubgameModuleFront:AccessToken"

func (t *jwtUtil) GetAccessTokenKey(account string, isAdmin bool) string {
	var accessTokenKeyPrefix string
	if isAdmin {
		accessTokenKeyPrefix = adminAccessTokenKeyPrefix
	} else {
		accessTokenKeyPrefix = frontAccessTokenKeyPrefix
	}
	return fmt.Sprintf("%s:%s", accessTokenKeyPrefix, account)
}

type basicInfo struct {
	Account string
	TokenID string
}

func (t *jwtUtil) GenToken(params map[string]string) (string, string, error) {
	config := t.config

	jti := uuid.NewV4().String()
	now := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Account: params["account"],
		UUID:    params["UUID"],
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Jwt.Issuer, //頒發者，是區分大小寫的字串，可以是一個字串或是網址
			IssuedAt:  now,               //頒發時間，是數字日期
			NotBefore: now,               //定義在什麼時間之前，不可用，是數字日期(目前沒用到)
			// ExpiresAt: now + int64(config.Jwt.access_token_exp_sec.(int)), //過期時間，是數字日期(目前沒用到)
			Id: jti, //唯一識別碼，是區分大小寫的字串
		},
	})

	result, err := token.SignedString([]byte(config.Jwt.Secret))
	if err != nil {
		return "", "", err
	}

	info := &basicInfo{Account: params["account"], TokenID: jti}
	infoJsonStr, _ := json.Marshal(info)

	return result, string(infoJsonStr), nil
}

// Parse validates token gotten from request and returns claims if it's legal
func (t *jwtUtil) Parse(tokenString string) (map[string]string, error) {
	config := t.config

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Jwt.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, err := validate(config, token)

	if err != nil {
		return nil, err
	}
	return claims, err
}

func validate(config *config.Config, token *jwt.Token) (map[string]string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok {
		if claims["iss"] != config.Jwt.Issuer {
			return nil, errors.New("token issuer mismatch")
		}
	} else {
		return nil, errors.New("get token claims failed")
	}

	return map[string]string{
		"account": claims["account"].(string),
		"UUID":    claims["UUID"].(string),
		"jti":     claims["jti"].(string),
		"iss":     claims["iss"].(string),
	}, nil
}

func (t *jwtUtil) RenewToken(claim map[string]string) (string, error) {
	config := t.config
	jti := claim["jti"]
	now := time.Now().Unix() - 60
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		Account: claim["account"],
		UUID:    claim["UUID"],
		StandardClaims: jwt.StandardClaims{
			Issuer:    claim["iss"],                                 //頒發者，是區分大小寫的字串，可以是一個字串或是網址
			IssuedAt:  now,                                          //頒發時間，是數字日期
			NotBefore: now,                                          //定義在什麼時間之前，不可用，是數字日期(目前沒用到)
			ExpiresAt: now + int64(config.Jwt.Access_token_exp_sec), //過期時間，是數字日期(目前沒用到)
			Id:        jti,                                          //唯一識別碼，是區分大小寫的字串
		},
	})

	result, err2 := token.SignedString([]byte(config.Jwt.Secret))
	if err2 != nil {
		return "", err2
	}

	return result, nil
}

func (t *jwtUtil) CheckTokenJti(claim map[string]string, infoJsonStr string) bool {
	var bi basicInfo
	json.Unmarshal([]byte(infoJsonStr), &bi)
	if bi.TokenID != claim["jti"] {
		return false
	}
	return true
}
