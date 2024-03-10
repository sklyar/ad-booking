DOCKER_COMPOSE_TEST_FILE = deployments/test/docker-compose.yaml
DOCKER_COMPOSE_DEV_FILE = deployments/dev/docker-compose.yaml
DEV_DB_URL = "postgres://dev_user:dev_password@localhost:15000/booking?sslmode=disable"
TEST_DB_URL = "postgres://test_user:test_password@localhost:16000/template?sslmode=disable"
MIGRATION_DIR = migrations

generate:
	$(call print-target)
	@buf generate

lint: deps
	$(call print-target)
	@goimports -w .
	@buf lint
	@golangci-lint run ./...

deps:
	$(call print-target)
	@go install golang.org/x/tools/cmd/goimports@latest

test: test-env-up
	$(call print-target)
	@go test -race ./...

test-env-up: test-env-down
	@docker-compose -f $(DOCKER_COMPOSE_TEST_FILE) up -d

	@while ! docker-compose -f $(DOCKER_COMPOSE_TEST_FILE) ps | grep "booking_test_postgres.*healthy" > /dev/null; do \
    sleep 1; \
    done

	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(TEST_DB_URL) goose -dir $(MIGRATION_DIR) up

test-env-down:
	@docker-compose -f $(DOCKER_COMPOSE_TEST_FILE) down

dev-env-up: dev-env-down
	@docker-compose -f $(DOCKER_COMPOSE_DEV_FILE) up -d

	@while ! docker-compose -f $(DOCKER_COMPOSE_DEV_FILE) ps | grep "booking_dev_postgres.*healthy" > /dev/null; do \
    sleep 1; \
    done

	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DEV_DB_URL) goose -dir $(MIGRATION_DIR) up

dev-env-down:
	@docker-compose -f $(DOCKER_COMPOSE_DEV_FILE) down

migration-create:
	$(call print-target)
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DEV_DB_URL) goose -dir $(MIGRATION_DIR) create $(name) sql

migration-up:
	$(call print-target)
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DEV_DB_URL) goose -dir $(MIGRATION_DIR) up

migration-down:
	$(call print-target)
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DEV_DB_URL) goose -dir $(MIGRATION_DIR) down

define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef
