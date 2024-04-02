package database

import (
	"auth/app/models"
	"context"

	"go.uber.org/zap"
)

func (p *Postgres) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	var strole string

	if err := p.DB.QueryRow(context.Background(), `
	SELECT id, email, pwdhash, role
	FROM users_schema.user
	WHERE email=$1`, email).Scan(
		&user.ID,
		&user.Email,
		&user.Pwdhash,
		&strole,
	); err != nil {
		p.Log.Error(err.Error(),
			zap.Error(err))
		return models.User{}, err
	}

	user.Role = models.Role(strole)

	return user, nil
}
