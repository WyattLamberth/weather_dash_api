# syntax=docker/dockerfile:1

# Build stage: compile the Go binary.
FROM golang:1.24 AS builder
WORKDIR /app

# Copy go.mod and go.sum to leverage Docker cache.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code from the src directory.
COPY src/ ./src/

# Change to the src directory and build the binary.
WORKDIR /app/src
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/api .

# Run stage: use a minimal image.
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copy the built binary from the builder stage.
COPY --from=builder /app/api .

# Expose the port the API server listens on.
EXPOSE 8080

# Run the API.
CMD ["./api"]