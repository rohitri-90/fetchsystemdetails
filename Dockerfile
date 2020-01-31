FROM golang:1.13 as builder

#ARG buildtime_variable=default_value
ENV MAC_ADDRESS=""
ENV MAC_API_TOKEN=""
WORKDIR /go/src/macapi
COPY . .
RUN go get
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build

# Run image
FROM docker:18.06.1-ce-dind

COPY --from=builder /go/src/macapi /bin/

ENTRYPOINT ["/usr/local/bin/dockerd-entrypoint.sh", "/bin/macapi"]