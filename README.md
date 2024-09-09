# go-secure-api

## Project Scope
This project focuses on building a small, cloud-native API service using Go, deployed on Kubernetes, with OAuth2 for authentication. The goal is to practice developing a scalable, secure API in a distributed cloud environment with modern security protocols and infrastructure automation.

### Set Up and Build API Service in Go

1. Set Up Development Environment:
* Install Go, Docker, and kubectl for Kubernetes.
* Install kind for local Kubernetes deployment.

2. Develop the API Service in Go:
* Build a simple REST API service in Golang to store music song info.
* Implement routing, data handling, and basic error handling.
* TODO: Connect to a cloud database like MongoDB or PostgreSQL.

3. API Security with OAuth2:
* Integrate OAuth2 for authentication, using Google OAuth 2.0 playground to get access tokens that can be verified with Google Token Info endpoint.
 - Get OAuth 2.0 access token from Google OAuth 2.0 playground: https://developers.google.com/oauthplayground/
 - Example curl request to verify the authentication middleware
 ```
 curl -X POST localhost:8080/songs -i -H "Content-Type: application/json" -H "Authorization: Bearer <access_token>" -d '{"title": "Monday","artist": "Leah Dou","rating": 4.5}'
 ```
* Protect specific API endpoints so only authenticated users can access them.

### Containerization and Kubernetes Deployment
1. Dockerize the Go Application:
* Write a Dockerfile to containerize the Go API service.
* Build and run the Docker container locally to ensure it works correctly.

2. Deploy to Kubernetes:
* Set up a kind cluster.
* Create Kubernetes manifests for deploying the Go API service.

3. Auto-scale deployment:
* Set up Horizontal Pod Autoscaler (HPA) to automatically scale the number of pods based on CPU usage.
```
# Install metrics-server, and then add "- --kubelet-insecure-tls" to `.containers.args`
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

### Testing and Documentation
1. Add unit tests for all the handlers package using Golang's testing package.
```
# To run the handlers tests
go test -v ./handlers
```

2. Write a brief README to explain the steps of this project.