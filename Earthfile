VERSION  0.8

#---
# Dev
#---

dev-up:
    LOCALLY
    RUN docker-compose -f ./docker/docker-compose.yml up

dev-down:
    LOCALLY
    RUN docker-compose -f ./docker/docker-compose.yml down

#---
# Building
# linux/riscv64 is beta supported by node. We only release versions that can be supported by all of our apps.
# linux/ppc64le is not suppored by node-alpine images.
#---
build-all-images:
     BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm64 \
        --platform=linux/arm/v7 \
        --platform=linux/s390x \
        +build-image-api
     BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm64 \
        --platform=linux/arm/v7 \
        --platform=linux/s390x \
        +build-image-worker

    BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm64 \
        --platform=linux/arm/v7 \
        --platform=linux/s390x \
        +build-image-webapp

#---
# Setup Dependencies
#---
be-setup-deps:
    FROM golang:1.23-alpine3.20
    WORKDIR /app
#    Only add if a tool is needed to be installed and ensure to pin the version
#    && apk add git=2.23.0
    RUN apk update
    COPY go.mod go.sum .
    COPY ./internal ./internal
    COPY ./services ./services
    COPY ./lib ./lib
    RUN go mod download
    RUN go mod tidy

fe-setup-deps:
    FROM oven/bun:latest
    WORKDIR /app
    COPY ./webapp .


#---
# Build Services
#---
build-api:
    FROM +be-setup-deps
    WORKDIR /app
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG VARIANT
    RUN GOARM=${VARIANT#v} go build -o api services/api/main.go
    SAVE ARTIFACT ./api

build-worker:
    FROM +be-setup-deps
    WORKDIR /app
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG VARIANT
    RUN GOARM=${VARIANT#v} go build -o worker services/worker/main.go
    SAVE ARTIFACT ./worker

build-webapp:
    FROM +fe-setup-deps
    RUN bun i && bun run build
    SAVE ARTIFACT .next/standalone

#---
# Build Docker Images
#---
build-image-api:
    ARG TARGETPLATFORM
    ARG TARGETOS
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM alpine:3.20
    COPY \
    # if this is set to --platform=$TARGETPLATFORM then the build will happen on that architecture.
    # However this is not garunteed to actually compile and the only one that is, is amd64.
        --platform=linux/amd64 \
        (+build-api/api --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) ./app
    ENTRYPOINT ["/app"]
    SAVE IMAGE --push ghcr.io/autoscalerhq/api:latest

build-image-worker:
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM alpine:3.20
    COPY \
    # if this is set to --platform=$TARGETPLATFORM then the build will happen on that architecture.
    # However this is not garunteed to actually compile and the only one that is, is amd64.
        --platform=linux/amd64 \
        (+build-worker/worker --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) ./app
    ENTRYPOINT ["/app"]
    SAVE IMAGE --push ghcr.io/autoscalerhq/worker:latest

build-image-webapp:
    FROM node:22-alpine3.20
    COPY --platform=linux/amd64 (+build-webapp/standalone ) ./app
    SAVE IMAGE --push ghcr.io/autoscalerhq/web:latest

