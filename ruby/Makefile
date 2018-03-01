.PHONY: test rails railsapi

test: rails railsapi

rails:
	cp Dockerfile-rails examples/rails/Dockerfile
	cd examples/rails && docker build -t docker-boilerplate-rails .
	docker rmi -f docker-boilerplate-rails

railsapi:
	cp Dockerfile-railsapi examples/railsapi/Dockerfile
	cd examples/railsapi && docker build -t docker-boilerplate-railsapi .
	docker rmi -f docker-boilerplate-railsapi