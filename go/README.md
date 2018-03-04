# Golang `Dockerfile`s

One thing that is great about Golang is you can compile your Go application into a single binary making the distributon of a full fledged web application as easy as copying a binary file. This repo demonstrates two approaches to getting `Go`ing on Docker:

* Simple method: copy the source code into a Docker container, compile the application into a binary and reuse this container as our server
* Preferred method: using a 2 step process we first build the Go binary as normal *but* we then extract the binary out and build a second Docker image with *just* the binary in it. This second image is the one we will distribute to our Docker servers.

The preferred method has a great advantage of having a slim image size, we're not shipping around the Go build tools and the source code in our Docker image which will reduce pull time from our servers.

We also pass down `GIT_SHA` as a Docker build arg all the way into the Go compiler and expose it as `X-App-Version` header via some small middleware.