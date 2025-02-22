# syntax=docker/dockerfile:1
# Build stage: use an official Go image to build your binary.
FROM golang:1.20 AS builder
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache for dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code.
COPY . .

# Build the binary. Set CGO_ENABLED=0 for a static binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o api .

# Run stage: use a lightweight image.
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app

# Copy the binary from the builder stage.
COPY --from=builder /app/api .

# Expose the port your API listens on (adjust if necessary).
EXPOSE 8080

# Command to run the API.
CMD ["./api"]