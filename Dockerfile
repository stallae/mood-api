# Use an official Golang base image
FROM golang:1.17

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the source code to the working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Run the application
