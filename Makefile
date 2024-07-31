ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

run: build
	@./bin/a

build: fmt
	@go build -o ./bin/a ./main.go

goose-create: ## Create migrarion file .Args name="create_user_table"
	@goose -dir ./db/schema create ${name} sql

goose-up: ## Up migration
	@goose -dir ./db/schema postgres ${DATABASE_URL} up

goose-down: ## Down migration
	@goose -dir ./db/schema postgres ${DATABASE_URL} down

sqlc-generate: ## Generate db functions
	@sqlc generate

db-test-up:
	@docker compose up db-test -d

db-test-down:
	@docker compose down db-test

db-test-ssh:
	@docker exec -it db-sales-register-test bash

test:
	@go test -v -cover ./...