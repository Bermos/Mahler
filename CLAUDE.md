# Claude Development Guidelines for Mahler

This document provides guidelines for AI-assisted development on the Mahler project.

## Project Overview

Mahler is a platform built with Go backend and Vue.js frontend for managing cloud resources and applications. The project follows a modular architecture with clear separation between API, business logic, and resource management.

## Technology Stack

- **Backend**: Go 1.23+
- **Frontend**: Vue.js 3 with Vite
- **API Framework**: Huma v2
- **CLI Framework**: Cobra

## General Principles

### Code Quality
- Write clean, maintainable, and well-documented code
- Follow established patterns and conventions in the codebase
- Prioritize readability and simplicity over cleverness
- Use meaningful names for variables, functions, and types
- Keep functions focused and single-purpose

### Version Control
- Write clear, descriptive commit messages
- Make atomic commits (one logical change per commit)
- Follow conventional commit format when applicable

## Go Development Guidelines

### Library Selection
- **Prefer the standard library** for common functionality (http, json, context, etc.)
- Use well-established third-party libraries only when necessary
- Avoid reinventing the wheel - leverage proven solutions
- Current approved libraries:
  - `github.com/danielgtaylor/huma/v2` - API framework
  - `github.com/spf13/cobra` - CLI framework
  - `github.com/google/uuid` - UUID generation

### SOLID Principles
Strictly adhere to SOLID principles:

1. **Single Responsibility Principle**
   - Each type/function should have one reason to change
   - Keep concerns separated (business logic, persistence, presentation)

2. **Open/Closed Principle**
   - Design for extension, not modification
   - Use interfaces and composition

3. **Liskov Substitution Principle**
   - Interfaces should be substitutable for their implementations
   - Maintain behavioral contracts

4. **Interface Segregation Principle**
   - Keep interfaces small and focused
   - Clients shouldn't depend on methods they don't use

5. **Dependency Inversion Principle**
   - Depend on abstractions, not concretions
   - Use dependency injection
   - Define interfaces in consumer packages, not implementation packages

### Code Organization

```
cmd/          - Application entry points
internal/     - Private application code
  api/        - API routes and handlers
  app/        - Application domain logic
  project/    - Project management
  service/    - Service orchestration
  resource/   - Resource implementations
    interface.go  - Resource interface definitions
    */        - Specific resource implementations
```

### Testing Requirements

#### Coverage Targets
- **Minimum**: 95% test coverage for all Go code
- Run coverage reports: `go test -cover ./...`
- Use `go test -coverprofile=coverage.out ./...` for detailed analysis

#### Test Organization
- Place tests in the same package as the code being tested (`*_test.go`)
- Use table-driven tests for multiple scenarios
- Separate unit tests, integration tests, and end-to-end tests

#### Test Design Patterns
Structure tests to support both integration and unit testing:

```go
// Define interfaces for external dependencies
type UserRepository interface {
    Get(id string) (*User, error)
    Save(user *User) error
}

// Production implementation
type postgresUserRepository struct { ... }

// Mock for testing
type mockUserRepository struct { ... }

// Service accepts interface, not concrete type
func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}
```

**Unit Tests**: Use mocked providers/repositories
```go
func TestUserService_CreateUser(t *testing.T) {
    mockRepo := &mockUserRepository{...}
    svc := NewUserService(mockRepo)
    // Test business logic in isolation
}
```

**Integration Tests**: Use real implementations with test infrastructure
```go
func TestUserService_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping integration test")
    }
    repo := setupTestDatabase(t)
    svc := NewUserService(repo)
    // Test with real dependencies
}
```

#### Test Best Practices
- Use `t.Helper()` for test helper functions
- Use `t.Parallel()` for independent tests
- Clean up resources with `t.Cleanup()`
- Use meaningful test names: `TestFunction_Scenario_ExpectedBehavior`
- Assert one thing per test when possible
- Use subtests for related test cases

### Error Handling
- Return errors, don't panic (except in truly exceptional cases)
- Wrap errors with context: `fmt.Errorf("failed to save user: %w", err)`
- Create custom error types for domain-specific errors
- Handle errors at appropriate levels

### Concurrency
- Use channels and goroutines appropriately
- Always handle context cancellation
- Avoid shared mutable state
- Use sync primitives when necessary (Mutex, RWMutex, WaitGroup)
- Document concurrency safety of types

### Code Style
- Follow standard Go formatting (`gofmt`, `goimports`)
- Use `golangci-lint` for linting
- Keep exported APIs minimal and well-documented
- Write godoc comments for all exported symbols
- Use Go's built-in tooling: `go vet`, `go test`, `go mod tidy`

## Vue.js Development Guidelines

### Project Structure
```
web/
  src/
    components/  - Reusable Vue components
    App.vue      - Root component
  public/        - Static assets
```

