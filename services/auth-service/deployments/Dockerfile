# Use golang base image
FROM golang:1.22.1-alpine

# Set working directory
WORKDIR /app

# Copy the module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Install Air for hot reloading - ensure to use a fixed version to have a predictable build
RUN go install github.com/cosmtrek/air@latest

# Copy the air configuration file into the container
COPY .air.toml /app/.air.toml

# Copy the rest of the application code
COPY . .

# Build the application for production use
# Adjustments for Air not needed here since Air is used in development
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-service .

# Copy the entrypoint script into the image and make it executable
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Expose port 8080 for the application
EXPOSE 8083

# Use the entrypoint script to decide what to run
ENTRYPOINT ["/entrypoint.sh"]
