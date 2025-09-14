.PHONY: migration-up migration-down migration-local migration-status migration-generate

migration-up:
	go run github.com/pressly/goose/v3/cmd/goose -dir ./migrations mysql "root:fiber_rest_api@tcp(localhost:3306)/fiber_db" up

migration-down:
	go run github.com/pressly/goose/v3/cmd/goose -dir ./migrations mysql "root:fiber_rest_api@tcp(localhost:3306)/fiber_db" down

# migrate-local:
# 	@export $$(grep -v '^#' .env | grep -v '^$$' | xargs) && make migrate

migration-generate:
	go run github.com/pressly/goose/v3/cmd/goose -dir ./migrations create $(name) sql

