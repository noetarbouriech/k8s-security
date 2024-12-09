# Kubernetes Security Project

## Requirements

- Kubernetes cluster with LoadBalancer (I use minikube with `minikube tunnel`)
- kubectl

## Install

```sh
cd k8s
kubectl apply -f '*.yaml'
curl $(kubectl get gateway -o jsonpath='{.items[0].status.addresses[0].value}')/aggregator
```
