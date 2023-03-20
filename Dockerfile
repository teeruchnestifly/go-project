# Use an official Golang runtime as a parent image
FROM golang:1.17

# Set the working directory to /app
WORKDIR /go-rest-api

# Copy the current directory contents into the container at /app
COPY . /go-rest-api

# Download any missing dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
