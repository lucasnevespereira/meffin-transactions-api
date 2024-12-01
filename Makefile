.PHONY: lint build run test clean deps mocks proxy-db connect-db deploy run-dev


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

proxy-db:
	fly proxy 5432 -a meffin-transactions-api-db

connect-db:
	fly postgres connect -a meffin-transactions-api-db

deploy:
	flyctl deploy
