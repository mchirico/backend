#!/bin/bash
export DOCKER_GIT_CREDENTIALS="$(cat ~/.git-for-docker)"
docker build . --no-cache --build-arg DOCKER_GIT_CREDENTIALS -t backend
docker tag backend acrDevelopment.azurecr.io/backend:latest
docker push acrDevelopment.azurecr.io/backend:latest

# az acr repository delete --name acrDevelopment  --image backend:v5
