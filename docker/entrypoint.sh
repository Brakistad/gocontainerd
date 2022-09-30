#!/bin/bash
set -u
set -o pipefail

source "/docker/base.sh"

echo "Starting the application..."


go mod download
# if file /bin/start_app exists, run it, if not build it and run it
if [ -f /bin/start_app ]; then
    start_app
else
    go build -o start_app
    ./start_app
fi
echo "Application started."