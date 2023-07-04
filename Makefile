run: build
	LISTEN_ADDRESS=127.0.0.1 LISTEN_PORT=8080 ./meteo-challenge

PHONY: clean

get:
	go get ./...

vendor: get
	go mod vendor

build: get vendor
	go build -o meteo-challenge ./cmd/meteo-challenge

gorun: vendor
	LISTEN_ADDRESS=127.0.0.1 LISTEN_PORT=8080 go run ./cmd/meteo-challenge

help: build
	./meteo-challenge --help


clean:
	rm -f ./meteo-challenge
