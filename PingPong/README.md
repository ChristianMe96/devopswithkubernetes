# Ping Pong Application

An API that return counter called pong which increases with each request.

## Setup Instructions

1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
    ```

2. Apply all Kubernetes manifests of the PingPong Application
    ```bash
    kubectl apply -f manifests/
    ```

3. Apply all Kubernetes manifests of the LogOutput Application (These should work together for now)
    ```bash
    cd ../LogOutput
   
    kubectl apply -f manifests/
    ```

4. Access the application via Ingress
    ```bash
    curl http://localhost:8081/pingpong
    ```