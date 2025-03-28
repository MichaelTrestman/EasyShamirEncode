FROM golang:1.22 AS builder

WORKDIR /app

COPY create_slip39_shares.go ./

RUN go mod init create-shares && \
    go get github.com/gavincarr/seedkit@latest && \
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o create-slip39-shares create_slip39_shares.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/create-slip39-shares .

ENTRYPOINT ["/app/create-slip39-shares"]
