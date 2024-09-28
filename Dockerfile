# Use the official Golang image to build the application
FROM golang:1.23 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o airport-api .

# Start a new stage from scratch
FROM debian:bullseye-slim

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=build /app/airport-api /app/airport-api

# Expose port 8080 to the outside world
EXPOSE 90900

# Command to run the executable
CMD ["./airport-api"]
