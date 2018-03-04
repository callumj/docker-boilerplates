GIT_SHA = $(shell git rev-parse --short HEAD)

.PHONY: test clean rails railsapi

test: rails railsapi

clean:
	@docker rm -f docker-boilerplate-rails-app docker-boilerplate-railsapi-app ||:

rails: clean
	cp Dockerfile-rails examples/rails/Dockerfile
	cd examples/rails && docker build -t docker-boilerplate-rails --build-arg GIT_SHA=$(GIT_SHA) .
	docker run --name docker-boilerplate-rails-app -d -p 3000:3000 -e SECRET_KEY_BASE=secret -e RAILS_ENV=production docker-boilerplate-rails rails s
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-rails-app

railsapi: clean
	cp Dockerfile-railsapi examples/railsapi/Dockerfile
	cd examples/railsapi && docker build -t docker-boilerplate-railsapi --build-arg GIT_SHA=$(GIT_SHA) .
	docker run --name docker-boilerplate-railsapi-app -d -p 3000:3000 -e SECRET_KEY_BASE=secret -e RAILS_ENV=production docker-boilerplate-railsapi rails s
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-railsapi-app