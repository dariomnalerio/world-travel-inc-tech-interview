#!/bin/bash

GITHUB_SHA=$1
DOCKERHUB_USERNAME=$2

# Build and push client
docker build -t $DOCKERHUB_USERNAME/client:$GITHUB_SHA ./client
docker push $DOCKERHUB_USERNAME/client:$GITHUB_SHA
docker tag $DOCKERHUB_USERNAME/client:$GITHUB_SHA $DOCKERHUB_USERNAME/client:latest
docker push $DOCKERHUB_USERNAME/client:latest

# Build and push server
docker build -t $DOCKERHUB_USERNAME/server:$GITHUB_SHA ./server
docker push $DOCKERHUB_USERNAME/server:$GITHUB_SHA
docker tag $DOCKERHUB_USERNAME/server:$GITHUB_SHA $DOCKERHUB_USERNAME/server:latest
docker push $DOCKERHUB_USERNAME/server:latest