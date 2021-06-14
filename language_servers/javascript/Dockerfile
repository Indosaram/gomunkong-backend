FROM golang:1.16.5-alpine3.13 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /temp
COPY . /temp
RUN go build -o main /temp/language_servers/javascript

FROM node:16-alpine3.11
ENV GIN_MODE=release
WORKDIR /usr/app/
COPY --from=builder /temp/main /
COPY --from=builder /temp/language_servers/javascript/package.json /usr/app/
RUN npm install
CMD ["/main"]