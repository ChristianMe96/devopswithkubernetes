FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY . .

# Build the application
# CGO_ENABLED=0 for static binary, GOOS=linux for Linux target
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pingpong-app .

FROM alpine:latest

# Install ca-certificates (useful for HTTPS requests if needed)
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN adduser -D -s /bin/sh appuser

WORKDIR /app

COPY --from=builder /app/pingpong-app .

# Create data directory
RUN mkdir -p /app/data

EXPOSE 8080

# Note: Running as root for simplicity with volume mounts

CMD ["./pingpong-app"]