SHELL := /bin/zsh

install-all:


generate-all:
	mkdir -p internal/generated/schema
	oapi-codegen -package dto -generate chi-server,types,spec api/schema.yaml > internal/generated/schema/dto.gen.go

launch-dev-env:

test:

lint:
	golangci-lint run

build-all:

launch-server:

launch-consumer: