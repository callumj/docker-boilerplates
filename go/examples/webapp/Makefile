GIT_SHA ?= $(shell git rev-parse --short HEAD)

run: build
	bin/app

build:
	mkdir -p bin
	go build -ldflags "-X main.gitSHA=$(GIT_SHA)" -o bin/app .