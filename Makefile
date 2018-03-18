.PHONY: test ruby go node python

test: ruby go node python

ruby:
	cd ruby && make test

go:
	cd go && make test

node:
	cd node && make test

python:
	cd python && make test