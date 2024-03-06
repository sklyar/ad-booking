DOCKER_COMPOSE_TEST_FILE = deployments/test/docker-compose.yaml
DOCKER_COMPOSE_DEV_FILE = deployments/dev/docker-compose.yaml
DEV_DB_URL = "postgres://dev_user:dev_password@localhost:16000/booking?sslmode=disable"
MIGRATION_DIR = migrations

generate:
	$(call print-target)
	@buf generate

lint:
	$(call print-target)
	@buf lint
	@golangci-lint run ./...

test: test-env-up
	$(call print-target)
	@go test -race ./...

test-env-up: test-env-down
	@docker-compose -f $(DOCKER_COMPOSE_TEST_FILE) up -d

test-env-down:
	@docker-compose -f $(DOCKER_COMPOSE_TEST_FILE) down

dev-env-up: dev-env-down
	@docker-compose -f $(DOCKER_COMPOSE_DEV_FILE) up -d

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
