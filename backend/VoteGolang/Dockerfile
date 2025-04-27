# Use newer Golang image
FROM golang:1.23.4

# Set environment variables
ENV GO111MODULE=on

# Set working directory inside container
WORKDIR /app

# Copy go.mod and go.sum first
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o main ./cmd/main.go

# Expose the port you want
EXPOSE 8080

# Command to run
CMD ["./main"]
