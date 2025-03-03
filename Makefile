-include .env
DEFAULT_GOAL := local


.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	go lint ./...

.PHONY: run
run:
	go run ./cmd/shortener/main.go

.PHONY: build
build:
	go build -o ./cmd/shortener/shortener.exe ./cmd/shortener/main.go

.PHONY: autotests
autotests:
	shortenertest -test.v -test.run=^TestIteration1$ -binary-path=./cmd/shortener/shortener.exe