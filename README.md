# Mahler

This project aims to enable platform teams to give their developers
the simplest possible way to deploy their code. No matter what the
underlying infrastructure looks like.

To that end it draws inspiration from the following projects:
- **Sevalla:** Lets developers easily deploy code (via either nixpacks,
  Dockerfiles or Docker images), databases, storage and static sites.
  Gives a neat overview of how services are interconnected.
- **Railway:** Lets developers deploy code or docker images and manages
  log and metrics collection for each deployed instance.
- **Coder:** Intended for dev environments, it allows the admins
  to define resources using terraform
- **Apollo:** The concept of environment owners that can define what an
  application/artifact needs to provide to advance to theirs. Create update
  plans and integrated change management

## Capabilities

- Let admins define resource templates using terraform
- Have a terraform provider that is responsible for syncing back information to Mahler
- Collect logs and metrics from the deployed applications
- Connect applications using a service mesh and expose using a gateway
- Collect connection information automatically using the service mesh and gateway
- Let Environment owners define their environments and required requisites for apps to deploy there

## Built on

- Go (backend and agent)
- Vue (frontend)
- Terraform [open tofu?]
- OTel for logs, metrics, traces etc.
- Consul for service mesh and gateway
