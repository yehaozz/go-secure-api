# Use the official Golang image as a build stage
FROM golang:alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application for ARM
RUN go build -o api .

# Use a lightweight image for the final deployment
FROM alpine:latest

# Copy the built binary from the builder stage
COPY --from=builder /app/api /app/api

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["/app/api"]
