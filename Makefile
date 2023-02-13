.PHONY: build

build:
	go build -v ./cmd/app

test:
	go test ./internal/repository/tests

.DEFAULT_GOAL := build