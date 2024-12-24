#!/bin/bash

GITHUB_SHA=$1
DOCKERHUB_USERNAME=$2
DOCKERHUB_TOKEN=$3
EC2_HOST=$4
SSH_USERNAME=$5

mkdir -p ~/.ssh

ssh-keyscan -H $EC2_HOST >> ~/.ssh/known_hosts

echo "StrictHostKeyChecking no" >> ~/.ssh/config

scp docker-compose.yml $SSH_USERNAME@$EC2_HOST:~/docker-compose.yml
scp scripts/install_docker.sh $SSH_USERNAME@$EC2_HOST:~/install_docker.sh

ssh $SSH_USERNAME@$EC2_HOST "
  chmod +x ~/install_docker.sh
  ~/install_docker.sh $SSH_USERNAME
  echo $DOCKERHUB_TOKEN | docker login --username $DOCKERHUB_USERNAME --password-stdin
  docker-compose pull
  docker-compose down
  docker-compose --profile prod up -d
  docker image prune -f
"