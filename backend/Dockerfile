FROM golang:1.22-alpine AS build

RUN apk add build-base

WORKDIR /server

COPY .env .env

COPY . .
RUN go mod download

RUN  CGO_ENABLED=1 GOOS=linux go build -o /myserver ./cmd

FROM alpine:edge

WORKDIR /server

COPY .env .env
COPY --from=build /myserver .

EXPOSE ${PORT}

ENTRYPOINT ["./myserver"]