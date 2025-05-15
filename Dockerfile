# Use Golang official image to build the Go binary
FROM golang:1.22-alpine as builder

WORKDIR /src
COPY . .

# Build the Go binary
RUN go build -o litehat ./cmd/litehat

# Final image
FROM alpine:latest

WORKDIR /root/
COPY --from=builder /src/litehat .

ENTRYPOINT ["/root/litehat"]
