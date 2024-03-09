CREATE TYPE IF NOT EXISTS users_schema.role AS ENUM ('admin', 'user');

CREATE TABLE IF NOT EXISTS users_schema.user(
    id SERIAL,
    email VARCHAR(80),
    pwdhash BINARY,
    role users_schema.role,
    PRIMARY KEY(id),
    UNIQUE(email)
);