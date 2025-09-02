# Start from the official Golang image
FROM golang:1.25-alpine

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o jouster *.go

# Expose port 8080 (change if your server uses a different port)
EXPOSE 3000

# Command to run the server
CMD ["/app/jouster"]
