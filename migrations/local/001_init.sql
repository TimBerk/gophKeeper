-- migrations/001_init.sql
-- +goose Up
CREATE TABLE IF NOT EXISTS secrets(
    id   TEXT primary key,
    type TEXT,
    data BLOB,
    meta TEXT
);
-- +goose Down
DROP TABLE secrets;

