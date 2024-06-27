#!/bin/sh

set -a
source .env
set +a

goose -dir /rates/internal/database/migrations postgres "host=$DB_HOST port=$DB_PORT user=$DB_USER password=$DB_PASSWORD dbname=$DB_NAME sslmode=disable" up

./cmd/main