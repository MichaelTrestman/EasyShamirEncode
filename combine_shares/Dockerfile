FROM golang:1.22 AS builder

WORKDIR /app

COPY combine_shares.go ./

RUN go mod init combine-shares && \
    go get github.com/SSSaaS/sssa-golang && \
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o combine-shares combine_shares.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/combine-shares .

ENTRYPOINT ["/app/combine-shares"]
