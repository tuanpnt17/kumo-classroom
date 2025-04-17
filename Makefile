# Makefile for Go project
APP_NAME := kumo-classroom

.PHONY: run build clean default di

default: run

run:
	go run cmd/server/main.go

di:
	wire ./internal/wire

build:
	go build -o bin/app cmd/server/main.go

clean:
	rm -rf bin/
