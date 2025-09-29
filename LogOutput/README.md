# Log Output Application

A multi-container application that generates random strings and serves them via HTTP. The application is split into two containers within a single pod:

1. **GenerateLog**: Generates a random string on startup and writes it with timestamps to a shared log file every 5 seconds
2. **ReadLog**: Reads the log file and exposes the content via HTTP JSON API

## Architecture

- Two containers in a single pod (`log-output`) sharing a volume
- `GenerateLog` writes to `/app/logs/random.log`
- `ReadLog` serves the log content via HTTP on port 8080
- Uses `emptyDir` volume for inter-container communication
- Service selects pods by `app: log-output` label and routes to the ReadLog container on port 8080

## Setup Instructions

1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
    ```

2. Apply all Kubernetes manifests
    ```bash
    kubectl apply -f manifests/
    ```

3. Access the application via Ingress
    ```bash
    # Get logs as JSON
    curl http://localhost:8081/
    ```

## API Response

```json
{
  "logs": [
    "[2025-09-29 18:23:34] L9sXQsSlYbBJUuop",
    "[2025-09-29 18:23:39] L9sXQsSlYbBJUuop"
  ],
  "count": 2
}
```

## Docker Images

- **GenerateLog**: `chrisme96/dwk-generate-log:1.10`
- **ReadLog**: `chrisme96/dwk-read-log:1.10`

## Technology Stack

- **Language**: Go 1.25
- **Container**: Multi-stage Docker builds with Alpine Linux
- **Orchestration**: Kubernetes
- **Ingress**: Traefik (k3d default)
- **Volume**: emptyDir for shared storage between containers