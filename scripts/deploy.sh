#!/bin/bash

GITHUB_SHA=$1
DOCKERHUB_USERNAME=$2
DOCKERHUB_TOKEN=$3
EC2_HOST=$4
SSH_USERNAME=$5

echo "StrictHostKeyChecking no" >> ~/.ssh/config

scp docker-compose.yml $SSH_USERNAME@$EC2_HOST:~/docker-compose.yml

ssh $SSH_USERNAME@$EC2_HOST "
  echo $DOCKERHUB_TOKEN | docker login --username $DOCKERHUB_USERNAME --password-stdin
  docker-compose pull
  docker-compose down
  docker-compose up -d
  docker image prune -f
"