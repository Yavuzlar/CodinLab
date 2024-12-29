#Backend
FROM golang:alpine3.20 AS backend
# Copy Backend And Version File
WORKDIR /app
COPY ./backend/. .
COPY VERSION .

ENV TZ=UTC \
    CGO_ENABLED=1 \
    GOOS=linux 
    
RUN VERSION=$(cat VERSION) && \
apk update && apk add --no-cache gcc musl-dev && \
echo "Building version $VERSION" && \
go build -ldflags "-s -w -X 'github.com/Yavuzlar/CodinLab/internal/config.APP_VERSION=$VERSION'" -o codinlab_backend cmd/codeinlab/main.go


#Frontend
FROM node:alpine AS frontend
WORKDIR /app
COPY ./frontend/. .
RUN apk update && apk add ca-certificates && update-ca-certificates
RUN npm install && npm run build


#Final Image
FROM nginx:alpine
WORKDIR /app
RUN apk update && apk add --no-cache docker-cli nodejs npm tzdata ca-certificates && update-ca-certificates
COPY --from=backend /app/codinlab_backend /app/back/codinlab_backend
COPY --from=backend /app/migrations /app/back/migrations
COPY --from=backend /app/object /app/back/object
COPY --from=backend /app/config/prod.yaml /app/back/config/config.yaml
COPY --from=frontend /app/next.config.js /app/front/next.config.js
COPY --from=frontend /app/public /app/front/public
COPY --from=frontend /app/.next /app/front/.next
COPY --from=frontend /app/node_modules /app/front/node_modules
COPY --from=frontend /app/package.json /app/front/package.json
COPY ./docker/nginx/prod.conf /etc/nginx/conf.d/default.conf
CMD ["sh", "-c", "cd back && ./codinlab_backend & sh -c 'cd front && npm run start' & nginx -g 'daemon off;'"]
