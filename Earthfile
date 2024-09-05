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

build-all-platforms:
     BUILD \
            --platform=linux/amd64 \
            --platform=linux/arm/v7 \
            --platform=linux/arm64 \
            --platform=linux/arm/v6 \
            +build-image

setup-deps:
    FROM golang:1.22-alpine3.20
    WORKDIR /app
    RUN apk update && apk add --no-cache git
    COPY go.mod go.sum .
    COPY ./services ./services
    COPY ./lib ./lib
    RUN go mod download
    RUN go mod tidy

build:
    FROM +setup-deps
    ARG GOOS=linux
    ARG GOARCH=amd64
    ARG VARIANT
    RUN GOARM=${VARIANT#v} go build -o app services/api/main.go
    SAVE ARTIFACT ./app

build-image:
    ARG TARGETPLATFORM
    ARG TARGETARCH
    ARG TARGETVARIANT
    FROM --platform=$TARGETPLATFORM alpine:3.20
    COPY \
        --platform=$TARGETPLATFORM \
         (+build/app --GOARCH=$TARGETARCH --VARIANT=$TARGETVARIANT) ./app
    ENTRYPOINT ["/app"]
    SAVE IMAGE --without-earthly-labels --push autoscaler/api:latest