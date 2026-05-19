# Go Architecture - task runner

set dotenv-load := false

# List all available recipes
default:
    @just --list

# ⎯⎯⎯⎯⎯ Local development ⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯

# Run the app locally
run:
    go run ./cmd/main/

# ⎯⎯⎯⎯⎯ Docker Compose ⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯

# Starts everything in Docker with the local definitions
local-up:
    docker compose -f .docker/compose/local.yml up -d

# Stops the Docker local stack
local-down:
    docker compose -f .docker/compose/local.yml down
