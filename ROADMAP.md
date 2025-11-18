# Mahler Platform Roadmap

**Vision:** A self-hostable platform that enables small teams to provide developers with a cloud-like experience for managing infrastructure and resources, powered by Terraform and the HashiCorp ecosystem.

**Similar to:** coder.com, but focused on resource provisioning and lifecycle management.

---

## Current State

### Existing Components
- ✅ Basic Go backend structure (Huma v2 API framework)
- ✅ Vue.js frontend with interactive canvas UI
- ✅ Resource interface and K8s Pod implementation example
- ✅ Project and Service domain models
- ✅ Development guidelines (CLAUDE.md)

### Critical Gaps
- ❌ No tests (requires 95% Go, 80% Vue coverage)
- ❌ No persistence layer
- ❌ No CI/CD pipeline
- ❌ No Terraform integration
- ❌ No authentication/authorization
- ❌ No observability stack integration

---

## Roadmap Phases

## Phase 0: Foundation & Infrastructure (Weeks 1-3)

**Goal:** Establish solid foundation with testing, CI/CD, and persistence to support rapid, confident development.

### 0.1 Testing Infrastructure
**Priority:** CRITICAL | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [x] **0.1.1** Set up Go testing framework
  - Create example table-driven tests for existing code
  - Configure coverage reporting (`go test -cover`)
  - Add test helpers and assertions package
  - **Deliverable:** `internal/app/app_test.go` with 95%+ coverage
  - **Tests:** N/A (this creates the testing foundation)

- [x] **0.1.2** Set up Vue/Vitest testing framework
  - Install Vitest and testing utilities (`@vue/test-utils`)
  - Configure `vitest.config.js`
  - Create example component tests
  - **Deliverable:** `web/src/components/ArchitectureCanvas.spec.js` with 80%+ coverage
  - **Tests:** N/A (this creates the testing foundation)

- [x] **0.1.3** Add test coverage enforcement
  - Create `scripts/check-coverage.sh`
  - Configure minimum coverage thresholds
  - **Deliverable:** Script that fails if coverage < 95% (Go) or < 80% (Vue)
  - **Tests:** Run script on current codebase

- [x] **0.1.4** Create mock infrastructure
  - Implement mock repository interfaces
  - Create test fixtures and data builders
  - **Deliverable:** `internal/testutil` package with mocks
  - **Tests:** Use mocks in at least 3 different test files

### 0.2 CI/CD Pipeline
**Priority:** CRITICAL | **Complexity:** Medium | **Duration:** 2 days

#### Tasks:
- [ ] **0.2.1** Create GitHub Actions workflow - Test job
  - Set up Go test job with coverage reporting
  - Set up Vue test job with coverage reporting
  - Upload coverage artifacts
  - **Deliverable:** `.github/workflows/ci.yml` with test jobs
  - **Tests:** Trigger workflow, verify it runs and reports coverage

- [ ] **0.2.2** Create GitHub Actions workflow - Lint job
  - Configure `golangci-lint` with config file
  - Set up ESLint for Vue.js
  - Add formatting checks (gofmt, prettier)
  - **Deliverable:** `.github/workflows/ci.yml` with lint jobs
  - **Tests:** Introduce a lint error, verify workflow fails

- [ ] **0.2.3** Create GitHub Actions workflow - Build job
  - Build Go binary
  - Build Vue frontend
  - Upload build artifacts
  - **Deliverable:** `.github/workflows/ci.yml` with build jobs
  - **Tests:** Download artifacts, verify binary runs

- [ ] **0.2.4** Create release workflow
  - Multi-platform binary builds (Linux, macOS, Windows)
  - Embed Vue frontend in Go binary using `embed`
  - Create GitHub releases with binaries
  - **Deliverable:** `.github/workflows/release.yml`
  - **Tests:** Create a test release, verify binaries work on each platform

- [ ] **0.2.5** Create Docker build workflow
  - Multi-stage Dockerfile (Go build + Vue build + minimal runtime)
  - Push to container registry (GitHub Container Registry)
  - Tag with version and latest
  - **Deliverable:** `Dockerfile` and `.github/workflows/docker.yml`
  - **Tests:** Pull and run container, verify application works

### 0.3 Persistence Layer
**Priority:** CRITICAL | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **0.3.1** Design database schema
  - Define tables: projects, services, resources, resource_templates, users, teams
  - Create migration system (using golang-migrate or similar)
  - Document schema with ER diagram
  - **Deliverable:** `internal/database/migrations/` directory with initial schema
  - **Tests:** Apply migrations up/down successfully

- [ ] **0.3.2** Implement database connection management
  - Create `internal/database/db.go` with connection pool
  - Support SQLite (dev) and PostgreSQL (prod)
  - Add health check endpoint
  - **Deliverable:** Database connection package with configuration
  - **Tests:** Unit tests for connection handling, integration test with real DB

- [ ] **0.3.3** Create repository interfaces
  - Define `ProjectRepository`, `ServiceRepository`, `ResourceTemplateRepository`
  - Follow DIP: interfaces in consumer packages
  - **Deliverable:** `internal/project/repository.go`, etc.
  - **Tests:** N/A (interfaces only, implementations tested separately)

- [ ] **0.3.4** Implement SQL repositories
  - `PostgresProjectRepository` implementing `ProjectRepository`
  - `PostgresServiceRepository` implementing `ServiceRepository`
  - `PostgresResourceTemplateRepository` implementing `ResourceTemplateRepository`
  - **Deliverable:** `internal/database/postgres/` package with implementations
  - **Tests:** Integration tests with test database, 95%+ coverage

- [ ] **0.3.5** Create in-memory repositories for testing
  - Fast, thread-safe in-memory implementations
  - Use for unit tests and local development
  - **Deliverable:** `internal/database/memory/` package
  - **Tests:** Unit tests verifying behavior matches SQL implementation

### 0.4 Configuration Management
**Priority:** HIGH | **Complexity:** Low | **Duration:** 1 day

#### Tasks:
- [ ] **0.4.1** Create configuration package
  - Support environment variables, config files (YAML), and defaults
  - Configuration for: database, server, logging, external services
  - **Deliverable:** `internal/config/config.go`
  - **Tests:** Unit tests for loading from different sources

- [ ] **0.4.2** Add configuration validation
  - Validate required fields
  - Validate formats (URLs, ports, etc.)
  - Provide helpful error messages
  - **Deliverable:** Validation in `config.go`
  - **Tests:** Test with invalid configs, verify error messages

- [ ] **0.4.3** Document configuration options
  - Create `docs/configuration.md`
  - Include example configuration files
  - **Deliverable:** Configuration documentation
  - **Tests:** Manual verification of documentation accuracy

