#!/bin/bash

# gateway
cd deploy
kubectl delete -f gateway.yaml
kubectl delete -f goods.yaml
kubectl apply -f .