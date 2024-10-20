# Build stage
# Use the Golang image for building the application
FROM golang:1.23.2-alpine3.20 AS build

# Install the gcc and musl-dev packages
RUN apk add --no-progress --no-cache gcc musl-dev

WORKDIR /go/src/github.com/akolybelnikov/xm-exercise
# Copy the source code into the container
ADD . .

# Build the application
RUN GOOS=linux go build -tags musl -ldflags '-extldflags "-static"'  -o /app ./cmd

# Production stage
# Use the Alpine image for the production image
FROM alpine:3.18

# Copy the built application into the production image
COPY --from=build /app /app
COPY --from=build /go/src/github.com/akolybelnikov/xm-exercise/config /config

# Set the entrypoint to the application
ENTRYPOINT ["/app"]

# Expose the port that the application listens on
EXPOSE 8080