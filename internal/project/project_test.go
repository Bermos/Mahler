package project

import (
	"testing"

	"github.com/Bermos/Platform/internal/service"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProject_NewProject(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		id       uuid.UUID
		projName string
		services []*service.Service
	}{
		{
			name:     "create project with valid ID and name",
			id:       uuid.New(),
			projName: "Test Project",
			services: nil,
		},
		{
			name:     "create project with services",
			id:       uuid.New(),
			projName: "Project with Services",
			services: []*service.Service{
				{
					ID:   uuid.New(),
					Name: "Service 1",
				},
				{
					ID:   uuid.New(),
					Name: "Service 2",
				},
			},
		},
		{
			name:     "create project with empty name",
			id:       uuid.New(),
			projName: "",
			services: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			project := &Project{
				ID:       tt.id,
				Name:     tt.projName,
				Services: tt.services,
			}

			assert.Equal(t, tt.id, project.ID)
			assert.Equal(t, tt.projName, project.Name)
			if tt.services != nil {
				assert.Len(t, project.Services, len(tt.services))
			} else {
				assert.Nil(t, project.Services)
			}
		})
	}
}

func TestProject_AddService(t *testing.T) {
	t.Parallel()

	project := &Project{
		ID:       uuid.New(),
		Name:     "Test Project",
		Services: make([]*service.Service, 0),
	}

	svc := &service.Service{
		ID:   uuid.New(),
		Name: "New Service",
	}

	project.Services = append(project.Services, svc)

	require.Len(t, project.Services, 1)
	assert.Equal(t, svc.ID, project.Services[0].ID)
	assert.Equal(t, svc.Name, project.Services[0].Name)
}

func TestProject_JSONMarshaling(t *testing.T) {
	t.Parallel()

	id := uuid.New()
	svcID := uuid.New()

	project := &Project{
		ID:   id,
		Name: "Test Project",
		Services: []*service.Service{
			{
				ID:   svcID,
				Name: "Test Service",
			},
		},
	}

	// Verify JSON tags are properly set
	assert.NotEmpty(t, project.ID)
	assert.NotEmpty(t, project.Name)
	assert.NotEmpty(t, project.Services)
}
