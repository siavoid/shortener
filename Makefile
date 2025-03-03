-include .env
DEFAULT_GOAL := local

.PHONY: lint
lint:
	go lint ./...

.PHONY: run
run:
	go run ./cmd/shortener/main.go

.PHONY: build
build:
	go build -o ./cmd/shortener/shortener.exe ./cmd/shortener/main.go

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: autotests1
autotests1:
	shortenertest -test.v -test.run=^TestIteration1$$ -binary-path=./cmd/shortener/shortener.exe

.PHONY: autotests2
autotests2:
	shortenertest -test.v -test.run=^TestIteration2$$ -binary-path=./cmd/shortener/shortener.exe

.PHONY: vet
vet:
	go vet -vettool=$(shell where statictest.exe) .\...

.PHONY: check
check: vet build autotests