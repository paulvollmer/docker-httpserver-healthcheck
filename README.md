# docker-httpserver-healthcheck [![CI](https://github.com/paulvollmer/docker-httpserver-healthcheck/actions/workflows/ci.yml/badge.svg)](https://github.com/paulvollmer/docker-httpserver-healthcheck/actions/workflows/ci.yml)

To check if your Docker Container running a HTTP Server is healthy you can use the [Healthcheck](https://docs.docker.com/engine/reference/builder/#healthcheck) feature in Docker.    

Docker describe the healthcheck feature as following:

> The HEALTHCHECK instruction tells Docker how to test a container to check that it is still working.
> This can detect cases such as a web server that is stuck in an infinite loop and unable to handle new connections, even though the server process is still running.

The goal of this Project is to creat a small binary that can be used to do the `HEALTHCKECK`.

## Why not using `curl`

The main reason to not use curl is to avoid attack surfaces. An other reason is that you need to install curl at the base image and this is not given at all. 
If the docker setup is based on a multistage build, what you should do, it is not as easy to install curl. 
The last reason i want to mention is that you need to maintain curl and you need to release a new image if there is a new update.

These are the reasons you should use `docker-httpserver-healthcheck` for that job. 

## How to use

To include the `docker-httpserver-healthcheck` tool at your docker image you can add the following layer to your docker build.

```dockerfile
# The application build environment
# ...

# The healtcheck tool build environment 
FROM golang:1.14 as healthcheck-env
RUN go get github.com/paulvollmer/docker-httpserver-healthcheck
WORKDIR /go/src/github.com/paulvollmer/docker-httpserver-healthcheck
RUN go build -ldflags "-s -w -X main.healthcheckURL=http://localhost:8080/check" -o /healthcheck

# The final environment
FROM gcr.io/distroless/base
# Copy the healthcheck tool to the final container
COPY --from=healthcheck-env /healthcheck /healthcheck
HEALTHCHECK --start-period=10s --interval=5s --retries=3 CMD /healthcheck
# ...
```

## License

MIT License
