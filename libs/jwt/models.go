package jwt

import (
	"fmt"
	"net/mail"
	"time"
)

type TokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type Payload struct {
	ID        int64
	Email     string
	ExpiresAt int64
	Role      string
	ContactId string
}

func (p *Payload) Valid() error {
	allowedRoles := []string{"admin", "client", "operator"}

	if _, err := mail.ParseAddress(p.Email); err != nil {
		return err
	}

	if time.Now().Unix() > time.Unix(p.ExpiresAt, 0).Unix() {
		return fmt.Errorf("token expired")
	}

	for _, role := range allowedRoles {
		if role == p.Role {
			return nil
		}
	}
	return fmt.Errorf("unallowed role")
}
