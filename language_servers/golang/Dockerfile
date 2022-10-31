FROM golang:1.19.2-alpine3.16 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /temp
COPY . /temp
RUN go build -o main /temp/language_servers/golang

FROM ubuntu
ENV GIN_MODE=release
COPY --from=builder /temp/main /
COPY --from=builder /usr/local/go/bin/gofmt /
CMD ["/main"]
