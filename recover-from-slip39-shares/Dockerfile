# Build the Go binary
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod tidy && go get github.com/SSSaaS/sssa-golang && go get github.com/gavincarr/go-slip39

COPY recover_from_slip39_shares.go ./


RUN CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o recover-from-slip39-shares recover_from_slip39_shares.go
