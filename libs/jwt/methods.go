package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte(os.Getenv("JWT_SECRET"))

func Generate(payload Payload) (string, error) {
	if err := payload.Valid(); err != nil {
		log.Error().Err(err).Stack()
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
		log.Error().Err(err).Stack()
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
		log.Error().Msg("invalid token")
		return fmt.Errorf("invalid token")
	}

	return nil
}

func ReissueAccessToken(refresh string) (string, error) {
	if err := Verify(refresh); err != nil {
		log.Error().Err(err).Stack()
		return "", err
	}

	token, err := jwt.Parse(refresh, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Error().Err(err).Stack()
		return "", err
	}

	return Generate(Payload{
		ID:        token.Claims.(jwt.MapClaims)["id"].(int64),
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
		log.Error().Err(err).Stack()
		return nil, err
	}

	return payload.Claims.(jwt.MapClaims), nil
}
