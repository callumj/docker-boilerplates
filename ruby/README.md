# Ruby `Dockerfile`s

* `Dockerfile-rails`: a Dockerfile supporting Rails 5.1 (Ruby 2.5.0)
  * NodeJS and Yarn are installed to support the Rails Asset Pipeline
  * Assets are precompile into the image, which may cover most usecases