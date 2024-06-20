# Use the official Golang image as the builder
FROM golang:1.18 as builder

# Set the working directory
WORKDIR /app

# Copy the Go modules files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application for multiple architectures
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# Use a minimal Docker image as the base image
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the built Go application from the builder stage
COPY --from=builder /app/main .

# Expose the port the service will run on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
