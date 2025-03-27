# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

# Add module files and initialize go.mod
COPY go.mod ./

# Fetch the module if not present yet
RUN go mod tidy && go get github.com/SSSaaS/sssa-golang


# Copy source code
COPY create_shares.go ./

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o create-shares create_shares.go

# Final image
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/create-shares .

ENTRYPOINT ["/app/create-shares"]
