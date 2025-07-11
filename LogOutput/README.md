# Log Output Application

A web application that generates a random string on startup and serves it via HTTP endpoints.

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
  "timestamp": "2023-12-07 15:30:45",
  "random_string": "a1b2c3d4e5f6789012345678901234567890"
}
```

## Kubernetes Deployment

```bash
# Apply all manifests
kubectl apply -f manifests/

# Check deployment status
kubectl get pods -l app=log-output
kubectl get svc log-output-svc
kubectl get ingress log-output-ingress
```

## Accessing the Application

### Via Ingress (External Access)

With k3d cluster started as:
```bash
k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
```

Access the application at: `http://localhost:8081`

### Via Port Forward (Development)

```bash
kubectl port-forward svc/log-output-svc 8080:8080
```

Then access the application at: `http://localhost:8080`

## Technology Stack

- **Runtime**: FrankenPHP (PHP 8.4)
- **Web Server**: Caddy (via FrankenPHP)
- **Container**: Docker
- **Orchestration**: Kubernetes
- **Ingress**: Traefik (k3d default)