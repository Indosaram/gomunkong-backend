FROM golang:1.19.2-alpine3.16 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /temp
COPY . /temp
RUN go build -o main /temp/language_servers/java

FROM openjdk:14-ea-jdk-alpine3.10
ENV GIN_MODE=release
COPY --from=builder /temp/main /
RUN wget https://github.com/google/google-java-format/releases/download/v1.10.0/google-java-format-1.10.0-all-deps.jar -O google-java-format.jar
CMD ["/main"]
