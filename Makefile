include .env

DEFAULT_GOAL := local
BIN_PATH=./cmd/shortener/shortener.exe
SHORTENER_TEST=shortenertest
SHORTENER_TEST_BETA=shortenertestbeta
TEMP_STORE_FILE=./tmp/short-url-db.json

POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=praktikum
PG_IMAGE_NAME=shorten_postgres
DATABASE_DSN='postgres://postgres:postgres@localhost:5432/praktikum?sslmode=disable'

.PHONY: lint
lint:
	go lint ./...

.PHONY: run
run:
	go run ./cmd/shortener/main.go

.PHONY: swaginit
swaginit:
	swag init -g ./internal/controllers/http/v1/server.go

.PHONY: build
build:
	go build -o $(BIN_PATH) ./cmd/shortener/main.go

.PHONY: pgbuild
pgbuild:
	docker build --build-arg POSTGRES_USER=$(POSTGRES_USER) \
	--build-arg POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	--build-arg POSTGRES_DB=$(POSTGRES_DB) \
	-t $(PG_IMAGE_NAME) ./deploy/pg/

.PHONY: pgrun
pgrun:
	docker run --rm -p 5432:5432 $(PG_IMAGE_NAME)

.PHONY: test
test:
	go test -count=1 ./...

.PHONY: vet
vet:
	go vet -vettool=$(shell where statictest.exe) .\...

.PHONY: check
check: vet build autotest

.PHONY: autotest
autotest: \
	build \
	autotest1 \
	autotest2 \
	autotest3 \
	autotest4 \
	autotest5 \
	autotest6 \
	autotest7 \
	autotest8 \
	autotest9 \
	autotest10


.PHONY: autotest1
autotest1:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST) -test.v -test.run=^TestIteration1$$ -binary-path=$(BIN_PATH)
	

.PHONY: autotest2
autotest2:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST) -test.v -test.run=^TestIteration2$$ -source-path=.

.PHONY: autotest3
autotest3:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST) -test.v -test.run=^TestIteration3$$ -source-path=.

.PHONY: autotest4
autotest4:
	@type nul > $(TEMP_STORE_FILE)
	$(eval SERVER_PORT=$(shell python -c "import socket; s=socket.socket(); s.bind(('', 0)); print(s.getsockname()[1]); s.close()"))
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration4$$ \
	-binary-path=$(BIN_PATH) \
	-server-port=$(SERVER_PORT) \
	-source-path=.

.PHONY: autotest5
autotest5:
	@type nul > $(TEMP_STORE_FILE)
	$(eval SERVER_PORT=$(shell python -c "import socket; s=socket.socket(); s.bind(('', 0)); print(s.getsockname()[1]); s.close()"))
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration5$$ \
	-binary-path=$(BIN_PATH) \
	-server-port=$(SERVER_PORT) \
	-source-path=.


.PHONY: autotest6
autotest6:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration6$$ -source-path=.

.PHONY: autotest7
autotest7:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration7$$ -source-path=. \
	-binary-path=$(BIN_PATH)

.PHONY: autotest8
autotest8:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration8$$ \
	-binary-path=$(BIN_PATH)

.PHONY: autotest9
autotest9:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration9$$ \
	-binary-path=$(BIN_PATH) \
	-source-path=. \
	-file-storage-path=$(TEMP_STORE_FILE)


.PHONY: autotest10
autotest10:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration10$$ \
	-binary-path=$(BIN_PATH) \
	-source-path=. \
	-database-dsn=$(DATABASE_DSN)

.PHONY: autotest11
autotest11:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration11$$ \
	-binary-path=$(BIN_PATH) \
	-database-dsn=$(DATABASE_DSN)

.PHONY: autotest12
autotest12:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration12$$ \
	-binary-path=$(BIN_PATH) \
	-database-dsn=$(DATABASE_DSN)

.PHONY: autotest13
autotest13:
	@type nul > $(TEMP_STORE_FILE)
	$(SHORTENER_TEST_BETA) -test.v -test.run=^TestIteration13$$ \
	-binary-path=$(BIN_PATH) \
	-database-dsn=$(DATABASE_DSN)

