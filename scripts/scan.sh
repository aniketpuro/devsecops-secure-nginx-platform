#!/usr/bin/env bash
set -euo pipefail

if ! command -v semgrep >/dev/null 2>&1; then
  echo "semgrep is required but not installed." >&2
  exit 1
fi

if ! command -v gitleaks >/dev/null 2>&1; then
  echo "gitleaks is required but not installed." >&2
  exit 1
fi

if ! command -v trivy >/dev/null 2>&1; then
  echo "trivy is required but not installed." >&2
  exit 1
fi

echo "Running static tests..."
chmod +x ./scripts/test.sh

echo "Running Semgrep..."
semgrep --config security/semgrep.yml

echo "Running Gitleaks..."
gitleaks detect --config security/gitleaks.toml --source .

echo "Running Trivy filesystem scan..."
trivy fs --severity HIGH,CRITICAL --ignore-unfixed --exit-code 1 .

if command -v docker >/dev/null 2>&1; then
  echo "Building local image for image scanning..."
  docker build -f docker/Dockerfile -t devsecops-secure-nginx-platform:local .

  echo "Running Trivy image scan..."
  trivy image --severity HIGH,CRITICAL --ignore-unfixed --exit-code 1 devsecops-secure-nginx-platform:local
fi

echo "All local scans passed."
