FROM golang:1.24-alpine3.20 AS builder

WORKDIR /app

ENV CGO_ENABLED=1

RUN apk update && apk add --no-cache docker-cli gcc musl-dev

# Spesifik ve stabil air versiyonu kurulumu
RUN go install github.com/air-verse/air@v1.62.0

ENTRYPOINT ["air", "dev"]
