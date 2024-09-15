# Start from a base Go image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go app
RUN go build -o go-backend-app

# Expose the port the app runs on (as defined in your Fiber app)
EXPOSE 8000

# Command to run the application
CMD ["/app/go-backend-app"]
