# Node `Dockerfile`s

Two different Node applications are being used here with Dockerfiles, one running an express web server (a server side application) and another building a client side application using webpack.

* Web: the whole application is copied into the Docker image to easily transport the Docker image.
* Webpack: the Docker image only contains the installed dependencies from package.json and then we spin it up as a Docker image with the whole directory mounted int so we can quickly compile the Webpack client app. This keeps the build time and size down on the image because we should _eventually_ rarely change our `package.json`.

We also pass down `GIT_SHA` as a Docker build arg as a environment variable and expose it as `X-App-Version` header via some small middleware.