# Dockerfile pour déploiement Render
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy go mod files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY main.go ./

# Build the application
RUN go build -o go-svelte-philosophy main.go

# Final stage - minimal runtime
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/go-svelte-philosophy .

# Expose port
EXPOSE 8080

# Run the application
CMD ["./go-svelte-philosophy"]