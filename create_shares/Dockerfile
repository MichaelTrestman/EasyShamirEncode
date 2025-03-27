FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod tidy && go get github.com/SSSaaS/sssa-golang

COPY create_shares.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o create-shares create_shares.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/create-shares .

ENTRYPOINT ["/app/create-shares"]
