GIT_SHA = $(shell git rev-parse --short HEAD)

.PHONY: test clean web webpack

test: web webpack

clean:
	@docker rm -f docker-boilerplate-node-web-app docker-boilerplate-node-webpack-staging ||:

web: clean
	cp Dockerfile-web examples/web/Dockerfile
	cd examples/web && docker build -t docker-boilerplate-node-web --build-arg GIT_SHA=$(GIT_SHA) .
	docker run --name docker-boilerplate-node-web-app -d -p 3000:3000 docker-boilerplate-node-web
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-node-web-app

webpack:
	mkdir -p examples/build_area
	cp Dockerfile-webpack examples/webpack-sample/Dockerfile
	cd examples/webpack-sample && docker build -t docker-boilerplate-node-webpack --build-arg GIT_SHA=$(GIT_SHA) .
	docker run -v $(PWD)/examples/build_area:/host_tmp --name docker-boilerplate-node-webpack-staging docker-boilerplate-node-webpack cp -r dist /host_tmp/
	docker rm -f docker-boilerplate-node-webpack-staging
	test examples/build_area/dist/index.html
	test examples/build_area/dist/bundle.js