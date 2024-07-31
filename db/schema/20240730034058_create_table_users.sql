-- +goose Up
-- +goose StatementBegin
BEGIN;
CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar NOT NULL,
    password varchar NOT NULL,
    account_id uuid NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE users ADD FOREIGN KEY (account_id) REFERENCES accounts (id); 
COMMIT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
