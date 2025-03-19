FROM golang:1.23-bookworm AS base

WORKDIR /blogging-api

COPY go.mod go.sum ./

RUN go mod download

COPY . .
CMD ["go","run","main.go"]