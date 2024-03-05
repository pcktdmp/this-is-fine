FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/cmd/fine
COPY src $GOPATH/src
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/fine
FROM scratch
COPY --from=builder /go/bin/fine /go/bin/fine
EXPOSE 8080
USER 9999:9999
ENTRYPOINT ["/go/bin/fine"]
