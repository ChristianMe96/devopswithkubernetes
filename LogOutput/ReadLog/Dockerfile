FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod .

# Download dependencies (if any)
RUN go mod download

# Copy source code
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o read-log-app .

FROM alpine:latest

# Install ca-certificates (useful for HTTPS requests if needed)
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN adduser -D -s /bin/sh appuser

WORKDIR /app

COPY --from=builder /app/read-log-app .

# Create logs directory and set permissions
RUN mkdir -p /app/logs && \
    chown -R appuser:appuser /app

USER appuser

EXPOSE 8080

CMD ["./read-log-app"]