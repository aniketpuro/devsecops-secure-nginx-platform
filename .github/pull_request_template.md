## Summary

- Describe the change
- Link the related issue or risk
- Explain the rollback plan

## Security Checklist

- [ ] I did not hardcode secrets
- [ ] I assessed RBAC, network policy, and runtime impact
- [ ] I reviewed CI/CD scan results
- [ ] I documented any operational follow-up

## Validation

- [ ] `make test`
- [ ] `make scan`
- [ ] Helm render or Kubernetes manifest review completed
