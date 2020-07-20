# The application build environment
#----------------------------------
FROM golang:1.14 as build-env
WORKDIR /workspace
COPY . .
RUN go build -o server

# The healtcheck tool build environment
#--------------------------------------
FROM golang:1.14 as healthcheck-env
RUN go get github.com/paulvollmer/docker-httpserver-healthcheck
WORKDIR /go/src/github.com/paulvollmer/docker-httpserver-healthcheck
RUN go build -ldflags "-s -w -X main.healthcheckURL=http://localhost:8080/check" -o /healthcheck

# The final environment
#----------------------
FROM gcr.io/distroless/base
# Copy the healthcheck tool to the final container
COPY --from=healthcheck-env /healthcheck /healthcheck
HEALTHCHECK --start-period=10s --interval=5s --retries=3 CMD /healthcheck
# Copy the example http server
COPY --from=build-env /workspace/server /server
CMD [ "/server" ]

