package testutil

import (
	"github.com/Bermos/Platform/internal/project"
	"github.com/Bermos/Platform/internal/service"
	"github.com/google/uuid"
)

// ProjectBuilder helps build test projects
type ProjectBuilder struct {
	project *project.Project
}

// NewProjectBuilder creates a new project builder with default values
func NewProjectBuilder() *ProjectBuilder {
	return &ProjectBuilder{
		project: &project.Project{
			ID:       uuid.New(),
			Name:     "Test Project",
			Services: make([]*service.Service, 0),
		},
	}
}

// WithID sets the project ID
func (b *ProjectBuilder) WithID(id uuid.UUID) *ProjectBuilder {
	b.project.ID = id
	return b
}

// WithName sets the project name
func (b *ProjectBuilder) WithName(name string) *ProjectBuilder {
	b.project.Name = name
	return b
}

// WithServices sets the project services
func (b *ProjectBuilder) WithServices(services []*service.Service) *ProjectBuilder {
	b.project.Services = services
	return b
}

// AddService adds a service to the project
func (b *ProjectBuilder) AddService(svc *service.Service) *ProjectBuilder {
	b.project.Services = append(b.project.Services, svc)
	return b
}

// Build returns the built project
func (b *ProjectBuilder) Build() *project.Project {
	return b.project
}

// ServiceBuilder helps build test services
type ServiceBuilder struct {
	service *service.Service
}

// NewServiceBuilder creates a new service builder with default values
func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		service: &service.Service{
			ID:   uuid.New(),
			Name: "Test Service",
		},
	}
}

// WithID sets the service ID
func (b *ServiceBuilder) WithID(id uuid.UUID) *ServiceBuilder {
	b.service.ID = id
	return b
}

// WithName sets the service name
func (b *ServiceBuilder) WithName(name string) *ServiceBuilder {
	b.service.Name = name
	return b
}

// WithResource sets the service resource
func (b *ServiceBuilder) WithResource(resource MockResource) *ServiceBuilder {
	b.service.Resource = &resource
	return b
}

// Build returns the built service
func (b *ServiceBuilder) Build() *service.Service {
	return b.service
}
