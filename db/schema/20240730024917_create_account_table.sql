-- +goose Up
-- +goose StatementBegin
BEGIN;
CREATE EXTENSION IF NOT EXISTS "pgcrypto";
CREATE TABLE accounts (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    status varchar NOT NULL CHECK (status IN ('ACTIVE', 'INACTIVE')),
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
BEGIN;
DROP EXTENSION IF EXISTS "pgcrypto";
DROP TABLE IF EXISTS accounts;
COMMIT;
-- +goose StatementEnd
