# backend/Dockerfile

# Build stage
FROM golang:1.25-alpine AS builder

# Install git (needed for some Go modules)
RUN apk add --no-cache git

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build optimized binary (disable CGO, static linking, smaller binary)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o server ./cmd/server

# Final stage: minimal image
FROM alpine:latest

# Add certificates for HTTPS calls
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from build stage
COPY --from=builder /app/server .

# Expose port (optional here, useful for documentation)
EXPOSE 8080

# Run the binary
CMD ["./server"]
