FROM golang:1.16.5-alpine3.13 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /temp
COPY . /temp
RUN go build -o main /temp/formatter 

FROM scratch
ENV GIN_MODE=release
COPY --from=builder /temp/main /
CMD ["/main"]