### 0.5 Build & Development Tooling
**Priority:** MEDIUM | **Complexity:** Low | **Duration:** 1 day

#### Tasks:
- [ ] **0.5.1** Create Makefile
  - Targets: `test`, `build`, `run`, `lint`, `coverage`, `migrate`
  - Cross-platform compatibility
  - **Deliverable:** `Makefile`
  - **Tests:** Run each target, verify expected behavior

- [ ] **0.5.2** Create development environment setup
  - Docker Compose for local dependencies (PostgreSQL, etc.)
  - Scripts for initializing development environment
  - **Deliverable:** `docker-compose.dev.yml`, `scripts/setup-dev.sh`
  - **Tests:** Fresh clone, run setup, verify development environment works

- [ ] **0.5.3** Create embedding for Vue frontend
  - Use Go `embed` package to include built Vue app
  - Serve embedded files from Go binary
  - Fallback to filesystem in development
  - **Deliverable:** Update `cmd/main.go` with embedding
  - **Tests:** Build binary, verify frontend is accessible

---

## Phase 1: Core Platform (Weeks 4-7)

**Goal:** Implement complete CRUD operations for projects and services, connect frontend to backend, establish basic resource template system.

### 1.1 Complete Project Management API
**Priority:** CRITICAL | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **1.1.1** Implement project business logic
  - `CreateProject(ctx, name, description)` with validation
  - `GetProject(ctx, id)` with not-found handling
  - `UpdateProject(ctx, id, updates)` with validation
  - `DeleteProject(ctx, id)` with cascade rules
  - `ListProjects(ctx, filter, pagination)`
  - **Deliverable:** `internal/project/service.go` with complete business logic
  - **Tests:** Unit tests with mocks, 95%+ coverage

- [ ] **1.1.2** Implement project API endpoints
  - `POST /api/v1/projects` - Create project
  - `GET /api/v1/projects/{id}` - Get project
  - `PUT /api/v1/projects/{id}` - Update project
  - `DELETE /api/v1/projects/{id}` - Delete project
  - `GET /api/v1/projects` - List projects
  - **Deliverable:** `internal/api/v1/projects.go`
  - **Tests:** API integration tests, verify all endpoints

- [ ] **1.1.3** Add request/response validation
  - Huma schema validation for all requests
  - Proper error responses (400, 404, 500)
  - **Deliverable:** Request/response schemas in API handlers
  - **Tests:** Test invalid inputs, verify error responses

### 1.2 Complete Service Management API
**Priority:** CRITICAL | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **1.2.1** Implement service business logic
  - `CreateService(ctx, projectID, name, resourceTemplateID, config)`
  - `GetService(ctx, id)`
  - `UpdateService(ctx, id, updates)`
  - `DeleteService(ctx, id)`
  - `ListServices(ctx, projectID, filter)`
  - **Deliverable:** `internal/service/service.go`
  - **Tests:** Unit tests with mocks, 95%+ coverage

- [ ] **1.2.2** Implement service API endpoints
  - `POST /api/v1/projects/{projectID}/services`
  - `GET /api/v1/services/{id}`
  - `PUT /api/v1/services/{id}`
  - `DELETE /api/v1/services/{id}`
  - `GET /api/v1/projects/{projectID}/services`
  - **Deliverable:** `internal/api/v1/services.go`
  - **Tests:** API integration tests

- [ ] **1.2.3** Add service-resource template association
  - Link services to resource templates
  - Validate configuration against template schema
  - **Deliverable:** Validation logic in service package
  - **Tests:** Test with valid/invalid configs

### 1.3 Resource Template System
**Priority:** CRITICAL | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **1.3.1** Design resource template schema
  - Template metadata (name, description, version, icon)
  - Input schema (JSON Schema for configuration)
  - Output schema (what the resource provides)
  - Pricing information
  - **Deliverable:** `internal/resource/template/template.go`
  - **Tests:** Create templates, validate schemas

- [ ] **1.3.2** Implement template repository
  - CRUD operations for templates
  - Version management
  - Template categorization/tagging
  - **Deliverable:** Template repository implementation
  - **Tests:** Repository tests with 95%+ coverage

- [ ] **1.3.3** Implement template validation
  - Validate input configs against JSON Schema
  - Validate template completeness
  - **Deliverable:** `internal/resource/template/validator.go`
  - **Tests:** Test with various valid/invalid templates

- [ ] **1.3.4** Create template API endpoints
  - `GET /api/v1/templates` - List available templates
  - `GET /api/v1/templates/{id}` - Get template details
  - `POST /api/v1/templates` - Create template (admin)
  - `PUT /api/v1/templates/{id}` - Update template (admin)
  - `DELETE /api/v1/templates/{id}` - Delete template (admin)
  - **Deliverable:** `internal/api/v1/templates.go`
  - **Tests:** API integration tests

- [ ] **1.3.5** Create built-in resource templates
  - Kubernetes Pod (enhance existing)
  - Kubernetes Deployment
  - Kubernetes Service
  - Database (PostgreSQL, MySQL)
  - Redis
  - **Deliverable:** Seed data with common templates
  - **Tests:** Verify templates validate correctly

### 1.4 Frontend-Backend Integration
**Priority:** CRITICAL | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **1.4.1** Create API client in Vue
  - TypeScript/JavaScript client for all endpoints
  - Error handling and loading states
  - **Deliverable:** `web/src/api/client.js`
  - **Tests:** Mock API tests in Vitest

- [ ] **1.4.2** Implement state management (Pinia)
  - Project store
  - Service store
  - Template store
  - **Deliverable:** `web/src/stores/` directory
  - **Tests:** Store tests with mocked API

- [ ] **1.4.3** Connect ArchitectureCanvas to real data
  - Load projects and services from API
  - Create/update/delete operations
  - Real-time updates
  - **Deliverable:** Updated `ArchitectureCanvas.vue`
  - **Tests:** Component tests with mocked stores

- [ ] **1.4.4** Create project management UI
  - Project list view
  - Project creation modal
  - Project details view
  - **Deliverable:** New Vue components
  - **Tests:** Component tests, 80%+ coverage

- [ ] **1.4.5** Create resource template browser
  - Template catalog view
  - Template details with schema
  - Add service from template
  - **Deliverable:** Template browser components
  - **Tests:** Component tests

### 1.5 Logging and Monitoring Foundation
**Priority:** MEDIUM | **Complexity:** Low | **Duration:** 2 days

#### Tasks:
- [ ] **1.5.1** Implement structured logging
  - Consistent log format (JSON in production)
  - Log levels and contexts
  - Request ID tracking
  - **Deliverable:** `internal/logging/logger.go`
  - **Tests:** Verify log output format

