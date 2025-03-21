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

.PHONY: autotest3
autotest3:
	$(SHORTENER_TEST) -test.v -test.run=^TestIteration3$$ -source-path=.

.PHONY: autotest4
autotest4:
	$(eval SERVER_PORT=$(shell python -c "import socket; s=socket.socket(); s.bind(('', 0)); print(s.getsockname()[1]); s.close()"))
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration4$$ \
	-binary-path=$(BIN_PATH) \
	-server-port=$(SERVER_PORT) \
	-source-path=.

.PHONY: autotest5
autotest5:
	$(eval SERVER_PORT=$(shell python -c "import socket; s=socket.socket(); s.bind(('', 0)); print(s.getsockname()[1]); s.close()"))
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration5$$ \
	-binary-path=$(BIN_PATH) \
	-server-port=$(SERVER_PORT) \
	-source-path=.


.PHONY: autotest
autotest: \
	build \
	autotest1 \
	autotest2 \
	autotest3 \
	autotest4 \
	autotest5


.PHONY: vet
vet:
	go vet -vettool=$(shell where statictest.exe) .\...

.PHONY: check
check: vet build autotest