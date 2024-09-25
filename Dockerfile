# Stage 1: Build the Go binary
FROM golang:1.22-alpine3.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Add this step before RUN go mod download
RUN apk add --no-cache curl && curl -I https://proxy.golang.org

# Download all the dependencies
RUN go mod download

# Copy the rest of the application code
COPY ./cmd ./cmd

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/main.go

# Stage 2: Create a small image to run the binary
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/myapp .

# Expose port 3000 to access the app
EXPOSE 3000

# Command to run the binary
CMD ["./myapp"]