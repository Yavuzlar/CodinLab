FROM golang:1.23-alpine3.19 as builder
RUN mkdir /app
WORKDIR /app
ENV CGO_ENABLED=1
RUN apk update && apk add --no-cache docker-cli gcc musl-dev
RUN go install github.com/air-verse/air@latest
ENTRYPOINT [ "air","dev" ]