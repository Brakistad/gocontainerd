#!/bin/bash
set -u
set -o pipefail

source "/docker/base.sh"

# Wait for the database to be ready

export PGPASSWORD=$POSTGRES_PASSWORD

# pg_isready -d $POSTGRES_DB -h $POSTGRES_HOSTNAME -p 5432 -U $POSTGRES_USER
until pg_isready -d $POSTGRES_DB -h $POSTGRES_HOSTNAME -p 5432 -U $POSTGRES_USER; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 1
done

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