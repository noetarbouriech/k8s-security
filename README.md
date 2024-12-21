# Kubernetes Security Project

## Requirements

- Kubernetes cluster with LoadBalancer (I use kind with [cloud-provider-kind](https://github.com/kubernetes-sigs/cloud-provider-kinds))
- kubectl
- [Linkerd CLI](https://linkerd.io/2.17/getting-started/#step-1-install-the-cli)

## Install

```sh
# Install Contour, Kyverno and Falco
kubectl apply --server-side -k k8s/base

# Configure Gateway, Policies, etc.
kubectl apply -k k8s/config

# Install Linkerd
linkerd install --crds | kubectl apply -f -
linkerd install | kubectl apply -f -
linkerd check # Check for the install to be completed

# Deploy the app
kubectl apply -k k8s/app

# Test the deployment (it may take a bit of time)
linkerd -n app check --proxy
curl $(kubectl get gateway -o jsonpath='{.items[0].status.addresses[0].value}')/aggregator
```
