# Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy && go get github.com/SSSaaS/sssa-golang

COPY create_shares.go ./


RUN CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o create-shares create_shares.go
