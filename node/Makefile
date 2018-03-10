GIT_SHA = $(shell git rev-parse --short HEAD)

.PHONY: test clean web webpack

test: web webpack

clean:
	@docker rm -f docker-boilerplate-node-web-app ||:

web: clean
	cp Dockerfile-web examples/web/Dockerfile
	cd examples/web && docker build -t docker-boilerplate-node-web --build-arg GIT_SHA=$(GIT_SHA) .
	docker run --name docker-boilerplate-node-web-app -d -p 3000:3000 docker-boilerplate-node-web
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-node-web-app

webpack:
	mkdir -p examples/build_area
	rm -rf examples/build_area/*
	cp Dockerfile-webpack examples/webpack-sample/Dockerfile
	cd examples/webpack-sample && docker build -t docker-boilerplate-node-webpack .
	docker run -e GIT_SHA=$(GIT_SHA) -v $(PWD)/examples/webpack-sample:/app -v $(PWD)/examples/build_area:/app/dist --rm docker-boilerplate-node-webpack npm start
	test examples/build_area/index.html
	test examples/build_area/bundle.js