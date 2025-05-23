# Todo App with FrankenPHP

A PHP-based Todo application using FrankenPHP server with PHP 8.4.

1. Start a k3d cluster with 2 agents
    ```bash
    k3d cluster create -a 2
    ```

2. Build the todo app image
    ```bash
    docker build -t chrisme96/dwk-the-project:1.2 .
    ```

3. Create new deployment with the todo app image
    ```bash
    kubectl create deployment todo-app-dep --image=chrisme96/dwk-the-project:1.2
    ```

4. Get the Pod name
    ```bash
    kubectl get pods
    ```

5. Get the logs of the Pod
    ```bash
    kubectl logs -f <pod-name>
    ```  