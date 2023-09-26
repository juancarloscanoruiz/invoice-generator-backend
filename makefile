include .env

dev:
	"$(GOP)/bin/air"
	
start:
	go run .

build:
	go build -o bin/invoice