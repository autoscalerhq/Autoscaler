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
# Note Earthly only supports amd64 and arm64
#---
build-all:
     BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm64 \
        +build-api
    BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm64 \
        +build-worker

build-all-images:
     BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm64 \
        +build-image-api
     BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm64 \
        +build-image-worker

#---
# Setup Dependencies
#---
setup-deps:
    FROM golang:1.23-alpine3.20
    WORKDIR /app
    RUN apk update && apk add --no-cache git
    COPY go.mod go.sum .
    COPY ./internal ./internal
    COPY ./services ./services
    COPY ./lib ./lib
    RUN go mod download
    RUN go mod tidy

#---
# Build Services
#---
build-api:
    FROM +setup-deps
    WORKDIR /app
    ARG GOOS=linux
    ARG TARGETARCH
    ARG VARIANT
    RUN GOARM=${VARIANT#v} GOARCH=$TARGETARCH go build -o api services/api/main.go
    SAVE ARTIFACT ./api
    #AS LOCAL ./tmp/api-$TARGETARCH

build-worker:
    FROM +setup-deps
    WORKDIR /app
    ARG GOOS=linux
    ARG TARGETARCH
    ARG VARIANT
    RUN GOARM=${VARIANT#v} GOARCH=$TARGETARCH go build -o worker services/worker/main.go
    SAVE ARTIFACT ./worker
    #AS LOCAL ./tmp/worker-$TARGETARCH

#---
# Build Docker Images
#---
build-image-api:

    ARG TARGETPLATFORM
    ARG TARGETOS
    ARG TARGETARCH
    ARG TARGETVARIANT
     FROM --platform=$TARGETPLATFORM alpine:3.20
    COPY (+build-api/api --VARIANT=$TARGETVARIANT) ./app
    ENTRYPOINT ["/app"]
    SAVE IMAGE --push autoscaler/api:latest

build-image-worker:
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM alpine:3.20
    COPY \
        --platform=linux/amd64 \
        (+build-worker/worker --VARIANT=$TARGETVARIANT) ./app
    ENTRYPOINT ["/app"]
    SAVE IMAGE --push autoscaler/worker:latest

