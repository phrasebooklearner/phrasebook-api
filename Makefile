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

# Check and inspect logic
INSPECT := $$(docker-compose -p $$1 -f $$2 ps -q $$3 | xargs -I ARGS docker inspect -f "{{ .State.ExitCode }}" ARGS)

CHECK := @bash -c '\
  if [[ $(INSPECT) -ne 0 ]]; \
  then exit $(INSPECT); fi' VALUE

# Use these settings to specify a custom Docker registry
DOCKER_REGISTRY ?= docker.io

# WARNING: Set DOCKER_REGISTRY_AUTH to empty for Docker Hub
# Set DOCKER_REGISTRY_AUTH to auth endpoint for private Docker registry
DOCKER_REGISTRY_AUTH ?=

# Application Service Name - must match Docker Compose release specification application service name
APP_SERVICE_NAME := app

# Build tag expression - can be used to evaulate a shell expression at runtime
BUILD_TAG_EXPRESSION ?= date -u +%Y%m%d%H%M%S

# Execute shell expression
BUILD_EXPRESSION := $(shell $(BUILD_TAG_EXPRESSION))

# Build tag - defaults to BUILD_EXPRESSION if not defined
BUILD_TAG ?= $(BUILD_EXPRESSION)

.PHONY: test
test:
	docker volume create --name cache
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) pull
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) build --pull test
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) run --rm agent
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) run --rm migrate
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up test
	docker cp $$(docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) ps -q test):/reports/. reports
	${CHECK} $(DEV_PROJECT) $(DEV_COMPOSE_FILE) test

.PHONY: build
build:
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) build builder
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) up builder
	${CHECK} $(DEV_PROJECT) $(DEV_COMPOSE_FILE) builder
	docker cp $$(docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) ps -q builder):/go/bin/. build

.PHONY: release
release:
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) pull test
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) build app
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) build --pull nginx
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) run --rm agent
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) run --rm migrate
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) up test
	docker cp $$(docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) ps -q test):/reports/. reports
	${CHECK} $(REL_PROJECT) $(REL_COMPOSE_FILE) test

.PHONY: clean
clean:
	docker-compose -p $(DEV_PROJECT) -f $(DEV_COMPOSE_FILE) down -v
	docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) down -v
	docker images -q -f dangling=true -f label=application=$(REPO_NAME) | xargs -I ARGS docker rmi -f ARGS

.PHONY: tag
tag:
	$(foreach tag,$(TAG_ARGS), docker tag $(IMAGE_ID) $(DOCKER_REGISTRY)/$(ORG_NAME)/$(REPO_NAME):$(tag);)

.PHONY: buildtag
buildtag:
	$(foreach tag,$(BUILDTAG_ARGS), docker tag $(IMAGE_ID) $(DOCKER_REGISTRY)/$(ORG_NAME)/$(REPO_NAME):$(tag).$(BUILD_TAG);)

.PHONY: login
login:
	docker login -u $$DOCKER_USER -p $$DOCKER_PASSWORD $$DOCKER_REGISTRY_AUTH

.PHONY: logout
logout:
	docker logout

.PHONY: publish
publish:
	$(foreach tag,$(shell echo $(REPO_EXPR)), docker push $(tag);)
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

# Get container id of application service container
APP_CONTAINER_ID := $$(docker-compose -p $(REL_PROJECT) -f $(REL_COMPOSE_FILE) ps -q $(APP_SERVICE_NAME))

# Get image id of application service
IMAGE_ID := $$(docker inspect -f '{{ .Image }}' $(APP_CONTAINER_ID))

# Repository Filter
ifeq ($(DOCKER_REGISTRY), docker.io)
  REPO_FILTER := $(ORG_NAME)/$(REPO_NAME)[^[:space:]|\$$]*
else
  REPO_FILTER := $(DOCKER_REGISTRY)/$(ORG_NAME)/$(REPO_NAME)[^[:space:]|\$$]*
endif

# Introspect repository tags
REPO_EXPR := $$(docker inspect -f '{{range .RepoTags}}{{.}} {{end}}' $(IMAGE_ID) | grep -oh "$(REPO_FILTER)" | xargs)

# Extract build tag arguments
ifeq (buildtag,$(firstword $(MAKECMDGOALS)))
  BUILDTAG_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  ifeq ($(BUILDTAG_ARGS),)
    $(error You must specify a tag)
  endif
  $(eval $(BUILDTAG_ARGS):;@:)
endif

# Extract tag arguments
ifeq (tag,$(firstword $(MAKECMDGOALS)))
  TAG_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  ifeq ($(TAG_ARGS),)
    $(error You must specify a tag)
  endif
  $(eval $(TAG_ARGS):;@:)
endif