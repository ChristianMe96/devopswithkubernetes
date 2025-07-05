1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create -a 2
    ```

2. Create new deployment with the todo app image
    ```bash
    kubectl apply -f manifests/deployment.yaml
    ```

3. Get the Pod name
    ```bash
    kubectl get pods
    ```

4. Make the Pod available through Port Forwarding
    ```bash
    kubectl port-forward <pod-name> 3003:8080
    ```

5. Open the web page
    ```bash
    https://localhost:3003/
    ```