- [ ] **1.5.2** Add request logging middleware
  - Log all HTTP requests with duration
  - Log errors with stack traces
  - **Deliverable:** Huma middleware for logging
  - **Tests:** Verify logs for sample requests

- [ ] **1.5.3** Add basic metrics endpoints
  - Health check (`/health`)
  - Readiness check (`/ready`)
  - Metrics endpoint (`/metrics`) for Prometheus
  - **Deliverable:** Health and metrics endpoints
  - **Tests:** Test each endpoint

---

## Phase 2: Multi-tenancy & Access Control (Weeks 8-10)

**Goal:** Add user authentication, role-based access control, and team/organization management.

### 2.1 User Authentication
**Priority:** CRITICAL | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **2.1.1** Design user and auth schema
  - Users table (id, email, hashed_password, role, created_at)
  - Sessions/tokens table
  - Choose auth strategy (JWT vs sessions)
  - **Deliverable:** Database migrations for auth
  - **Tests:** Apply migrations

- [ ] **2.1.2** Implement authentication service
  - User registration
  - Login (password verification)
  - Logout
  - Password reset flow
  - **Deliverable:** `internal/auth/service.go`
  - **Tests:** Unit tests with mocks, 95%+ coverage

- [ ] **2.1.3** Implement JWT token management
  - Token generation and signing
  - Token validation and parsing
  - Refresh token flow
  - **Deliverable:** `internal/auth/jwt.go`
  - **Tests:** Test token lifecycle

- [ ] **2.1.4** Create authentication API endpoints
  - `POST /api/v1/auth/register`
  - `POST /api/v1/auth/login`
  - `POST /api/v1/auth/logout`
  - `POST /api/v1/auth/refresh`
  - `POST /api/v1/auth/forgot-password`
  - **Deliverable:** `internal/api/v1/auth.go`
  - **Tests:** API integration tests

- [ ] **2.1.5** Add authentication middleware
  - Extract and validate JWT from requests
  - Inject user context
  - Handle unauthenticated requests (401)
  - **Deliverable:** Huma middleware for authentication
  - **Tests:** Test with valid/invalid/missing tokens

### 2.2 Authorization & RBAC
**Priority:** CRITICAL | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **2.2.1** Design RBAC schema
  - Roles: SuperAdmin, Admin, Developer, Viewer
  - Permissions: resource-based (project.*, service.*, template.*)
  - Role-permission associations
  - **Deliverable:** Database migrations for RBAC
  - **Tests:** Apply migrations

- [ ] **2.2.2** Implement authorization service
  - Check user permissions for actions
  - Role hierarchy
  - Resource ownership checks
  - **Deliverable:** `internal/auth/authorization.go`
  - **Tests:** Unit tests for various permission scenarios

- [ ] **2.2.3** Add authorization middleware
  - Check permissions before handler execution
  - Return 403 for unauthorized requests
  - **Deliverable:** Huma middleware for authorization
  - **Tests:** Test various permission scenarios

- [ ] **2.2.4** Update API handlers with permissions
  - Annotate handlers with required permissions
  - Check ownership for resource-specific operations
  - **Deliverable:** Updated API handlers
  - **Tests:** Test as different user roles

- [ ] **2.2.5** Create user management API (admin)
  - `GET /api/v1/users` - List users
  - `GET /api/v1/users/{id}` - Get user
  - `PUT /api/v1/users/{id}/role` - Change user role
  - `DELETE /api/v1/users/{id}` - Delete user
  - **Deliverable:** `internal/api/v1/users.go`
  - **Tests:** API tests as admin and non-admin

### 2.3 Team/Organization Management
**Priority:** HIGH | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **2.3.1** Design teams schema
  - Teams table (id, name, description)
  - Team members table (team_id, user_id, role)
  - Team resource ownership
  - **Deliverable:** Database migrations for teams
  - **Tests:** Apply migrations

- [ ] **2.3.2** Implement team service
  - Create/update/delete teams
  - Add/remove team members
  - Team-based access control
  - **Deliverable:** `internal/team/service.go`
  - **Tests:** Unit tests with mocks, 95%+ coverage

- [ ] **2.3.3** Create team API endpoints
  - `POST /api/v1/teams` - Create team
  - `GET /api/v1/teams` - List teams
  - `GET /api/v1/teams/{id}` - Get team
  - `POST /api/v1/teams/{id}/members` - Add member
  - `DELETE /api/v1/teams/{id}/members/{userID}` - Remove member
  - **Deliverable:** `internal/api/v1/teams.go`
  - **Tests:** API integration tests

- [ ] **2.3.4** Update resource ownership model
  - Resources can be owned by users or teams
  - Team members can access team resources
  - **Deliverable:** Updated project/service repositories
  - **Tests:** Test team-based access

### 2.4 Frontend Authentication UI
**Priority:** HIGH | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **2.4.1** Create login/registration views
  - Login form with validation
  - Registration form
  - Password reset flow
  - **Deliverable:** Auth components in `web/src/views/auth/`
  - **Tests:** Component tests

- [ ] **2.4.2** Add auth state management
  - Auth store (current user, token)
  - Protected routes (Vue Router guards)
  - **Deliverable:** `web/src/stores/auth.js` and router config
  - **Tests:** Store tests and routing tests

- [ ] **2.4.3** Add user profile UI
  - View/edit profile
  - Change password
  - API key management (for later)
  - **Deliverable:** Profile components
  - **Tests:** Component tests

- [ ] **2.4.4** Add admin UI for user management
  - User list with role display
  - Change user roles
  - Invite new users
  - **Deliverable:** Admin components
  - **Tests:** Component tests

---

## Phase 3: Terraform Integration (Weeks 11-15)

**Goal:** Integrate Terraform for actual resource provisioning, state management, and template creation.

### 3.1 Terraform Execution Engine
**Priority:** CRITICAL | **Complexity:** Very High | **Duration:** 5 days

#### Tasks:
- [ ] **3.1.1** Design Terraform integration architecture
  - Terraform workspace per resource/service
  - State file storage (S3, PostgreSQL, or local)
  - Execution isolation (containers or processes)
  - **Deliverable:** Architecture document in `docs/terraform-integration.md`
  - **Tests:** N/A (documentation)

- [ ] **3.1.2** Create Terraform executor interface
  - `Init(ctx, workingDir)`
  - `Plan(ctx, workingDir, vars)`
  - `Apply(ctx, workingDir, vars)`
  - `Destroy(ctx, workingDir, vars)`
  - `Show(ctx, workingDir)` - Get current state
  - **Deliverable:** `internal/terraform/executor.go` interface
  - **Tests:** N/A (interface only)

