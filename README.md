# Dockerfile boilerplates

[![Build Status](https://travis-ci.org/callumj/docker-boilerplates.svg?branch=master)](https://travis-ci.org/callumj/docker-boilerplates)

A community contributed source for Dockerfiles built for your language or framework. Feel free to open a PR and contribute new `Dockerfile`s or improvements.

Status: currently a work in progress. Feel free to submit `Dockerfiles`!

## Does it work?

Try `make test`

## Contents

* [Ruby](ruby/README.md)
  * Rails 5.1
  * Rails API 5.1
* [Go](go/README.md)
  * Building and running in one image
  * Building in one image, second slim container with just our application binary (preferred)
* [Node](node/README.md)
  * Web server
  * Webpack with a slim image, built and saved to local machine