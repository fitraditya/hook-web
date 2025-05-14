APP_NAME := hook-web
APP_VERSION := $(shell git describe --tags --always || git rev-parse HEAD)
APP_PKG := $(shell echo ${PWD} | sed -e "s\#${GOPATH}/src/\#\#g")

.PHONY: dev
dev:
	go run main.go server

.PHONY: build
build:
	@mkdir -p ./dist
	go mod tidy
	go build -ldflags "-X ${APP_PKG}/app/util.Name=${APP_NAME} -X ${APP_PKG}/app/util.Version=${APP_VERSION}" -o ./dist/hook-web main.go
