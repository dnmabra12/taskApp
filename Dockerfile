# Use the official Golang image to create a build artifact
FROM golang:1.17 as builder

# Set the working directory inside the container
WORKDIR /src

# Copy local code to the container image
COPY . .

# Build the application
RUN go build -o taskLoader .

# Use a minimal alpine image for the runtime container
FROM alpine:3.14

# Copy the binary from the build stage to the runtime container
COPY --from=builder /src/taskLoader /taskLoader

# Run the application
