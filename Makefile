-include .env
DEFAULT_GOAL := local


.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint:
	golint ./...

.PHONY: run
run:
	go run ./cmd/shortener/main.go