- [ ] **3.1.3** Implement Terraform CLI executor
  - Execute terraform commands via `os/exec`
  - Parse JSON output
  - Handle errors and timeouts
  - Stream logs
  - **Deliverable:** `internal/terraform/cli/executor.go`
  - **Tests:** Integration tests with real Terraform (requires Terraform binary)

- [ ] **3.1.4** Implement state management
  - Store Terraform state in database
  - Lock mechanism for concurrent operations
  - State versioning and rollback
  - **Deliverable:** `internal/terraform/state/manager.go`
  - **Tests:** Unit tests with mocks, integration tests

- [ ] **3.1.5** Create workspace management
  - Create isolated workspaces for each resource
  - Clean up workspaces on resource deletion
  - Workspace garbage collection
  - **Deliverable:** `internal/terraform/workspace/manager.go`
  - **Tests:** Unit and integration tests

- [ ] **3.1.6** Add Terraform execution logging
  - Capture terraform output
  - Store execution history
  - Stream logs to users
  - **Deliverable:** Logging in executor
  - **Tests:** Verify logs are captured

### 3.2 Terraform Module Registry
**Priority:** HIGH | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **3.2.1** Design module registry schema
  - Modules table (id, name, source, version)
  - Module inputs/outputs schema
  - Module metadata (provider, category)
  - **Deliverable:** Database migrations
  - **Tests:** Apply migrations

- [ ] **3.2.2** Implement module discovery
  - Parse Terraform module sources (local, Git, registry)
  - Extract variables and outputs from module
  - Validate module structure
  - **Deliverable:** `internal/terraform/module/discovery.go`
  - **Tests:** Test with various module sources

- [ ] **3.2.3** Create module-to-template converter
  - Convert Terraform module to resource template
  - Generate JSON schema from variables
  - Map outputs to resource capabilities
  - **Deliverable:** `internal/terraform/module/converter.go`
  - **Tests:** Test with example modules

- [ ] **3.2.4** Implement module API endpoints
  - `POST /api/v1/terraform/modules` - Add module
  - `GET /api/v1/terraform/modules` - List modules
  - `POST /api/v1/terraform/modules/{id}/generate-template` - Create template
  - **Deliverable:** `internal/api/v1/terraform.go`
  - **Tests:** API integration tests

### 3.3 Resource Provisioning via Terraform
**Priority:** CRITICAL | **Complexity:** Very High | **Duration:** 5 days

#### Tasks:
- [ ] **3.3.1** Design provisioning workflow
  - Service creation triggers Terraform apply
  - Status tracking (pending, provisioning, ready, failed)
  - Handle long-running operations (async)
  - **Deliverable:** State machine diagram in docs
  - **Tests:** N/A (documentation)

- [ ] **3.3.2** Implement provisioning service
  - Queue provisioning jobs
  - Execute Terraform apply for services
  - Update service status
  - Handle failures and retries
  - **Deliverable:** `internal/provisioning/service.go`
  - **Tests:** Unit tests with mocked Terraform executor

- [ ] **3.3.3** Add job queue system
  - Background job processing (consider using asynq or similar)
  - Job status tracking
  - Job priority and scheduling
  - **Deliverable:** `internal/jobs/queue.go`
  - **Tests:** Integration tests with job queue

- [ ] **3.3.4** Create provisioning worker
  - Worker process to execute provisioning jobs
  - Graceful shutdown handling
  - Concurrency control
  - **Deliverable:** `cmd/worker/main.go` or goroutine in main
  - **Tests:** Test worker lifecycle

- [ ] **3.3.5** Add provisioning API endpoints
  - `GET /api/v1/services/{id}/status` - Get provisioning status
  - `GET /api/v1/services/{id}/logs` - Get provisioning logs
  - `POST /api/v1/services/{id}/retry` - Retry failed provisioning
  - **Deliverable:** Update `internal/api/v1/services.go`
  - **Tests:** API tests

- [ ] **3.3.6** Implement resource outputs
  - Extract outputs from Terraform state
  - Store outputs with service
  - Expose outputs via API (connection strings, URLs, etc.)
  - **Deliverable:** Output handling in provisioning service
  - **Tests:** Test with modules that have outputs

### 3.4 Resource Lifecycle Management
**Priority:** HIGH | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **3.4.1** Implement resource updates
  - Detect configuration changes
  - Run Terraform plan and show diff
  - Apply updates
  - **Deliverable:** Update logic in provisioning service
  - **Tests:** Test updating service configurations

- [ ] **3.4.2** Implement resource deletion
  - Run Terraform destroy
  - Clean up state and workspaces
  - Handle deletion failures
  - **Deliverable:** Deletion logic in provisioning service
  - **Tests:** Test resource deletion flow

- [ ] **3.4.3** Add resource drift detection
  - Periodically run Terraform plan
  - Detect when actual state differs from desired
  - Alert on drift
  - **Deliverable:** `internal/provisioning/drift.go`
  - **Tests:** Test with manually modified infrastructure

- [ ] **3.4.4** Create frontend provisioning UI
  - Show provisioning status in service cards
  - Display provisioning logs
  - Show Terraform plan diffs before updates
  - **Deliverable:** Updated Vue components
  - **Tests:** Component tests

---

## Phase 4: Batteries Included - Observability (Weeks 16-19)

**Goal:** Integrate Prometheus, Loki, Grafana, and Consul service mesh for complete observability.

### 4.1 Prometheus Integration
**Priority:** HIGH | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **4.1.1** Design metrics collection architecture
  - Service discovery for Prometheus targets
  - Automatic target registration when services are created
  - Support for custom metrics
  - **Deliverable:** Architecture document
  - **Tests:** N/A (documentation)

- [ ] **4.1.2** Implement Prometheus configuration management
  - Generate Prometheus config based on services
  - Service discovery via file or API
  - Auto-reload Prometheus on config changes
  - **Deliverable:** `internal/observability/prometheus/config.go`
  - **Tests:** Test config generation

- [ ] **4.1.3** Add metrics endpoints to provisioned resources
  - Inject metrics sidecar or configure application metrics
  - Standard metrics (CPU, memory, request rate, etc.)
  - **Deliverable:** Update provisioning to include metrics config
  - **Tests:** Provision resource, verify metrics endpoint

- [ ] **4.1.4** Create metrics API endpoints
  - `GET /api/v1/services/{id}/metrics` - Proxy to Prometheus
  - `GET /api/v1/metrics/query` - PromQL query endpoint
  - **Deliverable:** `internal/api/v1/metrics.go`
  - **Tests:** API tests with mocked Prometheus

- [ ] **4.1.5** Create metrics dashboard in frontend
  - Embed Grafana or create custom charts
  - Show service metrics (CPU, memory, requests)
  - Time range selection
  - **Deliverable:** Metrics components in Vue
  - **Tests:** Component tests

