GIT_SHA = $(shell git rev-parse --short HEAD)

.PHONY: multi_stage

test: multi_stage

clean:
	@docker rm -f docker-boilerplate-go-simple-app ||:

multi_stage: clean
	cp Dockerfile examples/webapp/Dockerfile
	cd examples/webapp && docker build -t docker-boilerplate-go --build-arg GIT_SHA=$(GIT_SHA) .
	docker run --name docker-boilerplate-go-simple-app -d -p 3000:3000 docker-boilerplate-go-simple
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-go-simple-app