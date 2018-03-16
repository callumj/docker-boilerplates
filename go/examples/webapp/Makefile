GIT_SHA ?= $(shell git rev-parse --short HEAD)

run: build
	bin/app

build:
	mkdir -p bin
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.gitSHA=$(GIT_SHA)" -o bin/app .