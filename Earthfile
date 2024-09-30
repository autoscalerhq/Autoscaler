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
#---

temp:
    BUILD \
        --platform=linux/amd64 \
        +build-worker

build-all:
     BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm/v7 \
        --platform=linux/arm64 \
        --platform=linux/arm/v6 \
        +build-api

    BUILD \
        --platform=linux/amd64 \
        --platform=linux/arm/v7 \
        --platform=linux/arm64 \
        --platform=linux/arm/v6 \
        +build-worker

build-all-images:
     BUILD \
        --platform=linux/amd64 \
#        --platform=linux/arm/v7 \
#        --platform=linux/arm64 \
#        --platform=linux/arm/v6 \
        +build-image-api
#     BUILD \
#        --platform=linux/amd64 \
#        --platform=linux/arm/v7 \
#        --platform=linux/arm64 \
#        --platform=linux/arm/v6 \
#        +build-image-worker

#---
# Setup Dependencies
#---

setup-deps:
    FROM golang:1.22-alpine3.20
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
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG --required VARIANT
    RUN echo "TARGETPLATFORM=$TARGETPLATFORM, TARGETARCH=$TARGETARCH, TARGETVARIANT=$TARGETVARIANT"
    RUN echo "TARGETPLATFORM=$PLATFORM, TARGETARCH=$ARCH, TARGETVARIANT=$VARIANT"
    RUN exit 1
    RUN GOARM=${VARIANT#v} go build -o app services/api/main.go
    SAVE ARTIFACT ./api

build-worker:
    FROM +setup-deps
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG VARIANT
    RUN GOARM=${VARIANT#v} go build -o app services/worker/main.go
    SAVE ARTIFACT ./worker

#---
# Build Docker Images
#---

build-image-api:
    ARG TARGETPLATFORM
    ARG TARGETOS
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM alpine:3.20
    RUN false
    COPY --platform=$TARGETPLATFORM \
        (+build-api --GOARCH='${TARGETARCH}' --VARIANT='${TARGETVARIANT}') ./api
    ENTRYPOINT ["/api"]
    SAVE IMAGE --without-earthly-labels --push autoscaler/api:latest

build-image-worker:
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM alpine:3.20
    COPY \
        --platform=$TARGETPLATFORM \
        (+build-worker --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) ./worker
    ENTRYPOINT ["/worker"]
    SAVE IMAGE --without-earthly-labels --push autoscaler/worker:latest

