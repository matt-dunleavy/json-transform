# Use the official Golang image as the base image
FROM golang:1.18 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o json-transform main.go

# Start a new stage from scratch
FROM debian:buster

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/json-transform /usr/local/bin/json-transform

# Copy configuration file
COPY config.yaml /app/config.yaml

# Set the Current Working Directory inside the container
WORKDIR /app

# Expose port 8080 to the outside world (if your application listens on a port)
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["json-transform"]
