GIT_SHA = $(shell git rev-parse --short HEAD)

.PHONY: test simple preferred

test: simple preferred

clean:
	@docker rm -f docker-boilerplate-go-simple-staging docker-boilerplate-go-simple-app docker-boilerplate-go-preferred-app ||:

simple: clean
	cp Dockerfile-simple examples/webapp/Dockerfile
	cd examples/webapp && docker build -t docker-boilerplate-go-simple --build-arg GIT_SHA=$(GIT_SHA) .
	docker run --name docker-boilerplate-go-simple-app -d -p 3000:3000 docker-boilerplate-go-simple
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-go-simple-app

preferred: clean
	cp Dockerfile-simple examples/webapp/Dockerfile
	cd examples/webapp && docker build -t docker-boilerplate-go-simple --build-arg GIT_SHA=$(GIT_SHA) .
	docker run -v $(PWD)/examples/build_area:/host_tmp --name docker-boilerplate-go-simple-staging docker-boilerplate-go-simple cp bin/app /host_tmp
	docker rm -f docker-boilerplate-go-simple-staging
	cp Dockerfile-run examples/build_area/Dockerfile
	cd examples/build_area && docker build -t docker-boilerplate-go-preferred .
	docker run --name docker-boilerplate-go-preferred-app -d -p 3000:3000 docker-boilerplate-go-preferred
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-go-preferred-app