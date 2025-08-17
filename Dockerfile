# Use official Go image as builder
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

# Cache deps
COPY go.mod go.sum ./
RUN go mod download

# Copy everything
COPY . .

# Build the full cmd package (not just main.go)
RUN go build -o server ./cmd

# Final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 8080 50051

CMD ["./server"]
