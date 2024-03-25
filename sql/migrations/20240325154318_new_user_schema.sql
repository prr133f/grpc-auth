-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA users_schema;

CREATE TYPE users_schema.role AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS users_schema.user(
    id SERIAL,
    email VARCHAR(80),
    pwdhash TEXT,
    role users_schema.role,
    PRIMARY KEY(id),
    UNIQUE(email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA users_schema CASCADE;
-- +goose StatementEnd
