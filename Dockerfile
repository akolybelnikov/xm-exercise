# Build stage
FROM golang:1.23.2 AS build
WORKDIR /go/src/github.com/akolybelnikov/xm-exercise
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app

# Production stage
FROM alpine:3.18
COPY --from=build /app /app
COPY --from=build /go/src/github.com/akolybelnikov/xm-exercise/config/config.yaml /config.yaml
ENTRYPOINT ["/app/cmd/main/main.go"]

EXPOSE 8080