### Testing Requirements
- **Minimum**: 80% test coverage for Vue components
- Use Vitest for unit testing
- Test component behavior, not implementation details
- Test user interactions and edge cases

### Component Best Practices
- Keep components focused and composable
- Use Composition API (script setup)
- Properly type props and emits
- Follow Vue.js style guide
- Use meaningful component names

### State Management
- Start with props/events for component communication
- Use Pinia for complex state management if needed
- Keep state as local as possible

## CI/CD Guidelines

### Continuous Integration

#### Required Jobs
1. **Test Job**
   - Run all Go tests: `go test ./...`
   - Run Vue tests: `npm test`
   - Generate coverage reports
   - Fail if coverage drops below thresholds

2. **Coverage Reporting**
   - Generate Go coverage: `go test -coverprofile=coverage.out ./...`
   - Generate Vue coverage (via Vitest)
   - Report coverage metrics
   - Optional: Upload to coverage service (Codecov, Coveralls)

3. **Lint Job**
   - Run `golangci-lint` for Go
   - Run ESLint for Vue/JavaScript
   - Enforce code style standards

4. **Build Job**
   - Build Go binary: `go build -o mahler ./cmd`
   - Build Vue frontend: `npm run build`
   - Verify successful compilation

### Continuous Deployment

#### Binary Distribution
- Build a **single, working binary** containing:
  - Embedded Vue.js frontend (using `embed` package)
  - All necessary assets
  - Cross-platform builds (Linux, macOS, Windows)
- Make binary easy to deploy and run standalone
- Example: `./mahler serve`

#### Docker Container
- Publish Docker container for flexible deployment
- Use multi-stage builds for minimal image size
- Include both backend and frontend in single container
- Tag with version and latest
- Document required environment variables
- Expose necessary ports
- Example run: `docker run -p 8080:8080 mahler:latest`

#### Container Best Practices
```dockerfile
# Multi-stage build
FROM golang:1.23 AS backend-builder
# Build Go binary

FROM node:20 AS frontend-builder
# Build Vue frontend

FROM alpine:latest
# Copy binary and assets
# Single container, complete application
```

## Architecture Patterns

### Layered Architecture
```
Presentation Layer (API/CLI)
    ↓
Business Logic Layer (Services)
    ↓
Data Access Layer (Repositories)
    ↓
Resources/External Systems
```

### Dependency Injection
- Pass dependencies through constructors
- Use interfaces for testability
- Avoid global state and singletons
- Make dependencies explicit

### Resource Pattern
- Implement `Resource` interface for all managed resources
- Keep resource implementations isolated in their own packages
- Use dependency injection for resource dependencies

## Documentation

### Code Documentation
- Document all exported types, functions, and constants
- Include usage examples for complex functionality
- Document assumptions and constraints
- Keep documentation up-to-date with code changes

### API Documentation
- Use Huma's automatic OpenAPI generation
- Document all endpoints, parameters, and responses
- Include example requests/responses
- Document error cases

## Security Considerations

- Validate all input
- Use parameterized queries to prevent injection
- Handle sensitive data appropriately
- Follow principle of least privilege
- Keep dependencies updated
- Use context timeouts for external calls

## Performance Considerations

- Profile before optimizing
- Use appropriate data structures
- Avoid premature optimization
- Cache when beneficial
- Use database indexes appropriately
- Minimize allocations in hot paths

## Common Patterns

### Context Usage
```go
func (s *Service) DoWork(ctx context.Context, input Input) error {
    // Always accept context as first parameter
    // Propagate context to downstream calls
    // Check for cancellation in long-running operations
}
```

### Error Handling
```go
if err != nil {
    return fmt.Errorf("descriptive context: %w", err)
}
```

### Interface Definition
```go
// Define interfaces where they're used (consumer side)
type UserGetter interface {
    GetUser(ctx context.Context, id string) (*User, error)
}
```

## Pre-commit Checklist

Before committing code, ensure:
- [ ] All tests pass: `go test ./...` and `npm test`
- [ ] Coverage meets requirements (>95% Go, >80% Vue)
- [ ] Code is formatted: `gofmt`, `goimports`
- [ ] Linters pass: `golangci-lint run`
- [ ] No unnecessary dependencies added
- [ ] Documentation is updated
- [ ] SOLID principles are followed
- [ ] Interfaces are properly defined
- [ ] Tests use appropriate mocking/integration patterns

## Asking for Clarification

When requirements are ambiguous:
1. Ask specific questions about implementation approach
2. Propose options with trade-offs
3. Clarify architectural decisions
4. Confirm testing strategy
5. Validate assumptions before proceeding

## Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Vue.js Best Practices](https://vuejs.org/style-guide/)
- [SOLID Principles in Go](https://dave.cheney.net/2016/08/20/solid-go-design)
- [Huma Documentation](https://huma.rocks/)
