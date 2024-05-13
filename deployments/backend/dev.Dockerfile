FROM golang:1.22-alpine3.19 as builder
RUN mkdir /app
WORKDIR /app
ENV CGO_ENABLED=1
RUN apk update && apk add --no-cache gcc musl-dev
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT [ "air","dev" ]