# Kubernetes Deployment Guide

### 1. Get the Node's IP Address
  ```bash
  kubectl get nodes -o wide
  ```

### 2. Edit the `.env` File
Assuming your ip address is 192.168.x.x. 

- Open the `.env` file and set the following variables:
  ```env
  VITE_API_URL=http://192.168.x.x:30009
  FRONTEND_URL=http://192.168.x.x:30008
  ```

### 3. Build and Push Docker Images

Repeat the following 2 steps for the frontend image as well

1. Build the Docker image:
   ```bash
   cd backend
   docker build -t your-dockerhub-username/backend-image:latest .
   ```

2. Push the image:
   ```bash
   docker push your-dockerhub-username/backend-image:latest
   ```

### 4. Update the Configurations
update both `backend.yaml` and `frontend.yaml` with the new image names.

- on line 19 update the image name:
  ```yaml
    image: your-dockerhub-username/backend-image:latest
  ```

### 5. Apply the configrations file

   ```bash
   cd deploy/Kubernetes
   kubectl apply -f backend.yaml
   kubectl apply -f frontend.yaml
   ```

You can now access secret-note on `http://192.168.x.x:30008`


