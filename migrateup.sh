#!/bin/bash
set -e  # Exit on error

# Run database migrations using goose
export GOOSE_DRIVER=turso
export GOOSE_DBSTRING="$DATABASE_URL"
goose -dir migrations up
