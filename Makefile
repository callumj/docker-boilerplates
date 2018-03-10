.PHONY: test ruby go node

test: ruby go node

ruby:
	cd ruby && make test

go:
	cd go && make test

node:
	cd node && make test