### 4.2 Loki Integration (Logging)
**Priority:** HIGH | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **4.2.1** Implement Loki configuration management
  - Configure log shipping for provisioned services
  - Label management (service, project, environment)
  - **Deliverable:** `internal/observability/loki/config.go`
  - **Tests:** Test config generation

- [ ] **4.2.2** Add log collection to services
  - Configure Promtail or Loki agent
  - Collect stdout/stderr from services
  - Structured log parsing
  - **Deliverable:** Update provisioning to include log collection
  - **Tests:** Provision service, verify logs in Loki

- [ ] **4.2.3** Create logs API endpoints
  - `GET /api/v1/services/{id}/logs` - Query logs from Loki
  - Support filtering, time ranges
  - **Deliverable:** `internal/api/v1/logs.go`
  - **Tests:** API tests with mocked Loki

- [ ] **4.2.4** Create logs viewer in frontend
  - Real-time log streaming
  - Log filtering and search
  - Download logs
  - **Deliverable:** Logs components in Vue
  - **Tests:** Component tests

### 4.3 Consul Service Mesh Integration
**Priority:** HIGH | **Complexity:** Very High | **Duration:** 5 days

#### Tasks:
- [ ] **4.3.1** Design service mesh integration
  - Automatic service registration in Consul
  - Service-to-service communication via mesh
  - Traffic management (routing, load balancing)
  - **Deliverable:** Architecture document
  - **Tests:** N/A (documentation)

- [ ] **4.3.2** Implement Consul integration
  - Register services in Consul catalog
  - Configure health checks
  - Service deregistration on deletion
  - **Deliverable:** `internal/observability/consul/service.go`
  - **Tests:** Integration tests with Consul

- [ ] **4.3.3** Add sidecar proxy injection
  - Configure Envoy or Consul Connect sidecars
  - Automatic mTLS between services
  - **Deliverable:** Update provisioning to inject sidecars
  - **Tests:** Test service-to-service communication

- [ ] **4.3.4** Implement distributed tracing
  - Configure tracing (Jaeger or Zipkin)
  - Instrument services for tracing
  - Trace collection and storage
  - **Deliverable:** Tracing configuration
  - **Tests:** Generate traces, verify collection

- [ ] **4.3.5** Create service mesh UI
  - Service topology visualization
  - Service health and dependencies
  - Traffic metrics
  - **Deliverable:** Service mesh components in Vue
  - **Tests:** Component tests

### 4.4 Grafana Dashboards
**Priority:** MEDIUM | **Complexity:** Medium | **Duration:** 2 days

#### Tasks:
- [ ] **4.4.1** Create default dashboards
  - Platform overview dashboard
  - Service-specific dashboards
  - Resource utilization dashboard
  - **Deliverable:** Dashboard JSON in `deployments/grafana/dashboards/`
  - **Tests:** Import dashboards, verify they work

- [ ] **4.4.2** Implement dashboard provisioning
  - Auto-create dashboards for new services
  - Template-based dashboard generation
  - **Deliverable:** Dashboard provisioning service
  - **Tests:** Create service, verify dashboard exists

- [ ] **4.4.3** Embed Grafana in frontend
  - Iframe integration or API-based rendering
  - Single sign-on with Grafana
  - **Deliverable:** Grafana embedding in Vue
  - **Tests:** Verify dashboards load

### 4.5 Alerting
**Priority:** MEDIUM | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **4.5.1** Design alerting system
  - Alert rules (Prometheus alerts)
  - Alert routing and notification
  - Alert history
  - **Deliverable:** Alerting architecture document
  - **Tests:** N/A (documentation)

- [ ] **4.5.2** Implement alert management
  - Create/update/delete alert rules
  - Alert templating per service type
  - **Deliverable:** `internal/observability/alerts/manager.go`
  - **Tests:** Test alert CRUD operations

- [ ] **4.5.3** Add notification channels
  - Email, Slack, PagerDuty, webhooks
  - Per-team notification preferences
  - **Deliverable:** Notification service
  - **Tests:** Test notifications to different channels

- [ ] **4.5.4** Create alerts UI
  - Alert configuration interface
  - Alert history and status
  - Silence/acknowledge alerts
  - **Deliverable:** Alerts components in Vue
  - **Tests:** Component tests

---

## Phase 5: GitOps (Weeks 20-22)

**Goal:** Enable declarative resource management via Git repositories with automatic reconciliation.

### 5.1 Git Repository Integration
**Priority:** HIGH | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **5.1.1** Design GitOps workflow
  - Git repository structure for resource definitions
  - Reconciliation loop (detect changes, apply)
  - Conflict resolution strategy
  - **Deliverable:** GitOps workflow document
  - **Tests:** N/A (documentation)

- [ ] **5.1.2** Implement Git repository connector
  - Clone repositories (public and private with auth)
  - Pull updates
  - Webhook support for push events
  - **Deliverable:** `internal/gitops/repository/connector.go`
  - **Tests:** Integration tests with test repositories

- [ ] **5.1.3** Create resource definition format
  - YAML format for projects, services, configurations
  - Schema validation
  - Examples and documentation
  - **Deliverable:** YAML schemas and parser in `internal/gitops/schema/`
  - **Tests:** Parse valid and invalid YAML files

- [ ] **5.1.4** Implement repository API endpoints
  - `POST /api/v1/gitops/repositories` - Connect repository
  - `GET /api/v1/gitops/repositories` - List repositories
  - `POST /api/v1/gitops/repositories/{id}/sync` - Trigger sync
  - **Deliverable:** `internal/api/v1/gitops.go`
  - **Tests:** API integration tests

### 5.2 Reconciliation Engine
**Priority:** HIGH | **Complexity:** Very High | **Duration:** 5 days

#### Tasks:
- [ ] **5.2.1** Design reconciliation loop
  - Detect differences between Git and cluster state
  - Determine required actions (create, update, delete)
  - Order resources by dependencies
  - **Deliverable:** Reconciliation algorithm document
  - **Tests:** N/A (documentation)

- [ ] **5.2.2** Implement reconciliation service
  - Compare desired state (Git) with actual state (database)
  - Generate reconciliation plan
  - Execute plan (create/update/delete resources)
  - **Deliverable:** `internal/gitops/reconciliation/service.go`
  - **Tests:** Unit tests with various scenarios

- [ ] **5.2.3** Add reconciliation scheduling
  - Periodic reconciliation (e.g., every 5 minutes)
  - Webhook-triggered reconciliation
  - Manual sync trigger
  - **Deliverable:** Scheduler in reconciliation service
  - **Tests:** Test scheduling behavior

