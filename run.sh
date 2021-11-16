#!/usr/bin/env sh

docker build -t kvmayer/go-catalogs:latest .
docker push kvmayer/go-catalogs:latest
kubectl delete -f deploy.yaml
kubectl apply -f deploy.yaml