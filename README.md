# gomunkong-backend

A backend repository for `gomunkong` project, written in Golang.



## Structure

The repo mainly has `formatter`, `language_servers`, and `proto` directories. 



### Formatter

Code formatter main API server.



### Language_servers

Python, Go, Java, Javascript formatter servers.



### Proto

protobuf definitions for gRPC communication between containers.



## Deployment

```
minikube start --vm=true
minikube addons enable ingress

minikube service formatter-service
kubectl expose deployment formatter-deployment --type=NodePort --port=8080
```

`formatter-service` will be exposed in NodePort 30007 to communicate with frontend.



