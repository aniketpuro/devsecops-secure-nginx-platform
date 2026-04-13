#!/usr/bin/env bash
set -euo pipefail

VAULT_POD="$(kubectl get pods -n vault -l app.kubernetes.io/name=vault -o jsonpath='{.items[0].metadata.name}')"

if [[ -z "${VAULT_POD}" ]]; then
  echo "Vault pod not found in namespace 'vault'." >&2
  exit 1
fi

kubectl exec -n vault "${VAULT_POD}" -- sh -c 'vault auth enable kubernetes || true'
kubectl exec -n vault "${VAULT_POD}" -- sh -c 'vault secrets enable -path=secret kv-v2 || true'

kubectl exec -n vault "${VAULT_POD}" -- sh -c \
  'vault write auth/kubernetes/config \
    token_reviewer_jwt="$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)" \
    kubernetes_host="https://${KUBERNETES_PORT_443_TCP_ADDR}:443" \
    kubernetes_ca_cert=@/var/run/secrets/kubernetes.io/serviceaccount/ca.crt'

kubectl cp security/vault/vault-policy.hcl "vault/${VAULT_POD}:/tmp/vault-policy.hcl"
kubectl exec -n vault "${VAULT_POD}" -- sh -c 'vault policy write devsecops-secure-nginx /tmp/vault-policy.hcl'

kubectl exec -n vault "${VAULT_POD}" -- sh -c \
  'vault write auth/kubernetes/role/devsecops-secure-nginx \
    bound_service_account_names=secure-nginx-sa \
    bound_service_account_namespaces=devsecops-secure-nginx \
    policies=devsecops-secure-nginx \
    ttl=24h'

kubectl exec -n vault "${VAULT_POD}" -- sh -c \
  'vault kv put secret/devsecops-secure-nginx-platform/runtime \
    APP_ENV=production \
    CONTACT_EMAIL=platform-security@example.com'

echo "Vault configured for devsecops-secure-nginx-platform."
