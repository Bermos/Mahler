package testutil

import (
	"context"
	"fmt"
	"sync"

	"github.com/Bermos/Platform/internal/project"
	"github.com/Bermos/Platform/internal/service"
	"github.com/google/uuid"
)

// ProjectRepository defines the interface for project persistence operations
// This interface will be moved to internal/project package in Phase 0.3.3
type ProjectRepository interface {
	Create(ctx context.Context, proj *project.Project) error
	Get(ctx context.Context, id uuid.UUID) (*project.Project, error)
	Update(ctx context.Context, proj *project.Project) error
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context) ([]*project.Project, error)
}

// ServiceRepository defines the interface for service persistence operations
// This interface will be moved to internal/service package in Phase 0.3.3
type ServiceRepository interface {
	Create(ctx context.Context, svc *service.Service) error
	Get(ctx context.Context, id uuid.UUID) (*service.Service, error)
	Update(ctx context.Context, svc *service.Service) error
	Delete(ctx context.Context, id uuid.UUID) error
	ListByProject(ctx context.Context, projectID uuid.UUID) ([]*service.Service, error)
}

// MockProjectRepository is an in-memory mock implementation of ProjectRepository
type MockProjectRepository struct {
	mu       sync.RWMutex
	projects map[uuid.UUID]*project.Project
	// For simulating errors
	CreateError error
	GetError    error
	UpdateError error
	DeleteError error
	ListError   error
}

// NewMockProjectRepository creates a new mock project repository
func NewMockProjectRepository() *MockProjectRepository {
	return &MockProjectRepository{
		projects: make(map[uuid.UUID]*project.Project),
	}
}

// Create adds a project to the mock repository
func (r *MockProjectRepository) Create(ctx context.Context, proj *project.Project) error {
	if r.CreateError != nil {
		return r.CreateError
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.projects[proj.ID]; exists {
		return fmt.Errorf("project with ID %s already exists", proj.ID)
	}

	r.projects[proj.ID] = proj
	return nil
}

// Get retrieves a project from the mock repository
func (r *MockProjectRepository) Get(ctx context.Context, id uuid.UUID) (*project.Project, error) {
	if r.GetError != nil {
		return nil, r.GetError
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	proj, exists := r.projects[id]
	if !exists {
		return nil, fmt.Errorf("project with ID %s not found", id)
	}

	return proj, nil
}

// Update modifies a project in the mock repository
func (r *MockProjectRepository) Update(ctx context.Context, proj *project.Project) error {
	if r.UpdateError != nil {
		return r.UpdateError
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.projects[proj.ID]; !exists {
		return fmt.Errorf("project with ID %s not found", proj.ID)
	}

	r.projects[proj.ID] = proj
	return nil
}

// Delete removes a project from the mock repository
func (r *MockProjectRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if r.DeleteError != nil {
		return r.DeleteError
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.projects[id]; !exists {
		return fmt.Errorf("project with ID %s not found", id)
	}

	delete(r.projects, id)
	return nil
}

// List returns all projects from the mock repository
func (r *MockProjectRepository) List(ctx context.Context) ([]*project.Project, error) {
	if r.ListError != nil {
		return nil, r.ListError
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	projects := make([]*project.Project, 0, len(r.projects))
	for _, proj := range r.projects {
		projects = append(projects, proj)
	}

	return projects, nil
}

// MockServiceRepository is an in-memory mock implementation of ServiceRepository
type MockServiceRepository struct {
	mu       sync.RWMutex
	services map[uuid.UUID]*service.Service
	// For simulating errors
	CreateError        error
	GetError           error
	UpdateError        error
	DeleteError        error
	ListByProjectError error
}

// NewMockServiceRepository creates a new mock service repository
func NewMockServiceRepository() *MockServiceRepository {
	return &MockServiceRepository{
		services: make(map[uuid.UUID]*service.Service),
	}
}

// Create adds a service to the mock repository
func (r *MockServiceRepository) Create(ctx context.Context, svc *service.Service) error {
	if r.CreateError != nil {
		return r.CreateError
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.services[svc.ID]; exists {
		return fmt.Errorf("service with ID %s already exists", svc.ID)
	}

	r.services[svc.ID] = svc
	return nil
}

// Get retrieves a service from the mock repository
func (r *MockServiceRepository) Get(ctx context.Context, id uuid.UUID) (*service.Service, error) {
	if r.GetError != nil {
		return nil, r.GetError
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	svc, exists := r.services[id]
	if !exists {
		return nil, fmt.Errorf("service with ID %s not found", id)
	}

	return svc, nil
}

// Update modifies a service in the mock repository
func (r *MockServiceRepository) Update(ctx context.Context, svc *service.Service) error {
	if r.UpdateError != nil {
		return r.UpdateError
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.services[svc.ID]; !exists {
		return fmt.Errorf("service with ID %s not found", svc.ID)
	}

	r.services[svc.ID] = svc
	return nil
}

// Delete removes a service from the mock repository
func (r *MockServiceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if r.DeleteError != nil {
		return r.DeleteError
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.services[id]; !exists {
		return fmt.Errorf("service with ID %s not found", id)
	}

	delete(r.services, id)
	return nil
}

// ListByProject returns all services for a given project from the mock repository
// Note: In a real implementation, this would filter by project ID
// For now, this mock implementation returns all services
func (r *MockServiceRepository) ListByProject(ctx context.Context, projectID uuid.UUID) ([]*service.Service, error) {
	if r.ListByProjectError != nil {
		return nil, r.ListByProjectError
	}
	r.mu.RLock()
	defer r.mu.RUnlock()

	services := make([]*service.Service, 0, len(r.services))
	for _, svc := range r.services {
		services = append(services, svc)
	}

	return services, nil
}
