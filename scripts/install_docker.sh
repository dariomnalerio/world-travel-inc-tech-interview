#!/bin/bash

# Check if the username is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <username>"
  exit 1
fi

USERNAME=$1

echo "Installing Docker for $USERNAME..."

sudo yum update -y
sudo yum install -y docker python3-devel libcrypt-devel

# Create docker group if it doesn't exist
sudo groupadd -f docker

# Start and enable Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Set correct permissions
sudo chmod 666 /var/run/docker.sock
sudo usermod -aG docker $USERNAME

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose