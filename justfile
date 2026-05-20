# Go Architecture - task runner

# List all available recipes
default:
    @just --list

# ⎯⎯⎯⎯⎯ Local development ⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯

# Run the app locally
run:
    go run ./cmd/main/

# Runs the ddbb locally with docker
db-up:
    docker compose -f .docker/compose/local.yml up postgres pgadmin -d

# Stops the ddbb
db-down:
    docker compose -f .docker/compose/local.yml down postgres pgadmin

# Removes all data from local ddbb
db-clean:
    docker compose -f .docker/compose/local.yml down postgres pgadmin -v

# ⎯⎯⎯⎯⎯ Docker Compose ⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯⎯

# Starts everything in Docker with the local definitions
local-up:
    docker compose -f .docker/compose/local.yml up -d

# Stops the Docker local stack
local-down:
    docker compose -f .docker/compose/local.yml down
