#!/usr/bin/env bash
set -euo pipefail

mkdir -p sbom
IMAGE_NAME="${1:-devsecops-secure-nginx-platform:local}"
syft "${IMAGE_NAME}" -o cyclonedx-json > "sbom/$(echo "${IMAGE_NAME}" | tr '/:' '__').cdx.json"
