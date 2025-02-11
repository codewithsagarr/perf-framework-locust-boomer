#!/bin/sh

# Resolve the Locust master's IP address
LOCUST_MASTER_IP=$(getent hosts locust | awk '{ print $1 }')

# Export the IP address as an environment variable
export LOCUST_ADDRESS=$LOCUST_MASTER_IP

# Print the IP address for debugging
echo "Locust Master IP: $LOCUST_ADDRESS"

# Execute the original CMD
exec "$@"