-- +goose Up
-- +goose StatementBegin
BEGIN;
CREATE TABLE accounts (
    id int PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    status varchar NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE'))
);
CREATE TABLE users (
    id serial PRIMARY KEY,
    account_name varchar NOT NULL,
    username varchar NOT NULL,
    password varchar NOT NULL,
    rol varchar NOT NULL CHECK (rol IN ('ADMIN', 'EMPLOYEE')),
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE products (
    id serial PRIMARY KEY,
    account_name varchar NOT NULL,
    name varchar NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
CREATE TABLE payments (
    id serial PRIMARY KEY,
    account_name varchar NOT NULL,
    name varchar NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE users ADD FOREIGN KEY (account_name) REFERENCES accounts (name);
ALTER TABLE products ADD FOREIGN KEY (account_name) REFERENCES accounts (name);
ALTER TABLE payments ADD FOREIGN KEY (account_name) REFERENCES accounts (name);
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS payments;
COMMIT;
-- +goose StatementEnd
