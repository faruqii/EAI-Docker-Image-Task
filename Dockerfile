FROM golang:1.18.0-alpine

# Set the working directory to /app
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download Go modules
RUN go mod download

# Build the Go app
RUN go build ./cmd/main.go

# Expose port 3000 for the application
EXPOSE 3000

# Start the application
CMD ["./main"]
