DOCKER_COMPOSE_TEST_FILE = deployments/test/docker-compose.yaml
DOCKER_COMPOSE_DEV_FILE = deployments/dev/docker-compose.yaml

lint:
	$(call print-target)
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

define print-target
	@printf "Executing target: \033[36m$@\033[0m\n"
endef
