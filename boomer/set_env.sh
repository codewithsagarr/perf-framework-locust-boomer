#!/bin/bash

# Get the IP address of the locust-master container
LOCUST_MASTER_IP=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' locust)

# Export the IP address as an environment variable
export LOCUST_ADDRESS=$LOCUST_MASTER_IP

# Print the IP address for debugging
echo "Locust Master IP: $LOCUST_ADDRESS"