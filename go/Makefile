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
	# prepare the area
	mkdir -p examples/build_area
	rm -rf examples/build_area/app
	# copy in our build image Dockerfile
	cp Dockerfile-build examples/webapp/Dockerfile
	cd examples/webapp && docker build -t docker-boilerplate-go-build .
	# spin up the container that compiles our binary out into `build_area`
	docker run -e GIT_SHA=$(GIT_SHA) -v $(PWD)/examples/webapp:/go/src/github.com/callumj/docker-boilerplates/go/examples/webapp \
	-v $(PWD)/examples/build_area:/go/src/github.com/callumj/docker-boilerplates/go/examples/webapp/bin --rm docker-boilerplate-go-build make build
	cp Dockerfile-run examples/build_area/Dockerfile
	cd examples/build_area && docker build -t docker-boilerplate-go-preferred .
	docker run --name docker-boilerplate-go-preferred-app -d -p 3000:3000 docker-boilerplate-go-preferred
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-go-preferred-app