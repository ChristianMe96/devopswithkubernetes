# Ping Pong Application

An API that returns a counter called "pong" which increases with each request. The counter is persisted to a shared volume and survives pod restarts.

## Features

- Counter persisted to PersistentVolume at `/app/data/pingpong_count.txt`
- Thread-safe counter operations with mutex
- Atomic file writes to prevent corruption
- Counter survives pod restarts

## Architecture

- Uses PersistentVolumeClaim `log-output-claim` for counter storage
- Shares volume with LogOutput application for counter visibility

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

3. Apply PingPong manifests
    ```bash
    cd PingPong
    kubectl apply -f manifests/
    ```

4. Apply LogOutput manifests (to access via single ingress)
    ```bash
    cd ../LogOutput
    kubectl apply -f manifests/
    ```

5. Access the application via Ingress
    ```bash
    curl http://localhost:8081/pingpong
    ```

## API Response

```json
{
  "pong": 5
}
```

## Docker Image

- **Image**: `chrisme96/dwk-ping-pong:1.11`

## Technology Stack

- **Language**: Go 1.25
- **Container**: Multi-stage Docker builds with Alpine Linux
- **Orchestration**: Kubernetes
- **Storage**: PersistentVolume for counter persistence
