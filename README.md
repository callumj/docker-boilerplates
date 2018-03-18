# Dockerfile boilerplates

[![Build Status](https://travis-ci.org/callumj/docker-boilerplates.svg?branch=master)](https://travis-ci.org/callumj/docker-boilerplates)

A community contributed source for Dockerfiles built for your language or framework. Feel free to open a PR and contribute new `Dockerfile`s or improvements.

Status: currently a work in progress. Feel free to submit `Dockerfiles`!

## Does it work?

We ensure every example works by requiring each example to be tested using `make test`. You can see the test structure in each sub `Makefile`.

## Contents

* [Ruby](ruby/README.md)
  * Rails 5.1
  * Rails API 5.1
* [Go](go/README.md)
  * Building and running a Go web server using Docker's multi-stage builds
* [Node](node/README.md)
  * Web server with Express
  * Webpack with a slim image, built and saved to local machine
* [Python](python/README.md)
  * Web server with Flask & Gunicorn