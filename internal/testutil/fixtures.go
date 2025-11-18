package testutil

import (
	"github.com/Bermos/Platform/internal/project"
	"github.com/Bermos/Platform/internal/resource"
	"github.com/Bermos/Platform/internal/service"
	"github.com/google/uuid"
)

// ProjectBuilder builds test Project instances using the builder pattern
type ProjectBuilder struct {
	id       uuid.UUID
	name     string
	services []*service.Service
}

// NewProjectBuilder creates a new ProjectBuilder with default values
func NewProjectBuilder() *ProjectBuilder {
	return &ProjectBuilder{
		id:       uuid.New(),
		name:     "test-project",
		services: make([]*service.Service, 0),
	}
}

// WithID sets the project ID
func (b *ProjectBuilder) WithID(id uuid.UUID) *ProjectBuilder {
	b.id = id
	return b
}

// WithName sets the project name
func (b *ProjectBuilder) WithName(name string) *ProjectBuilder {
	b.name = name
	return b
}

// WithServices sets the project services
func (b *ProjectBuilder) WithServices(services []*service.Service) *ProjectBuilder {
	b.services = services
	return b
}

// AddService adds a service to the project
func (b *ProjectBuilder) AddService(svc *service.Service) *ProjectBuilder {
	b.services = append(b.services, svc)
	return b
}

// Build creates the Project instance
func (b *ProjectBuilder) Build() *project.Project {
	return &project.Project{
		ID:       b.id,
		Name:     b.name,
		Services: b.services,
	}
}

// ServiceBuilder builds test Service instances using the builder pattern
type ServiceBuilder struct {
	id       uuid.UUID
	name     string
	resource resource.Resource
}

// NewServiceBuilder creates a new ServiceBuilder with default values
func NewServiceBuilder() *ServiceBuilder {
	return &ServiceBuilder{
		id:       uuid.New(),
		name:     "test-service",
		resource: NewMockResource(),
	}
}

// WithID sets the service ID
func (b *ServiceBuilder) WithID(id uuid.UUID) *ServiceBuilder {
	b.id = id
	return b
}

// WithName sets the service name
func (b *ServiceBuilder) WithName(name string) *ServiceBuilder {
	b.name = name
	return b
}

// WithResource sets the service resource
func (b *ServiceBuilder) WithResource(res resource.Resource) *ServiceBuilder {
	b.resource = res
	return b
}

// Build creates the Service instance
func (b *ServiceBuilder) Build() *service.Service {
	return &service.Service{
		ID:       b.id,
		Name:     b.name,
		Resource: b.resource,
	}
}

// Test Fixtures - commonly used test data

// NewTestProject creates a simple test project
func NewTestProject() *project.Project {
	return NewProjectBuilder().Build()
}

// NewTestProjectWithServices creates a test project with the given number of services
func NewTestProjectWithServices(numServices int) *project.Project {
	builder := NewProjectBuilder()
	for i := 0; i < numServices; i++ {
		svc := NewServiceBuilder().
			WithName("service-" + string(rune(i+'a'))).
			Build()
		builder.AddService(svc)
	}
	return builder.Build()
}

// NewTestService creates a simple test service
func NewTestService() *service.Service {
	return NewServiceBuilder().Build()
}

// NewTestServiceWithResource creates a test service with a specific resource
func NewTestServiceWithResource(res resource.Resource) *service.Service {
	return NewServiceBuilder().WithResource(res).Build()
}
