#!/usr/bin/env bash
set -euo pipefail

helm template devsecops-secure-nginx-platform helm/devsecops-secure-nginx-platform --namespace devsecops-secure-nginx
