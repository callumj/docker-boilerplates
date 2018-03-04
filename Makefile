.PHONY: test ruby go

test: ruby go

ruby:
	cd ruby && make test

go:
	cd go && make test