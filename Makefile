.PHONY: lint build run test clean deps mocks

fly-db-proxy:
	flyctl proxy 5432 -a meffin-transactions-api-db

lint:
	golangci-lint run --exclude-use-default=true

clean:
		rm -rf mocks

deps: clean
		go get -d -v ./...

mocks: clean deps
		mockery --all

test: mocks
	go test -v ./...

run-dev:
	docker-compose up

build-dev:
	docker-compose up --build