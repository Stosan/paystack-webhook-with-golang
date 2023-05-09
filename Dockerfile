#########
# FINAL #
#########
#########
# FINAL #
#########

# Use golang:alpine as the base image for building the application
FROM golang:alpine AS build

# Create a work directory to copy the application files
WORKDIR /app

# Copy the entire project into the container's work directory
COPY . .

# Build the application and name the binary 'paystack_webhookapp'
RUN go build -o paystack_webhookapp .

# Use a new Alpine image as the base image
FROM alpine

# Set the working directory to /app
WORKDIR /app

# Copy the binary file 'paystack_webhookapp' from the previous build stage
COPY --from=build /app/paystack_webhookapp .

# COPY .env .
# Expose port 2000 to the outside world
EXPOSE 9000

# Start the application
CMD ["./paystack_webhookapp"]
