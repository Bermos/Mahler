package testutil

import (
	"testing"

	"github.com/Bermos/Platform/internal/service"
	"github.com/google/uuid"
)

func TestProjectBuilder(t *testing.T) {
	t.Run("creates_project_with_defaults", func(t *testing.T) {
		builder := NewProjectBuilder()
		proj := builder.Build()

		AssertNotNil(t, proj, "project should not be nil")
		AssertEqual(t, proj.Name, "test-project", "default name should be set")
		AssertNotNil(t, proj.Services, "services should not be nil")
		AssertEqual(t, len(proj.Services), 0, "services should be empty by default")
	})

	t.Run("builder_with_custom_name", func(t *testing.T) {
		proj := NewProjectBuilder().
			WithName("my-project").
			Build()

		AssertEqual(t, proj.Name, "my-project", "custom name should be set")
	})

	t.Run("builder_with_custom_id", func(t *testing.T) {
		id := uuid.New()
		proj := NewProjectBuilder().
			WithID(id).
			Build()

		AssertEqual(t, proj.ID, id, "custom ID should be set")
	})

	t.Run("builder_with_services", func(t *testing.T) {
		svc1 := NewTestService()
		svc2 := NewTestService()
		services := []*service.Service{svc1, svc2}

		proj := NewProjectBuilder().
			WithServices(services).
			Build()

		AssertEqual(t, len(proj.Services), 2, "should have 2 services")
	})

	t.Run("builder_add_service", func(t *testing.T) {
		svc1 := NewTestService()
		svc2 := NewTestService()

		proj := NewProjectBuilder().
			AddService(svc1).
			AddService(svc2).
			Build()

		AssertEqual(t, len(proj.Services), 2, "should have 2 services")
	})
}

func TestServiceBuilder(t *testing.T) {
	t.Run("creates_service_with_defaults", func(t *testing.T) {
		builder := NewServiceBuilder()
		svc := builder.Build()

		AssertNotNil(t, svc, "service should not be nil")
		AssertEqual(t, svc.Name, "test-service", "default name should be set")
		AssertNotNil(t, svc.Resource, "resource should not be nil")
	})

	t.Run("builder_with_custom_name", func(t *testing.T) {
		svc := NewServiceBuilder().
			WithName("my-service").
			Build()

		AssertEqual(t, svc.Name, "my-service", "custom name should be set")
	})

	t.Run("builder_with_custom_id", func(t *testing.T) {
		id := uuid.New()
		svc := NewServiceBuilder().
			WithID(id).
			Build()

		AssertEqual(t, svc.ID, id, "custom ID should be set")
	})

	t.Run("builder_with_custom_resource", func(t *testing.T) {
		mockResource := NewMockResource().WithName("custom-resource")
		svc := NewServiceBuilder().
			WithResource(mockResource).
			Build()

		AssertNotNil(t, svc.Resource, "resource should not be nil")
		AssertEqual(t, svc.Resource.Name(), "custom-resource", "custom resource should be set")
	})
}

func TestNewTestProject(t *testing.T) {
	t.Run("creates_simple_project", func(t *testing.T) {
		proj := NewTestProject()

		AssertNotNil(t, proj, "project should not be nil")
		AssertEqual(t, proj.Name, "test-project", "should have default name")
		AssertNotNil(t, proj.Services, "services should not be nil")
	})
}

func TestNewTestProjectWithServices(t *testing.T) {
	tests := []struct {
		name        string
		numServices int
	}{
		{
			name:        "zero_services",
			numServices: 0,
		},
		{
			name:        "one_service",
			numServices: 1,
		},
		{
			name:        "multiple_services",
			numServices: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proj := NewTestProjectWithServices(tt.numServices)

			AssertNotNil(t, proj, "project should not be nil")
			AssertEqual(t, len(proj.Services), tt.numServices, "should have correct number of services")

			// Verify each service has a unique name
			if tt.numServices > 0 {
				for i, svc := range proj.Services {
					AssertNotNil(t, svc, "service should not be nil")
					expectedName := "service-" + string(rune(i+'a'))
					AssertEqual(t, svc.Name, expectedName, "service should have correct name")
				}
			}
		})
	}
}

func TestNewTestService(t *testing.T) {
	t.Run("creates_simple_service", func(t *testing.T) {
		svc := NewTestService()

		AssertNotNil(t, svc, "service should not be nil")
		AssertEqual(t, svc.Name, "test-service", "should have default name")
		AssertNotNil(t, svc.Resource, "resource should not be nil")
	})
}

func TestNewTestServiceWithResource(t *testing.T) {
	t.Run("creates_service_with_custom_resource", func(t *testing.T) {
		mockResource := NewMockResource().WithName("database")
		svc := NewTestServiceWithResource(mockResource)

		AssertNotNil(t, svc, "service should not be nil")
		AssertNotNil(t, svc.Resource, "resource should not be nil")
		AssertEqual(t, svc.Resource.Name(), "database", "should have custom resource")
	})
}
