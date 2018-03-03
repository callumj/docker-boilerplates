.PHONY: test rails railsapi

test: rails railsapi

rails:
	cp Dockerfile-rails examples/rails/Dockerfile
	cd examples/rails && docker build -t docker-boilerplate-rails .
	docker run --name docker-boilerplate-rails-app -d -p 3000:3000 -e SECRET_KEY_BASE=secret -e RAILS_ENV=production docker-boilerplate-rails rails s
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-rails-app

railsapi:
	cp Dockerfile-railsapi examples/railsapi/Dockerfile
	cd examples/railsapi && docker build -t docker-boilerplate-railsapi .
	docker run --name docker-boilerplate-railsapi-app -d -p 3000:3000 -e SECRET_KEY_BASE=secret -e RAILS_ENV=production docker-boilerplate-railsapi rails s
	../scripts/wait-for-success.sh curl http://localhost:3000/hello_world
	docker rm -f docker-boilerplate-railsapi-app