#!/bin/sh
# wait-for-postgres.sh

set -e

host="$1"
shift
cmd="$@"

until PGPASSWORD=$DB_PASSWORD psql -h "$host" -U "postgres" -c '\q'; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

>&2 echo "Postgres is up - executing command"
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir migrations postgres \
		"host=localhost\
		port=5432\
		dbname=postgres\
		user=postgres\
		password=qwerty\
		sslmode=disable" up
exec $cmd