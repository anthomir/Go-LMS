# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory in the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Build the Golang application
RUN go build -o main

# Expose the port your Golang application will listen on
EXPOSE 8000

# Command to run your Golang application
CMD ["./main"]
