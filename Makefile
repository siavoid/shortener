-include .env
DEFAULT_GOAL := local
BIN_PATH=./cmd/shortener/shortener.exe
SHORTENER_TEST=shortenertest
SHORTENER_TEST_BETA=shortenertestbeta

.PHONY: lint
lint:
	go lint ./...

.PHONY: run
run:
	go run ./cmd/shortener/main.go

.PHONY: build
build:
	go build -o $(BIN_PATH) ./cmd/shortener/main.go

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: autotest1
autotest1:
	$(SHORTENER_TEST) -test.v -test.run=^TestIteration1$$ -binary-path=$(BIN_PATH)

.PHONY: autotest2
autotest2:
	$(SHORTENER_TEST) -test.v -test.run=^TestIteration2$$ -source-path=.

.PHONY: vet
vet:
	go vet -vettool=$(shell where statictest.exe) .\...

.PHONY: check
check: vet build autotest