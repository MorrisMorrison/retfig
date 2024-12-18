# Use the official Go image as the base image
FROM golang:1.22.2

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

RUN go get -u github.com/a-h/templ
RUN templ generate
# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

