test-cover:
	go test -race ./... -coverprofile res.coverprofile
	go tool cover -html=res.coverprofile

LINK_DB_HOST_MASTER?=localhost
LINK_DB_PORT?=5432
LINK_DB_USER?=postgres
LINK_DB_PASS?=qwerty
LINK_DB_DATABASE_NAME?=postgres


migrate:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir migrations postgres \
		"host=$(LINK_DB_HOST_MASTER)\
		port=$(LINK_DB_PORT)\
		dbname=$(LINK_DB_DATABASE_NAME)\
		user=$(LINK_DB_USER)\
		password=$(LINK_DB_PASS)\
		sslmode=disable" up
