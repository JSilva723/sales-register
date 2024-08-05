-- +goose Up
-- +goose StatementBegin
BEGIN;
CREATE TABLE accounts (
    id int PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    status varchar NOT NULL DEFAULT 'ACTIVE',
    CONSTRAINT validate_status CHECK(status IN ('ACTIVE', 'INACTIVE'))
);
CREATE TABLE users (
    id serial PRIMARY KEY,
    account_name varchar NOT NULL,
    username varchar NOT NULL,
    password varchar NOT NULL,
    rol varchar NOT NULL CHECK (rol IN ('ADMIN', 'EMPLOYEE')),
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (account_name) REFERENCES accounts (name)
);
CREATE TABLE products (
    id serial PRIMARY KEY,
    account_name varchar NOT NULL,
    name varchar NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (account_name) REFERENCES accounts (name)
);
CREATE TABLE payments (
    id serial PRIMARY KEY,
    account_name varchar NOT NULL,
    name varchar NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (account_name) REFERENCES accounts (name)
);
CREATE TABLE sale_orders (
    id serial PRIMARY KEY,
    account_name varchar NOT NULL,
    user_id int NOT NULL,
    payment_id int NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (account_name) REFERENCES accounts (name),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (payment_id) REFERENCES payments (id)
);
CREATE TABLE order_lines (
    id bigserial PRIMARY KEY,
    account_name varchar NOT NULL,
    sale_order_id int NOT NULL,
    ammount int NOT NULL,
    product_id int NOT NULL,
    FOREIGN KEY (account_name) REFERENCES accounts (name),
    FOREIGN KEY (sale_order_id) REFERENCES sale_orders (id),
    FOREIGN KEY (product_id) REFERENCES products (id)
);
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;
DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS sale_orders;
DROP TABLE IF EXISTS order_lines;
COMMIT;
-- +goose StatementEnd
