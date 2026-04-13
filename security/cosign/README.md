# Cosign Integration

The release workflow supports optional image signing with Cosign when these GitHub secrets are configured:

- `COSIGN_PRIVATE_KEY`
- `COSIGN_PASSWORD`

The workflow signs the immutable Docker image tag after a successful push. Verification example:

```bash
cosign verify --key cosign.pub docker.io/acmeplatform/devsecops-secure-nginx-platform:1.0.0-sha12345
```

For keyless signing, replace the private-key flow with OIDC-backed signing and update the workflow trust policy accordingly.
