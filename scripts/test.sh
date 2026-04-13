#!/usr/bin/env bash
set -euo pipefail

test -f app/index.html
test -f app/nginx.conf

grep -q "Secure NGINX Platform" app/index.html
grep -q "Content-Security-Policy" app/index.html
grep -q "listen 8080;" app/nginx.conf
grep -q "server_tokens off;" app/nginx.conf

echo "Static app tests passed."
