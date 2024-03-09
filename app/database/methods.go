package database

import (
	"auth/app/models"
	"context"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func (p *Postgres) GetUserByEmail(email string) (models.User, error) {
	row, err := p.DB.Query(context.Background(), `
	SELECT id, email, pwdhash, role
	FROM users_schema.user
	WHERE email=$1`, email)
	if err != nil {
		p.Log.Error("error while selecting user",
			zap.Error(err))
		return models.User{}, err
	}

	user, err := pgx.CollectOneRow[models.User](row, pgx.RowToStructByName[models.User])
	if err != nil {
		p.Log.Error("error while selecting user",
			zap.Error(err))
		return models.User{}, err
	}

	return user, nil
}