- [ ] **5.2.4** Implement conflict detection
  - Detect manual changes vs Git changes
  - Alert on conflicts
  - Conflict resolution strategies (Git wins, manual wins, merge)
  - **Deliverable:** Conflict detection logic
  - **Tests:** Create conflicts, verify detection

- [ ] **5.2.5** Add reconciliation status tracking
  - Track sync status per repository
  - Store reconciliation history
  - Show drift from Git
  - **Deliverable:** Status tracking in database
  - **Tests:** Verify status updates during reconciliation

### 5.3 GitOps UI
**Priority:** MEDIUM | **Complexity:** Medium | **Duration:** 2 days

#### Tasks:
- [ ] **5.3.1** Create repository management UI
  - Add/remove Git repositories
  - Configure branch and path
  - Show connection status
  - **Deliverable:** Repository components in Vue
  - **Tests:** Component tests

- [ ] **5.3.2** Create reconciliation status UI
  - Show sync status
  - Display drift and conflicts
  - Trigger manual sync
  - Show reconciliation logs
  - **Deliverable:** Reconciliation components in Vue
  - **Tests:** Component tests

- [ ] **5.3.3** Add resource source indicator
  - Show if resource is managed by Git or manually
  - Prevent manual changes to Git-managed resources
  - **Deliverable:** Update existing components
  - **Tests:** Component tests

---

## Phase 6: Advanced Features (Weeks 23-27)

**Goal:** Add advanced capabilities like special resource types, dependencies, cost tracking, and lifecycle management.

### 6.1 Special Resource Types (Stretch Goal)
**Priority:** MEDIUM | **Complexity:** High | **Duration:** 5 days

#### Tasks:
- [ ] **6.1.1** Design special resource type system
  - Define categories: databases, compute, jobs, search engines, message queues
  - Standard interfaces for each category
  - Type-specific operations
  - **Deliverable:** Design document
  - **Tests:** N/A (documentation)

- [ ] **6.1.2** Implement database resource type
  - Standard inputs: size, version, backup schedule
  - Standard outputs: connection string, credentials
  - Lifecycle hooks: backup, restore, scale
  - **Deliverable:** `internal/resource/types/database/` package
  - **Tests:** Unit tests for database resource logic

- [ ] **6.1.3** Implement compute resource type
  - Standard inputs: CPU, memory, image, command
  - Auto-scaling configuration
  - **Deliverable:** `internal/resource/types/compute/` package
  - **Tests:** Unit tests

- [ ] **6.1.4** Implement job resource type
  - Scheduled and event-triggered jobs
  - Job history and logs
  - **Deliverable:** `internal/resource/types/job/` package
  - **Tests:** Unit tests

- [ ] **6.1.5** Create UI for special resource types
  - Type-specific configuration forms
  - Type-specific dashboards and actions
  - **Deliverable:** Type-specific Vue components
  - **Tests:** Component tests

### 6.2 Resource Dependencies
**Priority:** HIGH | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **6.2.1** Design dependency system
  - Directed acyclic graph (DAG) for dependencies
  - Dependency types: hard (required), soft (optional)
  - Provisioning order based on dependencies
  - **Deliverable:** Dependency design document
  - **Tests:** N/A (documentation)

- [ ] **6.2.2** Implement dependency graph
  - Store dependencies in database
  - Detect cycles
  - Topological sort for provisioning order
  - **Deliverable:** `internal/dependencies/graph.go`
  - **Tests:** Unit tests with various dependency scenarios

- [ ] **6.2.3** Update provisioning for dependencies
  - Provision dependencies first
  - Pass outputs to dependent resources
  - Handle dependency failures
  - **Deliverable:** Update provisioning service
  - **Tests:** Integration tests with dependent services

- [ ] **6.2.4** Add dependency API endpoints
  - `POST /api/v1/services/{id}/dependencies` - Add dependency
  - `DELETE /api/v1/services/{id}/dependencies/{depID}` - Remove
  - `GET /api/v1/services/{id}/dependencies` - List
  - **Deliverable:** Dependency endpoints
  - **Tests:** API tests

- [ ] **6.2.5** Create dependency visualization UI
  - Dependency graph visualization
  - Add/remove dependencies via UI
  - Show dependency impact (what breaks if I delete this?)
  - **Deliverable:** Dependency graph component
  - **Tests:** Component tests

### 6.3 Cost Tracking and Budgets
**Priority:** MEDIUM | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **6.3.1** Implement cost calculation
  - Calculate resource costs based on templates
  - Aggregate project costs
  - Track cost over time
  - **Deliverable:** `internal/billing/calculator.go`
  - **Tests:** Unit tests with various scenarios

- [ ] **6.3.2** Add budget management
  - Set budgets per project or team
  - Alert when approaching budget limits
  - Enforce budget limits (prevent provisioning)
  - **Deliverable:** `internal/billing/budgets.go`
  - **Tests:** Unit tests for budget enforcement

- [ ] **6.3.3** Create cost API endpoints
  - `GET /api/v1/projects/{id}/cost` - Get current and historical costs
  - `POST /api/v1/projects/{id}/budget` - Set budget
  - `GET /api/v1/billing/reports` - Cost reports
  - **Deliverable:** Billing API endpoints
  - **Tests:** API tests

- [ ] **6.3.4** Create cost dashboard UI
  - Cost breakdown by project/service
  - Budget status and alerts
  - Cost trends over time
  - **Deliverable:** Cost dashboard components
  - **Tests:** Component tests

### 6.4 Vault Integration (Secrets Management)
**Priority:** HIGH | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **6.4.1** Design Vault integration
  - Secrets storage in Vault
  - Dynamic secrets for databases
  - Secret injection into services
  - **Deliverable:** Vault integration design document
  - **Tests:** N/A (documentation)

- [ ] **6.4.2** Implement Vault client
  - Authenticate with Vault
  - CRUD operations for secrets
  - Lease renewal for dynamic secrets
  - **Deliverable:** `internal/secrets/vault/client.go`
  - **Tests:** Integration tests with Vault

- [ ] **6.4.3** Add secret management service
  - Store secrets in Vault
  - Inject secrets as environment variables or files
  - Rotate secrets
  - **Deliverable:** `internal/secrets/service.go`
  - **Tests:** Unit tests with mocked Vault

- [ ] **6.4.4** Update provisioning for secrets
  - Configure secret injection in Terraform
  - Pass secrets securely to services
  - **Deliverable:** Update provisioning service
  - **Tests:** Provision service with secrets

- [ ] **6.4.5** Create secrets management UI
  - Add/update/delete secrets
  - Secret history and rotation
  - Secret usage (which services use this secret)
  - **Deliverable:** Secrets components in Vue
  - **Tests:** Component tests

