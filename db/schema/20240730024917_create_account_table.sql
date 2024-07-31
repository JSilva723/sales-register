-- +goose Up
-- +goose StatementBegin
CREATE TABLE accounts (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    status varchar NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
