#!/bin/bash

# Check if the username is provided as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <username>"
  exit 1
fi

USERNAME=$1

echo "Installing Docker for $USERNAME..."

sudo yum update -y
sudo yum install -y docker

# Create docker group if it doesn't exist
sudo groupadd -f docker

# Start and enable Docker service
sudo systemctl start docker
sudo systemctl enable docker

# Add user to docker group
sudo usermod -aG docker $USERNAME

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

# Verify installations
if command -v docker &> /dev/null; then
    echo "Docker installed successfully"
    sudo docker --version
else
    echo "Docker installation failed"
    exit 1
fi

if command -v docker-compose &> /dev/null; then
    echo "Docker Compose installed successfully"
    sudo docker-compose --version
else
    echo "Docker Compose installation failed"
    exit 1
fi

echo "Installation complete!"