### 6.5 Resource Lifecycle Policies
**Priority:** MEDIUM | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **6.5.1** Design lifecycle policies
  - Auto-delete after inactivity
  - Auto-scale based on metrics
  - Scheduled start/stop (cost saving)
  - **Deliverable:** Lifecycle policy design document
  - **Tests:** N/A (documentation)

- [ ] **6.5.2** Implement policy engine
  - Evaluate policies periodically
  - Execute policy actions
  - Policy audit log
  - **Deliverable:** `internal/lifecycle/policies.go`
  - **Tests:** Unit tests for policy evaluation

- [ ] **6.5.3** Add lifecycle API endpoints
  - `POST /api/v1/services/{id}/policies` - Add policy
  - `GET /api/v1/services/{id}/policies` - List policies
  - `DELETE /api/v1/services/{id}/policies/{policyID}` - Remove
  - **Deliverable:** Lifecycle API endpoints
  - **Tests:** API tests

- [ ] **6.5.4** Create policy configuration UI
  - Configure auto-delete, auto-scale, scheduling
  - Show policy status and history
  - **Deliverable:** Policy components in Vue
  - **Tests:** Component tests

### 6.6 Multi-Cloud and Hybrid Cloud Support
**Priority:** MEDIUM | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **6.6.1** Design multi-cloud architecture
  - Support multiple cloud providers (AWS, GCP, Azure, on-prem)
  - Provider abstraction layer
  - Cross-cloud networking
  - **Deliverable:** Multi-cloud design document
  - **Tests:** N/A (documentation)

- [ ] **6.6.2** Implement provider registry
  - Register cloud providers with credentials
  - Select provider per resource template
  - **Deliverable:** `internal/providers/registry.go`
  - **Tests:** Unit tests

- [ ] **6.6.3** Add provider-specific Terraform backends
  - Configure Terraform to use correct provider
  - Manage provider credentials securely (via Vault)
  - **Deliverable:** Update Terraform executor
  - **Tests:** Integration tests with multiple providers

- [ ] **6.6.4** Create provider management UI
  - Add/remove cloud providers
  - Configure provider credentials
  - Show resources per provider
  - **Deliverable:** Provider management components
  - **Tests:** Component tests

---

## Phase 7: Production Readiness (Weeks 28-30)

**Goal:** Harden the platform for production use with backups, disaster recovery, performance optimization, and comprehensive documentation.

### 7.1 Backup and Disaster Recovery
**Priority:** CRITICAL | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **7.1.1** Implement database backups
  - Automated periodic backups
  - Backup to object storage (S3, GCS, etc.)
  - Backup retention policies
  - **Deliverable:** Backup service
  - **Tests:** Create backup, verify it can be restored

- [ ] **7.1.2** Implement Terraform state backups
  - Backup state files periodically
  - State versioning
  - **Deliverable:** State backup in state manager
  - **Tests:** Backup and restore state

- [ ] **7.1.3** Create disaster recovery procedures
  - Document backup/restore process
  - Test recovery procedures
  - RTO/RPO targets
  - **Deliverable:** DR documentation
  - **Tests:** Full DR drill

- [ ] **7.1.4** Add backup management UI
  - View backup history
  - Trigger manual backups
  - Restore from backup
  - **Deliverable:** Backup components
  - **Tests:** Component tests

### 7.2 Performance Optimization
**Priority:** HIGH | **Complexity:** Medium | **Duration:** 3 days

#### Tasks:
- [ ] **7.2.1** Add database query optimization
  - Add indexes for common queries
  - Optimize N+1 queries
  - Query performance profiling
  - **Deliverable:** Database migrations with indexes
  - **Tests:** Benchmark queries before/after

- [ ] **7.2.2** Add API response caching
  - Cache frequently accessed data (templates, etc.)
  - Cache invalidation strategy
  - **Deliverable:** Caching middleware
  - **Tests:** Test cache hits/misses

- [ ] **7.2.3** Optimize frontend performance
  - Code splitting
  - Lazy loading components
  - Bundle size optimization
  - **Deliverable:** Optimized Vue build
  - **Tests:** Lighthouse performance score

- [ ] **7.2.4** Add rate limiting
  - Rate limit API requests per user
  - Prevent abuse
  - **Deliverable:** Rate limiting middleware
  - **Tests:** Test rate limiting behavior

### 7.3 Security Hardening
**Priority:** CRITICAL | **Complexity:** High | **Duration:** 4 days

#### Tasks:
- [ ] **7.3.1** Security audit
  - Review code for common vulnerabilities (OWASP Top 10)
  - Dependency vulnerability scanning
  - **Deliverable:** Security audit report
  - **Tests:** Fix identified vulnerabilities

- [ ] **7.3.2** Add input validation and sanitization
  - Validate all inputs against schemas
  - Sanitize outputs to prevent XSS
  - Parameterized queries (already done with sqlx)
  - **Deliverable:** Enhanced validation
  - **Tests:** Test with malicious inputs

- [ ] **7.3.3** Implement audit logging
  - Log all sensitive operations (user actions, resource changes)
  - Tamper-proof logs
  - Log retention
  - **Deliverable:** Audit logging service
  - **Tests:** Verify audit logs are created

- [ ] **7.3.4** Add security headers
  - CSP, HSTS, X-Frame-Options, etc.
  - CORS configuration
  - **Deliverable:** Security headers middleware
  - **Tests:** Verify headers in responses

- [ ] **7.3.5** Implement API key management
  - Generate API keys for programmatic access
  - Key rotation
  - Key permissions (scoped keys)
  - **Deliverable:** API key service
  - **Tests:** Test authentication with API keys

### 7.4 Documentation
**Priority:** HIGH | **Complexity:** Low | **Duration:** 3 days

#### Tasks:
- [ ] **7.4.1** Create user documentation
  - Getting started guide
  - How-to guides (create project, add resources, etc.)
  - FAQ
  - **Deliverable:** User docs in `docs/user/`
  - **Tests:** Manual review by non-developers

- [ ] **7.4.2** Create admin documentation
  - Installation guide
  - Configuration reference
  - Terraform module guide
  - Observability stack setup
  - **Deliverable:** Admin docs in `docs/admin/`
  - **Tests:** Follow installation guide on fresh system

- [ ] **7.4.3** Create developer documentation
  - Architecture overview
  - API reference (OpenAPI)
  - Contributing guide
  - Development environment setup
  - **Deliverable:** Developer docs in `docs/developer/`
  - **Tests:** New developer follows setup guide

- [ ] **7.4.4** Create video tutorials
  - Platform overview video
  - Creating first project walkthrough
  - Admin setup walkthrough
  - **Deliverable:** Videos on YouTube or hosted
  - **Tests:** User feedback on clarity

