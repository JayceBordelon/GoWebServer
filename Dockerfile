FROM golang:1.24.1-alpine

ENV GO111MODULE=on

WORKDIR /app


RUN apk add --no-cache git


COPY go.mod go.sum ./
RUN go mod download


COPY . .


RUN go build -o app ./cmd/server

EXPOSE 8080

ENTRYPOINT ["./app"]
