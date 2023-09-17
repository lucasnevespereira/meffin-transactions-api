.PHONY: lint build run test clean deps mocks

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

run:
	docker-compose up

build:
	docker-compose up --build