### 7.5 Operational Tooling
**Priority:** MEDIUM | **Complexity:** Medium | **Duration:** 2 days

#### Tasks:
- [ ] **7.5.1** Create admin CLI tool
  - User management commands
  - Database migrations
  - System health checks
  - **Deliverable:** `cmd/admin/` CLI tool
  - **Tests:** Test each CLI command

- [ ] **7.5.2** Add system diagnostics
  - Health check for all dependencies (DB, Vault, Consul, etc.)
  - Version information
  - System resource usage
  - **Deliverable:** Diagnostics endpoints and CLI commands
  - **Tests:** Run diagnostics, verify output

- [ ] **7.5.3** Create upgrade guide
  - Version compatibility matrix
  - Migration guides between versions
  - Rollback procedures
  - **Deliverable:** Upgrade documentation
  - **Tests:** Test upgrade from previous version

---

## Success Metrics

### Phase 0 (Foundation)
- [ ] 95%+ test coverage for all Go code
- [ ] 80%+ test coverage for all Vue code
- [ ] All CI/CD pipelines green
- [ ] Single binary deploys successfully on Linux, macOS, Windows
- [ ] Docker container runs application successfully

### Phase 1 (Core Platform)
- [ ] Can create, read, update, delete projects via API
- [ ] Can create, read, update, delete services via API
- [ ] Resource templates can be created and browsed
- [ ] Frontend successfully communicates with backend
- [ ] Users can visually create projects with services

### Phase 2 (Auth & Access Control)
- [ ] Users can register and log in
- [ ] Role-based permissions work correctly
- [ ] Teams can be created and managed
- [ ] Resources are properly scoped to users/teams

### Phase 3 (Terraform)
- [ ] Terraform modules can be imported as templates
- [ ] Services are provisioned via Terraform apply
- [ ] Service state is tracked and displayed
- [ ] Services can be updated and destroyed via Terraform

### Phase 4 (Observability)
- [ ] Prometheus collects metrics from services
- [ ] Loki collects logs from services
- [ ] Grafana dashboards show service metrics
- [ ] Service mesh provides mTLS and tracing
- [ ] Alerts fire and notify correctly

### Phase 5 (GitOps)
- [ ] Git repositories can be connected
- [ ] Resources defined in Git are created automatically
- [ ] Changes in Git trigger reconciliation
- [ ] Drift is detected and displayed

### Phase 6 (Advanced Features)
- [ ] Special resource types (databases, compute, jobs) work
- [ ] Resource dependencies are enforced
- [ ] Costs are calculated and tracked
- [ ] Secrets are managed via Vault
- [ ] Lifecycle policies execute correctly

### Phase 7 (Production)
- [ ] Backups can be created and restored
- [ ] Performance meets SLAs (p95 API response time < 200ms)
- [ ] Security audit passes with no critical issues
- [ ] Documentation is complete and accurate

---

## Technology Stack Summary

### Backend
- **Language:** Go 1.23+
- **API Framework:** Huma v2
- **CLI Framework:** Cobra
- **Database:** PostgreSQL (production), SQLite (dev/testing)
- **Migrations:** golang-migrate or similar
- **Background Jobs:** asynq or custom queue
- **Testing:** Go testing package, testify

### Frontend
- **Framework:** Vue 3 (Composition API)
- **Build Tool:** Vite
- **UI Library:** PrimeVue
- **State Management:** Pinia
- **Routing:** Vue Router
- **Testing:** Vitest, @vue/test-utils
- **Icons:** Lucide

### Infrastructure
- **Provisioning:** Terraform
- **Service Mesh:** Consul Connect
- **Secrets:** Vault
- **Metrics:** Prometheus
- **Logs:** Loki + Promtail
- **Dashboards:** Grafana
- **Tracing:** Jaeger or Zipkin

### DevOps
- **CI/CD:** GitHub Actions
- **Containers:** Docker
- **Container Registry:** GitHub Container Registry or Docker Hub

---

## Risk Mitigation

### Technical Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Terraform state conflicts | High | High | Implement robust locking, state versioning |
| Long-running Terraform operations timeout | Medium | Medium | Async job queue, proper timeout handling |
| Service mesh complexity | Medium | High | Start with simple use cases, comprehensive docs |
| Multi-tenancy security issues | Medium | Critical | Security audit, extensive testing, input validation |
| Performance issues with many resources | Medium | Medium | Database indexing, caching, pagination |
| Vault unavailability breaks provisioning | Medium | High | Graceful degradation, local fallback for dev |

### Organizational Risks

| Risk | Probability | Impact | Mitigation |
|------|-------------|--------|------------|
| Scope creep | High | Medium | Strict phase definitions, defer stretch goals |
| Lack of testing discipline | Medium | Critical | CI enforcement, coverage requirements |
| Documentation becomes stale | Medium | Medium | Docs as part of definition of done |
| Unclear requirements | Medium | High | Frequent user feedback, iterative development |

---

## Dependencies and Prerequisites

### Required Infrastructure (for full feature set)
- **Kubernetes cluster** (for deploying user services)
- **PostgreSQL database** (for application data)
- **Prometheus instance** (for metrics)
- **Loki instance** (for logs)
- **Grafana instance** (for dashboards)
- **Consul cluster** (for service mesh)
- **Vault instance** (for secrets)

### Development Dependencies
- Go 1.23+
- Node.js 20+
- Terraform CLI
- Docker
- PostgreSQL (for integration tests)

### Optional Dependencies
- Kubernetes (minikube or kind for local dev)
- S3-compatible storage (for backups)

---

## Iteration and Feedback Loops

1. **Weekly demos** at end of each week showing progress
2. **End-of-phase reviews** to validate success metrics
3. **User feedback sessions** after phases 1, 3, and 5
4. **Security review** before phase 7
5. **Performance testing** at end of phases 3, 4, and 7

---

## Next Steps

1. **Review and approval** of this roadmap
2. **Set up development environment** (Phase 0.5.2)
3. **Begin Phase 0.1** (Testing Infrastructure)
4. **Establish weekly sync meetings** for progress tracking

---

## Glossary

- **Resource:** An infrastructure component (e.g., database, VM, k8s deployment)
- **Resource Template:** A reusable definition for creating resources
- **Service:** An instance of a resource template within a project
- **Project:** A collection of related services
- **Terraform Module:** A reusable Terraform configuration
- **GitOps:** Managing infrastructure declaratively via Git
- **Service Mesh:** Infrastructure layer for service-to-service communication
- **Reconciliation:** Process of making actual state match desired state

---

**Document Version:** 1.0
**Last Updated:** 2025-11-17
**Owner:** Mahler Platform Team
