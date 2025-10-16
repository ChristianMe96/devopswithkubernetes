# Log Output Application

A multi-container application that generates random strings and serves them via HTTP along with the PingPong counter. The application is split into two containers within a single pod:

1. **GenerateLog**: Generates a random string on startup and writes it with timestamps to a shared log file every 5 seconds
2. **ReadLog**: Reads the log file and PingPong counter, exposes both via HTTP JSON API

## Architecture

- Two containers in a single pod (`log-output`) sharing volumes
- `GenerateLog` writes to `/app/logs/random.log` (emptyDir volume)
- `ReadLog` serves the log content via HTTP on port 8080
- `ReadLog` also reads PingPong counter from `/app/data/pingpong_count.txt` (PersistentVolume)
- Uses `emptyDir` volume for inter-container communication
- Uses PersistentVolume to share data with PingPong application
- Service selects pods by `app: log-output` label and routes to the ReadLog container on port 8080

## Setup Instructions

1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
    ```

2. Apply persistent volume infrastructure (cluster admin task)
    ```bash
    cd /Users/christian.meinhard/Projects/github/devopswithkubernetes
    kubectl apply -f persistent-volumes/
    ```

3. Apply PingPong manifests (for counter tracking)
    ```bash
    cd PingPong
    kubectl apply -f manifests/
    ```

4. Apply LogOutput manifests
    ```bash
    cd ../LogOutput
    kubectl apply -f manifests/
    ```

5. Access the applications
    ```bash
    # Get logs with pingpong count as JSON
    curl http://localhost:8081/
    
    # Increment pingpong counter
    curl http://localhost:8081/pingpong
    ```

## API Response

```json
{
  "logs": [
    "[2025-10-16 20:23:34] L9sXQsSlYbBJUuop",
    "[2025-10-16 20:23:39] L9sXQsSlYbBJUuop"
  ],
  "count": 2,
  "pingpongs": 5
}
```

The `pingpongs` field shows the current PingPong counter value from the shared persistent volume.

## Docker Images

- **GenerateLog**: `chrisme96/dwk-generate-log:1.11`
- **ReadLog**: `chrisme96/dwk-read-log:1.11`

## Technology Stack

- **Language**: Go 1.25
- **Container**: Multi-stage Docker builds with Alpine Linux
- **Orchestration**: Kubernetes
- **Ingress**: Traefik (k3d default)
- **Volumes**: 
  - emptyDir for shared storage between containers (logs)
  - PersistentVolume for shared storage between applications (counter)
