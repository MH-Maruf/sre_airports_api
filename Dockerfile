# Use the official Golang image to build the application
FROM golang:1.23 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the container to download dependencies
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app (make sure the main.go file is present in the copied files)
RUN go build -o airport-api .

# Use a smaller base image for the runtime environment
FROM debian:bullseye-slim

# Set the working directory in the runtime container
WORKDIR /app

# Copy the pre-built binary from the build stage
COPY --from=build /app/airport-api /app/airport-api

# Expose the port on which the app runs
EXPOSE 9090

# Health check to ensure the app is running
HEALTHCHECK --interval=30s --timeout=5s CMD curl -f http://localhost:9090/ || exit 1

# Run the Go binary when the container starts
CMD ["./airport-api"]
