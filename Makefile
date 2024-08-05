ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

run: build
	@./bin/a

build:
	@go build -o ./bin/a ./main.go

db-up:
	@docker compose up db -d

db-down:
	@docker compose down db

db-ssh:
	@docker exec -it db-sales-register bash

goose-create:
	@goose -dir ./db/schema create ${name} sql

goose-up:
	@goose -dir ./db/schema postgres ${DATABASE_URL} up

goose-down:
	@goose -dir ./db/schema postgres ${DATABASE_URL} down

sqlc-generate:
	@sqlc generate

test:
	@bash script/test.sh