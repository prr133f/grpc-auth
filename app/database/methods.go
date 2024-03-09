package database

import (
	"auth/app/models"
	"context"

	"github.com/jackc/pgx/v5"
)

func (p *Postgres) GetUserByEmail(email string) (models.User, error) {
	row, err := p.DB.Query(context.Background(), `
	SELECT id, email, pwdhash, role
	FROM users_schema.user
	WHERE email=$1`, email)
	if err != nil {
		p.Log.Error().Err(err).Stack()
		return models.User{}, err
	}

	user, err := pgx.RowToStructByName[models.User](row)
	if err != nil {
		p.Log.Error().Err(err).Stack()
		return models.User{}, err
	}

	return user, nil
}
