FROM node:lts-alpine AS build-js

WORKDIR /app

COPY client/package*.json ./

RUN npm install

COPY ./client ./

RUN npm run build 

FROM golang:alpine AS build-go

WORKDIR /app

COPY ./go.* ./

COPY ./*.go ./

COPY ./handler ./handler

COPY ./model ./model

RUN CGO_ENABLED=0 go build -o main main.go router.go

FROM busybox

WORKDIR /app

COPY --from=build-js /app/dist ./dist

COPY --from=build-go /app/main ./main

EXPOSE 8080

CMD ["./main"]
