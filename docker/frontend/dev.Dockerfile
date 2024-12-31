FROM node:20.13.1-alpine AS builder
ENV TZ=UTC 
WORKDIR /app
RUN apk update && apk add ca-certificates && update-ca-certificates
COPY ./package.json /app

COPY . .

CMD npm install && npm run dev
