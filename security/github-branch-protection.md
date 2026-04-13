# Branch Protection Strategy

Apply the following protection to `main`:

- Require a pull request before merging
- Require at least 2 approvals
- Require review from Code Owners
- Dismiss stale approvals on new commits
- Require conversation resolution before merge
- Require status checks:
  - `Pull Request Validation / validate`
- Require branches to be up to date before merging
- Restrict direct pushes to administrators only if your governance model requires it
- Block force pushes
- Block deletions

Recommended repository settings:

- Default branch: `main`
- Feature branches: `feature/*`, `bugfix/*`, `hotfix/*`
- Merge strategy: squash or rebase only
- Signed commits encouraged or required
