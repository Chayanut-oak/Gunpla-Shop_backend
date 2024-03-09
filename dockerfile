# Use the official golang image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the entire project to the container's workspace
COPY . .

# Download all the dependencies
RUN go mod download

COPY ./cmd/myapps/.env .
# Build the Go app
RUN go build -o gunpla-store ./cmd/myapps

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
CMD ["./gunpla-store"]