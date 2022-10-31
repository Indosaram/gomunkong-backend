FROM golang:1.19.2-alpine3.16 as builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /temp
COPY . /temp
RUN go build -o main /temp/language_servers/python

FROM python:3.8.10-slim-buster
ENV GIN_MODE=release
COPY --from=builder /temp/main /
COPY --from=builder /temp/language_servers/python/requirements.txt /
RUN python3 -m pip install --upgrade pip &&
    python3 -m pip install -r requirements.txt
CMD ["/main"]
