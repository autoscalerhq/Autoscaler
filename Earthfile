VERSION 0.8

all-unit-test:
    BUILD ./libs/hello+unit-test
    BUILD ./services/api+unit-test
    BUILD ./services/worker+unit-test

all-docker:
    BUILD ./services/api+docker
    BUILD ./services/worker+docker

all-release:
    BUILD ./services/api+release
    BUILD ./services/worker+release

dev-up:
    LOCALLY
    RUN docker-compose -f ./docker/docker-compose.yml up

dev-down:
    LOCALLY
    RUN docker-compose -f ./docker/docker-compose.yml down

mon-up:
    LOCALLY
    RUN docker-compose -f ./docker/docker-compose.monitoring.yml up

mon-down:
    LOCALLY
    RUN docker-compose -f ./docker/docker-compose.monitoring.yml down