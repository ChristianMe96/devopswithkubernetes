1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create -a 2
    ```

2. Build the random-string image
    ```bash
    docker build -t chrisme96/dwk-log-output:1.01 .
    ```
3. Create new deployment with the random-string image
    ```bash
    kubectl create deployment log-output-dep --image=chrisme96/dwk-log-output:1.01
    ```

4. Get the Pod name
    ```bash
    kubectl get pods
    ```

5. Get the logs of the Pod
    ```bash
    kubectl logs -f <pod-name>
    ```