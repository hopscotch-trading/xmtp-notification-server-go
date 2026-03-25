# Design: Better Handle Buf GitHub Action in Forks

**Issue:** https://github.com/xmtp/example-notification-server-go/issues/77

## Problem

The current `buf.yml` workflow uses `BUF_TOKEN` secret for both setup and push steps. In forks, this secret is unavailable, causing the workflow to fail on PRs from forks.

## Requirements (EARS)

1. **When** a pull request is opened or updated, **the system shall** run `buf lint` and `buf breaking` on the `proto` directory to validate schema correctness.
2. **When** a pull request validation fails lint or breaking checks, **the system shall** fail the workflow run.
3. **When** code is pushed to the `main` branch, **the system shall** push the proto schema to the Buf Schema Registry (BSR).
4. **When** the `BUF_TOKEN` secret is unavailable during a push to `main`, **the system shall** skip the BSR push step without failing the workflow.

## Design

Split the single `buf` job into two jobs:

### Job: `validate` (runs on PRs)
- Checkout code
- Install `buf` CLI (no API token needed)
- Run `buf lint proto`
- Run `buf breaking proto --against 'https://github.com/xmtp/example-notification-server-go.git#branch=main,subdir=proto'`

### Job: `push` (runs on push to main only)
- Conditional on `github.ref == 'refs/heads/main'` and `secrets.BUF_TOKEN` being available
- Checkout code
- Install `buf` CLI with token
- Push to BSR using `buf-push-action`

## Testing & Validation

- Verify workflow YAML is valid
- Confirm lint/breaking steps don't require authentication
- Confirm push step only runs on main with token available
