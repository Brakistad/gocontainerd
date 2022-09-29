#!/bin/bash
echo "Starting the application..."

cd /app && \
go mod download && \
go build -o /usr/bin/start_app

echo "Application started."