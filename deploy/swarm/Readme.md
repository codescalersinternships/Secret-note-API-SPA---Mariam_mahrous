# Docker Swarm Deployment Guide

### 1. Get your ip address 
  ```bash
   hostname -I
  ```

### 2. Create the docker swarm
Assuming your ip address is 172.23.x.x. 

  ```bash
   docker swarm init --advertise-addr 172.23.x.x:2377 --listen-addr 172.23.x.x:2377
  ```

### 3. Edit the `.env` File
- Open the `.env` file and set the following variables:
  ```env
  VITE_API_URL=http://172.23.x.x:30009
  FRONTEND_URL=http://172.23.x.x:30008
  ```

### 4. Build Docker Images

1. Build the Docker image:
   ```bash
   cd frontend
   docker build --tag secretnote-fe .
   cd ../backend
   docker build --tag secretnote-be .
   ```


### 5. Deploy using docker swarm

   ```bash
   cd deploy/swarm
   docker stack deploy --compose-file docker-stack.yaml secret-note
   ```

You can now access secret-note on `http://172.24.x.x:3000`


