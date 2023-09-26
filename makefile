include .env

dev:
	"$(GOPATH)/bin/air"
	
start:
	go run .

build:
	go build -o bin/invoice