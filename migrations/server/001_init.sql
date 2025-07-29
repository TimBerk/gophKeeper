-- migrations/001_init.sql
-- +goose Up
CREATE TABLE IF NOT EXISTS users (
     id       text PRIMARY KEY,
     username text UNIQUE NOT NULL,
     hash     bytea NOT NULL
);

CREATE TABLE IF NOT EXISTS secrets (
    id      text PRIMARY KEY,
    user_id text NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    type    text NOT NULL,
    data    bytea NOT NULL,
    meta    jsonb
);

-- +goose Down
DROP TABLE IF EXISTS secrets;
DROP TABLE IF EXISTS users;
