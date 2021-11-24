# syntax=docker/dockerfile:1

FROM golang:1.17.2-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY /certs ./certs
COPY /pages ./pages
COPY /templates ./templates

RUN go build -o main .

EXPOSE 9090

CMD [ "/app/main" ]