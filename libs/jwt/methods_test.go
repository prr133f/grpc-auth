package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	payload := Payload{
		ID:        1,
		Email:     "test@gmail.com",
		ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		Role:      "admin",
		ContactId: "550e8400-e29b-41d4-a716-446655440000",
	}

	token, err := Generate(payload)
	assert.Nil(t, err)

	access, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return key, err
	})
	assert.Nil(t, err)

	assert.Equal(t, 1, int(access.Claims.(jwt.MapClaims)["id"].(float64)))
	assert.Equal(t, "admin", access.Claims.(jwt.MapClaims)["role"].(string))
}

func TestGenerateWrongPayload(t *testing.T) {
	payloads := []Payload{{
		ID:        1,
		Email:     "arraabba",
		ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		Role:      "admin",
	}, {
		ID:        1,
		Email:     "test@gmail.com",
		ExpiresAt: time.Now().Add(-time.Minute).Unix(),
		Role:      "admin",
	}, {
		ID:        1,
		Email:     "test@gmail.com",
		ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		Role:      "GOD",
	},
	}

	for _, payload := range payloads {
		_, err := Generate(payload)
		assert.NotNil(t, err, "payload wasn't validate correctly")
	}
}

func TestVerify(t *testing.T) {
	validToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":   1,
		"sub":  "testmail@gmail.com",
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
		"role": "admin",
	}).SignedString(key)
	assert.Nil(t, err)

	assert.Nil(t, Verify(validToken))

	invalidToken, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"id":   1,
		"sub":  "testmail@gmail.com",
		"exp":  time.Now().Add(-time.Minute * 20).Unix(),
		"role": "admin",
	}).SignedString(key)
	assert.Nil(t, err)

	assert.NotNil(t, Verify(invalidToken))
}
