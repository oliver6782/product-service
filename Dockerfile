# Use the latest official Golang image as the base image
FROM golang:latest

# Create a directory for your service
RUN mkdir -p /product-service

# Set the working directory inside the container
WORKDIR /product-service

# Copy go.mod and go.sum to the container
COPY go.mod go.sum ./

# Download all dependencies. This will leverage Docker caching.
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Expose the port that your Go application will run on
EXPOSE 8080

# Command to run the application
CMD ["go", "run", "/product-service/cmd/server/main.go"]
