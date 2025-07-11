FROM dunglas/frankenphp:php8.4

# Set working directory
WORKDIR /app

# Copy configuration files first (better layer caching)
COPY Caddyfile ./

# Copy PHP application files
COPY index.php ./
COPY README.md ./

# Expose the default port
EXPOSE 8080

# Start FrankenPHP directly with the Caddyfile
CMD ["frankenphp", "run", "--config", "/app/Caddyfile"]