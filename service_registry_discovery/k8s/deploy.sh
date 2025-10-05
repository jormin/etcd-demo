#!/bin/bash

kubectl delete -f namespace.yaml

cd deploy
kubectl delete -f .
kubectl apply -f .