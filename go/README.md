# Golang `Dockerfile`s

One thing that is great about Golang is you can compile your Go application into a single binary making the distributon of a full fledged web application as easy as copying a binary file. This repo demonstrates using mutli-stage Docker builds to getting `Go`ing on Docker:

We have a single `Dockerfile` that actually ends up producing two images, the first image is used to build the Go binary and the second image takes the single `app` binary from the previous image and copyies it into itself.

We also pass down `GIT_SHA` as a Docker build arg all the way into the Go compiler and expose it as `X-App-Version` header via some small middleware.