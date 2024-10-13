FROM golang:1.22 AS build

# Set to your project's root directory
WORKDIR /app

# Copy your module files first
COPY go.mod go.sum ./ 

# Download dependencies based on go.mod
RUN go mod download 

COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]