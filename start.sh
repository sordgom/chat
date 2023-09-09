#!/bin/sh

set -e

echo "Starting Auth server..."
exec go run main.go -server auth