# Use the official Go image as a parent image.
FROM golang:1.21.3-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy the local package files to the container's workspace.
ADD . /app

# Build the Go app as a static binary.
RUN go build -o /hello-world

# Document that the service listens on port 8080.
EXPOSE 8080

# Run the hello-world command by default when the container starts.
CMD ["/hello-world"]
