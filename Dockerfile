# Getting the latest image of go lang 
FROM golang:latest

# Description label
LABEL description="Go lang server for sorting of Arrays"

# Set the working directory inside the container
WORKDIR /app

# Copy the local code into the container
COPY . .

# Build the Go application
RUN go build -o main.go

# Exposing the port of the application
EXPOSE 8000

# Command to run the executable
CMD ["./sort.go"]