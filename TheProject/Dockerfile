FROM dunglas/frankenphp:php8.4-bookworm

# Set working directory
WORKDIR /app

# Copy configuration files first (better layer caching)
COPY Caddyfile ./

# Copy PHP application files
COPY index.php ./
COPY README.md ./

# Create a simple startup script that shows the message and starts FrankenPHP
RUN echo '#!/bin/sh' > /startup.sh && \
    echo 'PORT=${PORT:-8080}' >> /startup.sh && \
    echo 'echo "Server started in port $PORT"' >> /startup.sh && \
    echo 'exec frankenphp run --config /app/Caddyfile' >> /startup.sh && \
    chmod +x /startup.sh

# Expose the default port
EXPOSE 8080

# Start FrankenPHP with the startup script
CMD ["/startup.sh"] 