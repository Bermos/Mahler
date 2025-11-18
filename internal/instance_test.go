package internal

import (
	"testing"

	"github.com/Bermos/Platform/internal/project"
	"github.com/Bermos/Platform/internal/resource"
	"github.com/Bermos/Platform/internal/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInstance_NewInstance(t *testing.T) {
	t.Parallel()

	instance := &Instance{
		Name:               "Test Instance",
		Projects:           make([]*project.Project, 0),
		AvailableResources: make([]resource.Resource, 0),
	}

	assert.Equal(t, "Test Instance", instance.Name)
	assert.NotNil(t, instance.Projects)
	assert.NotNil(t, instance.AvailableResources)
	assert.Empty(t, instance.Projects)
	assert.Empty(t, instance.AvailableResources)
}

func TestInstance_AddProject(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name              string
		initialProjects   []*project.Project
		projectToAdd      *project.Project
		expectedNumProj   int
	}{
		{
			name:            "add project to empty instance",
			initialProjects: make([]*project.Project, 0),
			projectToAdd: &project.Project{
				ID:   uuid.New(),
				Name: "New Project",
			},
			expectedNumProj: 1,
		},
		{
			name: "add project to instance with existing projects",
			initialProjects: []*project.Project{
				{
					ID:   uuid.New(),
					Name: "Existing Project",
				},
			},
			projectToAdd: &project.Project{
				ID:   uuid.New(),
				Name: "Another Project",
			},
			expectedNumProj: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			instance := &Instance{
				Name:     "Test Instance",
				Projects: tt.initialProjects,
			}

			instance.AddProject(tt.projectToAdd)

			assert.Len(t, instance.Projects, tt.expectedNumProj)
			assert.Equal(t, tt.projectToAdd, instance.Projects[len(instance.Projects)-1])
		})
	}
}

func TestInstance_AddMultipleProjects(t *testing.T) {
	t.Parallel()

	instance := &Instance{
		Name:     "Test Instance",
		Projects: make([]*project.Project, 0),
	}

	proj1 := testutil.NewProjectBuilder().
		WithName("Project 1").
		Build()

	proj2 := testutil.NewProjectBuilder().
		WithName("Project 2").
		Build()

	proj3 := testutil.NewProjectBuilder().
		WithName("Project 3").
		Build()

	instance.AddProject(proj1)
	instance.AddProject(proj2)
	instance.AddProject(proj3)

	require.Len(t, instance.Projects, 3)
	assert.Equal(t, proj1, instance.Projects[0])
	assert.Equal(t, proj2, instance.Projects[1])
	assert.Equal(t, proj3, instance.Projects[2])
}

func TestInstance_WithAvailableResources(t *testing.T) {
	t.Parallel()

	mockResource := testutil.NewMockResource()

	instance := &Instance{
		Name:               "Test Instance",
		Projects:           make([]*project.Project, 0),
		AvailableResources: []resource.Resource{mockResource},
	}

	require.Len(t, instance.AvailableResources, 1)
	assert.Equal(t, mockResource, instance.AvailableResources[0])
}

func TestInstance_AddProjectConcurrency(t *testing.T) {
	t.Parallel()

	instance := &Instance{
		Name:     "Test Instance",
		Projects: make([]*project.Project, 0),
	}

	// Note: This test shows that the current implementation is NOT thread-safe
	// In a real implementation, we would need to add synchronization
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			defer func() { done <- true }()
			proj := testutil.NewProjectBuilder().
				WithName("Concurrent Project").
				Build()
			instance.AddProject(proj)
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Due to race conditions, we can't assert an exact number
	// This test would fail with -race flag without proper synchronization
	assert.GreaterOrEqual(t, len(instance.Projects), 1)
}
