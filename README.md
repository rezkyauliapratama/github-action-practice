# GitHub Action Practice

Hands-on practice for GitHub Actions: container builds, composite actions, and CI/CD patterns.

## Structure

```
.github/workflows/
├── build-image.yml          # Docker image build workflow
└── composite/echo/          # reusable composite action example
docker/
└── Dockerfile.public-ip     # containerized Go webservice
webservice/public-ip/        # Go service: reports public IP + brute-force checker
```

## Workflows

- **build-image** — builds and pushes a Docker image on push/PR
- **composite/echo** — demonstrates composite action reuse

## Quick start

```bash
cd webservice/public-ip
go run main.go
```

## Status

Learning repo — patterns extracted here get reused in real projects.
