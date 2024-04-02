package database

import (
	"regexp"
	"testing"

	"github.com/pashagolub/pgxmock/v3"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestPostgres_GetUserByEmail(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)

	email := "test@email.com"

	mock.ExpectQuery(
		regexp.QuoteMeta(`
		SELECT id, email, pwdhash, role
		FROM users_schema.user
		WHERE email=$1`),
	).WithArgs(email).WillReturnRows(pgxmock.NewRows([]string{"id", "email", "pwdhash", "role"}).AddRow(1, email, "somehash", "admin"))

	logger := zap.Must(zap.NewDevelopment())

	pgMock := Postgres{
		Log: logger,
		DB:  mock,
	}

	_, err = pgMock.GetUserByEmail(email)
	require.NoError(t, err)
}
