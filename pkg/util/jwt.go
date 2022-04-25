package util

import (
	"github.com/golang-jwt/jwt"
	"github.com/whoismarcode/go-chat-room/global"
	"github.com/whoismarcode/go-chat-room/pkg/logging"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	secretKey := []byte(global.Config.Jwt.SecretKey)
	expiration := time.Now().Add(time.Minute * 10).Unix()

	// Create the Claims
	claims := MyClaims{
		username, //TODO: encode
		password,
		jwt.StandardClaims{
			ExpiresAt: expiration,
			Issuer:    global.Config.Jwt.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secretKey)
	if err != nil {
		logging.Error(err)
		return "", err
	}

	return token, nil
}

func VerifyToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.SecretKey), nil
	})

	if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
