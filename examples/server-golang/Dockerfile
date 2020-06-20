# The application build environment
#----------------------------------
FROM golang:1.14 as app-build
WORKDIR /workspace
COPY . .
RUN go build -o server

# The healtcheck tool build environment
#--------------------------------------
FROM golang:1.14 as healthcheck-build
RUN go get github.com/paulvollmer/docker-httpserver-healthcheck
WORKDIR /go/src/github.com/paulvollmer/docker-httpserver-healthcheck
RUN go build -ldflags "-s -w -X main.healthcheckURL=http://localhost:8080/check" -o /healthcheck

# The final environment
#----------------------
FROM gcr.io/distroless/base
# Copy the healthcheck tool to the sinal container
COPY --from=healthcheck-build /healthcheck /healthcheck
HEALTHCHECK --start-period=10s --interval=5s CMD ["/healthcheck"]
# Copy the example http server
COPY --from=app-build /workspace/server /server
CMD [ "/server" ]

