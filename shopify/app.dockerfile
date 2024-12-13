# Stage 1: Build the Go app
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . ./

# Generate the .pb.go files from .proto files if needed
# Uncomment if you need to run protoc to generate the Go files from proto definitions
# RUN apk add --no-cache protobuf
# RUN protoc --go_out=. --go-grpc_out=. ./shopify/shopify.proto

# Build the Go application
RUN GO111MODULE=on go build -o /app/shopify/shopify ./shopify/cmd/shopify

# Stage 2: Prepare the runtime image
FROM alpine:latest

# Install CA certificates required for HTTPS communication
RUN apk --no-cache add ca-certificates

# Copy the built Go application from the builder stage
COPY --from=builder /app/shopify /app/shopify

# Set the working directory to where the app was copied
WORKDIR /app/shopify

# Make the Go application executable
RUN chmod +x /app/shopify/shopify

# Expose the port your application will run on (adjust as needed)
EXPOSE 8080

# Command to run the application
CMD ["/app/shopify/shopify"]
