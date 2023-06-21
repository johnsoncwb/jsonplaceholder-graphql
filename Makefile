M = $(shell printf "\033[34;1m▶\033[0m")

server: ; $(info $(M) Starting development server...)
	@ go run cmd/app/main.go
.PHONY: server

image: ; $(info $(M) Creating docker image...)
	@ docker build -t graphql -f docker/app/dockerfile .
.PHONY: image

build: ; $(info $(M) Starting docker container with server...)
	@ docker run --publish 8000:8000 graphql
.PHONY: build