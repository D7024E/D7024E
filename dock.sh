#! /bin/bash
# This script is based on; https://linuxconfig.org/bash-script-flags-usage-with-arguments-examples
while getopts ':usd' OPTION; do
    case "$OPTION" in 
    u)  
        docker-compose up -d --scale kademliaNodes=1 --no-recreate
        docker-compose up -d --no-recreate
        echo "Started docker with default settings"
        ;;
    s)
        echo "Number of nodes after scale:"
        read rescale
        echo "Rescaling number of nodes to $rescale, no recreation"
        docker-compose up -d --scale kademliaNodes=$rescale --no-recreate
        echo "Nodes rescaled to $"
        ;;
    d)
        docker-compose down
        echo "Shutting down docker"
        ;;
    esac
done