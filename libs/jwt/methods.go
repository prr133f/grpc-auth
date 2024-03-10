package jwt

import (
	"auth/utils"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

var logger = utils.InitZap()

func Generate(payload Payload) (string, error) {
	if err := payload.Valid(); err != nil {
		logger.Error(err.Error())
		return "", err
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":        payload.ID,
		"sub":       payload.Email,
		"exp":       payload.ExpiresAt,
		"role":      payload.Role,
		"contactId": payload.ContactId,
	}).SignedString(key)
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return token, nil
}

func Verify(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		logger.Error(err.Error())
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ReissueAccessToken(refresh string) (string, error) {
	if err := Verify(refresh); err != nil {
		logger.Error(err.Error())
		return "", err
	}

	token, err := jwt.Parse(refresh, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		logger.Error(err.Error())
		return "", err
	}

	return Generate(Payload{
		ID:        int64(token.Claims.(jwt.MapClaims)["id"].(float64)),
		Email:     token.Claims.(jwt.MapClaims)["sub"].(string),
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Role:      token.Claims.(jwt.MapClaims)["role"].(string),
		ContactId: token.Claims.(jwt.MapClaims)["contactId"].(string),
	})
}

func GetPayload(token string) (jwt.MapClaims, error) {
	payload, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return payload.Claims.(jwt.MapClaims), nil
}
