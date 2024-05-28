#!/bin/bash

# Function to check if the Pulsar broker is ready
function check_broker_ready {
    broker_container_name="broker"
    broker_container_ip=$(docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' "$broker_container_name")

    ping_response=$(docker exec "$broker_container_name" ping -c 1 "$broker_container_ip")
    if echo "$ping_response" | grep "1 packets transmitted, 1 received"; then
        sleep 5
        echo "Pulsar broker is ready."
    else
        echo "Waiting for Pulsar broker to be ready..."
        sleep 1
        check_broker_ready
    fi
}
# Wait for the broker to be ready

docker compose up -d

check_broker_ready

## Execute pulsar-admin command inside a new shell session in the broker container
docker exec -it broker sh -c "./bin/pulsar-admin topics create persistent://public/default/processing-topic"
docker exec -it broker sh -c "./bin/pulsar-admin topics create persistent://public/default/scaler-topic"

# Add superuser for Pulsar Manager

# Get CSRF Token
CSRF_TOKEN=$(curl http://127.0.0.1:7750/pulsar-manager/csrf-token)

# Create or update user using the CSRF token
curl \
-H "X-XSRF-TOKEN: $CSRF_TOKEN" \
-H "Cookie: XSRF-TOKEN=$CSRF_TOKEN;" \
-H 'Content-Type: application/json' \
-X PUT http://127.0.0.1:7750/pulsar-manager/users/superuser \
-d '{"name": "admin", "password": "apachepulsar", "description": "test", "email": "username@test.org"}'