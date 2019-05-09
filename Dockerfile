# Multi Stage Dockerfile
## First container will build our application
## Build is performed in such a way that all dependencies (including c dependencies) are statically linked
## This gives us a very lightweight final image (roughly 20mb) as measured via Dive

# Builder Container
# using latest golang image by default
FROM golang:latest as builder

# assuming root out of convenience
USER 0

# force enable go mod
ENV GO111MODULE=on

# update OS in case anything is out of date
RUN apt-get update && apt-get install -y

# create our app's working directory where we will compile our binary
RUN mkdir -p /go/src/github.com/apaz037/go-metadata-api

# set working directory
WORKDIR /go/src/github.com/apaz037/go-metadata-api

# copy over current source code
COPY . .

# compile statically linked linux binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' .

# this will be our final application container, we use alpine because it bundles busybox which is nice for troubleshooting
FROM alpine:latest

# update OS in case anything is out of date
RUN apk upgrade && apk update

# add ca-certs for https in the future
RUN apk add ca-certificates

# create a directory to hold our binary
RUN mkdir -p /var/app/

# set working directory
WORKDIR /var/app

# copy the compiled binary from our builder container
COPY --from=builder /go/src/github.com/apaz037/go-metadata-api/go-metadata-api /var/app

# expose port 4200, the default for our server.  This could easily be an environment variable
EXPOSE 4200

# healthcheck our container to ensure server is up and responding
HEALTHCHECK CMD wget -O/dev/null -q https://localhost:4200/health || exit 1

# run our application server
ENTRYPOINT ["/var/app/go-metadata-api", "serve"]