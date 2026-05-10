#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$SCRIPT_DIR"

SPEC_URL="https://api.zarban.io/assets/openapi.yaml"
SPEC_FILE="$PROJECT_ROOT/api_specs/service.openapi.yaml"
CLIENT_FILE="$PROJECT_ROOT/service/client.go"

echo "Downloading OpenAPI spec from $SPEC_URL..."
curl -sSL "$SPEC_URL" -o "$SPEC_FILE"
echo "Spec saved to $SPEC_FILE"

echo "Generating service client..."
oapi-codegen -package service -generate "types,client" "$SPEC_FILE" > "$CLIENT_FILE"
echo "Client generated at $CLIENT_FILE"
