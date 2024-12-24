#!/bin/bash

# Check if the username is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <username>"
  exit 1
fi

USERNAME=$1

sudo yum update -y
sudo amazon-linux-extras install docker -y
sudo service docker start
sudo chkconfig docker on
sudo usermod -a -G docker $USERNAME

sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

sudo reboot

sudo docker --version
sudo docker-compose --version

