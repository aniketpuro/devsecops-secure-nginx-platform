# Security Policy

## Supported Versions

| Version | Supported          | Notes |
|---------|--------------------|-------|
| main    | :white_check_mark: | Active development |
| v1.x    | :white_check_mark: | Latest stable release |
| < 1.0   | :x:                | Not supported |

---

## Reporting a Vulnerability

We take security seriously. If you discover a security vulnerability in this project, please report it responsibly.

### How to Report

- **Preferred Method**: Use GitHub's **Private Vulnerability Reporting** (enabled for this repository)
- **Alternative**: Send an email to `aniketpurohit17@gmail.com` with subject "**Security Vulnerability Report - devsecops-secure-nginx-platform**"

### What to Include in Report

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Affected component (Dockerfile, Kyverno policy, Helm chart, NGINX config, CI/CD workflow, etc.)
- Suggested fix (if any)

### Our Response Timeline

| Timeframe       | Response |
|-----------------|----------|
| Initial Acknowledgment | Within 48 hours |
| Detailed Analysis & Triage | Within 5 business days |
| Patch Release (if critical) | Within 7-14 days |

### Disclosure Policy

- We follow **responsible disclosure**.
- We will publicly acknowledge reporters (unless you prefer to stay anonymous).
- We will not take legal action against researchers who follow these guidelines.

---

## Security Features of This Project

This platform is built with strong DevSecOps principles:
- Image signing with Cosign (mandatory)
- SBOM generation
- Trivy, Semgrep & Gitleaks scanning in CI/CD
- Kyverno policies
- Vault secret management
- Read-only filesystem + dropped capabilities
- Network policies

---

## Contact

- **Maintainer**: Aniket Purohit (@aniketpuro)
- For urgent issues: Open a private vulnerability report on GitHub

Thank you for helping keep this project secure! 🔒
