FROM golang:1.10
# build the full go path for src
# this will ensure imports will work correctly
WORKDIR /go/src/github.com/callumj/docker-boilerplates/go/examples/webapp
COPY . .

# accept GIT_SHA as a build arg from Docker
# so we can pass into the Go compiler the value of gitSHA
# for X-App-Version
ARG GIT_SHA
ENV GIT_SHA ${GIT_SHA:-unknown}

# invoke the build process
RUN make build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# begin a new Dockerfile which will contain our slim
# image for hosting just our Go binary

# use alpine to keep our image small
FROM alpine:latest

# install certs for SSL
RUN apk --no-cache add ca-certificates

# setup directory to host binary
WORKDIR /root/

# copy in from previous build step
COPY --from=0 /go/src/github.com/callumj/docker-boilerplates/go/examples/webapp/bin/app .

CMD ["./app"]