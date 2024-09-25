# Start from the official Go image
FROM golang:1.20-alpine

# Install necessary build tools
RUN apk add --no-cache git gcc musl-dev

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o /app/main .

# Command to run the executable
CMD ["./main"]