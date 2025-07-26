# Use official Go image as builder
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Install git (needed for go get) and other tools
RUN apk add --no-cache git

# Copy go.mod and go.sum and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN go build -o server ./cmd/main.go

# Final stage: minimal runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

# Expose your app port
EXPOSE 8080

CMD ["./server"]
