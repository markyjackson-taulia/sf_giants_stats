# Use the golang Docker image as the base image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go get -d -v ./...
RUN go install -v ./...

# Expose port 8080 for incoming traffic
EXPOSE 8080

# Add an entrypoint script
COPY /scripts/entrypoint.go /scripts/entrypoint.go
RUN chmod +x /scripts/entrypoint.go
ENTRYPOINT ["/scripts/entrypoint.go"]
