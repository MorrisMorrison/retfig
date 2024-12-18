# Use the official Go image as the base image
FROM golang:1.22.2

# Set the working directory inside the container
WORKDIR /app

# Copy the local package files to the container's workspace
COPY . .

# Install the `templ` binary
RUN go install github.com/a-h/templ/cmd/templ@latest

# Ensure the installed binary is in the PATH
ENV PATH="/go/bin:${PATH}"

# Run the `templ generate` command
RUN templ generate

# Build the Go application
RUN go build -o main .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
