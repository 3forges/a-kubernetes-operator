#!/bin/bash
set -o errexit

# We need to build the executable outside of contianer for now

go mod tidy
GOOS=linux GOARCH=amd64 go build -o operator

# then we build the container image
docker rmi pesto-operator:0.0.3
docker build -t pesto-operator:0.0.3 .

docker tag pesto-operator:0.0.3 pesto-operator:latest

# then we load the built image to the fresh kind cluster
export KND_CLUSTER_NAME=${KND_CLUSTER_NAME:-"koperator-play"}
kind load docker-image -n "${KND_CLUSTER_NAME}" pesto-operator:0.0.3
kind load docker-image -n "${KND_CLUSTER_NAME}" pesto-operator:latest

# and we finally deploy the kubernetes operator

kubectl apply -f ./kubernetes/operator-ns.yaml

kubectl apply -f ./kubernetes/



