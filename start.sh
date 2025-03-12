#!/bin/bash

# # Move to the directory where the script is located
# cd "$(dirname "$0")"

# # Move to the project root if `start.sh` is inside a subfolder
# if [ "$(basename "$(pwd)")" != "stockwise" ]; then
#     cd ..
# fi

# echo "📂 Running script from: $(pwd)"

# echo "🚀 Stopping any running CockroachDB instances..."
# pkill -9 cockroach || true

# echo "🚀 Starting CockroachDB on a different HTTP port..."
# nohup cockroach start-single-node --insecure --listen-addr=localhost --http-addr=localhost:8081 > cockroach.log 2>&1 &

# # Wait until CockroachDB is ready
# echo "⏳ Waiting for CockroachDB to start..."
# sleep 5

# # Check if CockroachDB is running
# if ! ps aux | grep "[c]ockroach" > /dev/null; then
#     echo "❌ CockroachDB failed to start!"
#     cat cockroach.log
#     exit 1
# fi

# echo "✅ CockroachDB is running!"

# # Verify that `cmd/main.go` exists
# if [ ! -f "cmd/main.go" ]; then
#     echo "❌ ERROR: cmd/main.go not found in $(pwd)!"
#     exit 1
# fi

echo "🚀 Starting API Server..."
go run cmd/main.go

# cockroach start-single-node --insecure --listen-addr=localhost --http-addr=localhost:8081 --store=cockroach-data
