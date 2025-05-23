1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create -a 2
    ```
    ```
2. Create new deployment with the random-string image
    ```bash
    kubectl apply -f manifests/deployment.yaml
    ```

3. Get the Pod name
    ```bash
    kubectl get pods
    ```

4. Get the logs of the Pod
    ```bash
    kubectl logs -f <pod-name>
    ```