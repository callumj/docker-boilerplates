GIT_SHA = $(shell git rev-parse --short HEAD)

.PHONY: test clean web

test: web

clean:
	@docker rm -f docker-boilerplate-python ||:

web: clean
	cp Dockerfile examples/web/Dockerfile
	cd examples/web && docker build -t docker-boilerplate-python --build-arg GIT_SHA=$(GIT_SHA) .
	docker run --name docker-boilerplate-python-app -d -p 3000:3000 docker-boilerplate-python
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-python-app