#!/bin/bash

# Default port (if not provided)
PORT=${1:-8088}

# Ensure script runs from project root
cd "$(dirname "$0")"/..

echo "🚀 Stopping any running CockroachDB instances..."
pkill -9 cockroach || true

echo "🚀 Cleaning up CockroachDB temporary files..."
rm -rf cockroach-data

echo "🚀 Starting CockroachDB on a different HTTP port..."
cockroach start-single-node --insecure --listen-addr=localhost --http-addr=localhost:8081 --background

# Wait a few seconds to ensure CockroachDB starts properly
sleep 3

echo "🚀 Starting API Server on port $PORT..."
go run cmd/server.go -port=$PORT
