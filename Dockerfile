# Start from the official Golang Alpine image
FROM golang:1.24.2-alpine3.21 as builder

# Set the working directory in the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum* ./

# Download dependencies (if go.sum exists)
RUN go mod download

# Copy the source code
COPY . .

# Build the API application
RUN go build -o /app/api ./cmd/api/main.go

# Build the Queue application
RUN go build -o /app/queue ./cmd/queue/main.go

# Use a minimal alpine image for the final stage
FROM alpine:3.21

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the compiled binaries from the builder stage
COPY --from=builder /app/api /app/api
COPY --from=builder /app/queue /app/queue

# Default command can be overridden in docker-compose.yml
CMD ["/app/api"]