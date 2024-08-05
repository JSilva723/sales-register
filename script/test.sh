#!/bin/bash

CONTAINER="db-sales-register-test"
CONTAINER_PORT="5432"
DB_SERVICE="db-test"
DB_NAME="salesregister"
DB_USER="salesregister"
DB_PASSWORD="salesregister"
DB_HOST="localhost"
DB_PORT="5433"
DB_URL="postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable"
MIGRATIONS_DIR="$(pwd)/db/schema"

wait_for_container() {
    local retries=10

    until docker inspect -f {{.State.Running}} $CONTAINER | grep true > /dev/null; do
        sleep 1
        ((retries--))

        if [ $retries -le 0 ]; then
            echo "Error: $CONTAINER no está listo a tiempo."
            exit 1
        fi
    done
}

wait_for_db() {
    local retries=10

    until docker exec $CONTAINER pg_isready -U $DB_USER -d $DB_NAME -h $DB_HOST -p $CONTAINER_PORT; do
        sleep 1
        ((retries--))

        if [ $retries -le 0 ]; then
            echo "Error: La base de datos no está lista a tiempo."
            exit 1
        fi
    done
}

docker compose up $DB_SERVICE -d
wait_for_container $CONTAINER
wait_for_db

if ! goose -dir "$MIGRATIONS_DIR" postgres "$DB_URL" up; then
    echo "Error: Falló la ejecución de las migraciones."
    docker compose down $DB_SERVICE
    exit 1
fi

if ! go test -v -cover ./...; then
    echo "Error: Fallaron las pruebas."
    docker compose down $DB_SERVICE
    exit 1
fi

docker compose down $DB_SERVICE