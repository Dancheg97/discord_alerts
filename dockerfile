FROM golang:1.18.0-alpine3.15
WORKDIR /
COPY go.mod .
COPY go.sum .
COPY *.go .
CMD go run .