# Makefile for Go project
APP_NAME := kumo-classroom

.PHONY: run build clean default

default: run

run:
	go run cmd/server/main.go

build:
	go build -o bin/app cmd/server/main.go

clean:
	rm -rf bin/
