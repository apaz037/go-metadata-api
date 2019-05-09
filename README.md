# Golang Application Metadata API Server
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

A RESTful Golang API for persisting application metadata.

## Install
If a personal access token is available
```console
go get -u github.com/apaz037/go-metadata-api
```

>Note: if you are not prompted for git credentials when using go get, use `env GIT_TERMINAL_PROMPT=1` as a prefix

If a personal access token is not available but the repository is accessible by the user
```console
mkdir -p $GOPATH/src/github.com/apaz037/ && \
cd $GOPATH/src/github.com/apaz037 && \
git clone https://github.com/apaz037/go-metadata-api && \
cd go-metadata-api && \
env GO111MODULE=on go install
```
>Note: go-metadata-api uses go mod for dependencies, you must have Go 1.11 or higher installed in order to use. 

>Note: if go-metadata-api is not immediately available as a command line tool close your terminal window and start a new one.


## Serve
This command will start an https server on port `4200` 

> Note: go-metadata-api generates its own certs.  The code that generates these self signed certs is available in `api/utils/tls.go`.
> In order to avoid TLS handshake errors, you must add the `cert.pem` and `key.pem` to your http client.
> cert.pem and key.pem will only be generated the first time you run `go-metadata-api serve`, there is a check that looks for them before generation

```console
go-metadata-api serve
```

## Generate Docs
This command will generate a routes.md file detailing middlewares and linking to http handlers
```console
go-metadata-api gendoc
```

## Verify
For now, the postman collection is the best way to ensure functionality.
It can be found in test/postman_collection/

## Tests - WIP
```console
cd test && go test
```

## Docker
go-metadata-api utilizes a multi stage docker file.

The first container builds the binary with all of its dependencies.

The second container runs the application inside an alpine linux container and provides a lightweight way of interacting with the API over http while providing busybox for debugging purposes.

To run `go-metadata-api` via Docker
```console
docker run -p 4200:4200 aaronpaz/go-metadata-api:latest
```

[DockerHub](https://hub.docker.com/r/aaronpaz/go-metadata-api)

## TODO:
```console
/application
- GET: /application - done
- POST: /application - done
- GET: /application/{id} - done
- PUT: /application/{id} - done
- DELETE: /application/{id} - done
- GET: /application/search/{params} - done

/health
- GET: /health - done

miscellaneous:
json pretty printing - done
validation - done
doc generation command - done
statically linked go binary - done
Docker - done
/test - 15%
ci - 0% (docker container is built and published upon every commit to master however)
```