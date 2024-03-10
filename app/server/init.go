package main

import (
	"auth/app/database"
	"auth/app/models"
	"context"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func insertDefaultUser(p *database.Postgres) error {
	pwdhash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("DEFAULTUSER_PWD")), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	defaultUser := models.User{
		Email:   os.Getenv("DEFAULTUSER_EMAIL"),
		Pwdhash: string(pwdhash),
		Role:    models.Role(os.Getenv("DEFAULTUSER_ROLE")),
	}
	if _, err := p.DB.Exec(context.Background(), `
	INSERT INTO users_schema.user(email, pwdhash, role)
	VALUES ($1, $2, $3)
	ON CONFLICT (email)
	DO NOTHING`, defaultUser.Email, defaultUser.Pwdhash, defaultUser.Role); err != nil {
		return err
	}

	return nil
}
