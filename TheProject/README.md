1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create --port 8082:30080@agent:0 -p 8081:80@loadbalancer --agents 2
    ```

2. Create deployment and service (NodePort) with the todo app image
    ```bash
    kubectl apply -f manifests
    ```

5. Open the web page
    ```bash
    http://localhost:8081/
    ```