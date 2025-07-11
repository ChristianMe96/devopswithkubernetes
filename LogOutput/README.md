# Log Output Application

A web application that generates a random string on startup and serves it via HTTP endpoints.

## Setup Instructions

1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
    ```

2. Apply all Kubernetes manifests
    ```bash
    kubectl apply -f manifests/
    ```

3. Verify deployment status
    ```bash
    kubectl get pods -l app=log-output
    kubectl get svc log-output-svc
    kubectl get ingress log-output-ingress
    ```

4. Access the application via Ingress
    ```bash
    # Web interface
    curl http://localhost:8081/
    
    # JSON API endpoint
    curl http://localhost:8081/status
    ```

## Features

- Generates a random string on application startup
- Stores the random string in memory
- Provides a web interface to view the current status
- Exposes a JSON API endpoint for status information
- Accessible via Ingress for external access

## Endpoints

- `/` - Web interface showing current status
- `/status` - JSON API endpoint returning timestamp and random string

## Example Response

```json
{
  "timestamp": "2025-07-11 13:30:45",
  "random_string": "a1b2c3d4e5f6789012345678901234567890"
}
```

## Technology Stack

- **Runtime**: FrankenPHP (PHP 8.4)
- **Web Server**: Caddy (via FrankenPHP)
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Ingress**: Traefik (k3d default)
- **Docker Image**: Available on Docker Hub as `chrisme96/dwk-log-output:1.7`