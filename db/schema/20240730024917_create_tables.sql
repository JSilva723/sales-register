-- +goose Up
-- +goose StatementBegin
BEGIN;
CREATE TABLE accounts (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    status varchar NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE'))
);
CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar NOT NULL,
    password varchar NOT NULL,
    rol varchar NOT NULL CHECK (rol IN ('ADMIN', 'EMPLOYEE')),
    account_name varchar NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE users ADD FOREIGN KEY (account_name) REFERENCES accounts (name);
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;
DROP TABLE IF EXISTS accounts;
COMMIT;
-- +goose StatementEnd
