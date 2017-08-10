# Project variables
PROJECT_NAME ?= phrasebook-api
ORG_NAME ?= phrasebooklearner
REPO_NAME ?= phrasebook-api

# Docker Compose Project Names
REL_PROJECT := $(PROJECT_NAME)$(BUILD_ID)
DEV_PROJECT := $(REL_PROJECT)dev
TE := $

# Filenames
DEV_COMPOSE_FILE := docker/dev/docker-compose.yml
REL_COMPOSE_FILE := docker/release/docker-compose.yml
DOCKER_DEV_COMPOSE_FILE := docker-dev/docker-compose.yml

.PHONY: test
test:
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) build
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up agent
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up migrate
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up test

.PHONY: build
build:
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up builder
	docker cp $$(docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) ps -q builder):/go/bin/. build

.PHONY: release
release:
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) build
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) up agent
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) up migrate
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) up test

.PHONY: clean
clean:
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) kill
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) rm -f -v # -v removing dangling volumes
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) kill
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) rm -f -v # -v removing dangling volumes
	docker images -q -f dangling=true -f label=application=$(REPO_NAME) | xargs -I ARGS docker rmi -f ARGS

# =================================================================================================

.PHONY: env-up
env-up:
	${INFO} "Completed"

.PHONY: migrate-create
migrate-create:
	${INFO} "Creating migration..."
	docker-compose -p $(DEV_PROJECT) -f $(DOCKER_DEV_COMPOSE_FILE) run --rm migrate create -ext sql -dir migrations migration
	${INFO} "Completed"

.PHONY: migrate-up
migrate-up:
	${INFO} "Applying all migrations..."
	docker-compose -p $(DEV_PROJECT) -f $(DOCKER_DEV_COMPOSE_FILE) run --rm migrate -path migrations up
	${INFO} "Completed"

.PHONY: migrate-down
migrate-down:
	${INFO} "Rolling back last migration..."
	docker-compose -p $(DEV_PROJECT) -f $(DOCKER_DEV_COMPOSE_FILE) run --rm migrate -path migrations down 1
	${INFO} "Completed"

.PHONY: fmt
fmt:
	${INFO} "Applying fmt to the project..."
	go fmt $(go list ./... | grep -v /vendor/)

.PHONY: unit-test
unit-test:
	${INFO} "Running unit tests..."
	go test -v phrasebook-api/src/... -tags unit

.PHONY: test-all
test-all:
	${INFO} "Running unit and integration tests..."
	go test phrasebook-api/src/...

# =================================================================================================

# Cosmetics
YELLOW := "\e[1;33m"
NC := "\e[0m"

# Shell Functions
INFO := @bash -c '\
  printf $(YELLOW); \
  echo "=> $$1"; \
  printf $(NC)' VALUE