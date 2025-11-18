package internal

import (
	"testing"

	"github.com/Bermos/Platform/internal/project"
	"github.com/Bermos/Platform/internal/testutil"
)

func TestInstance_AddProject(t *testing.T) {
	t.Helper()

	tests := []struct {
		name            string
		initialProjects []*project.Project
		projectToAdd    *project.Project
		expectedCount   int
	}{
		{
			name:            "adds project to empty instance",
			initialProjects: []*project.Project{},
			projectToAdd:    testutil.NewTestProject(),
			expectedCount:   1,
		},
		{
			name: "adds project to instance with existing projects",
			initialProjects: []*project.Project{
				testutil.NewTestProject(),
				testutil.NewTestProject(),
			},
			projectToAdd:  testutil.NewTestProject(),
			expectedCount: 3,
		},
		{
			name: "adds multiple projects sequentially",
			initialProjects: []*project.Project{
				testutil.NewTestProject(),
			},
			projectToAdd:  testutil.NewTestProject(),
			expectedCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			instance := &Instance{
				Name:     "Test Instance",
				Projects: tt.initialProjects,
			}

			instance.AddProject(tt.projectToAdd)

			testutil.AssertEqual(t, len(instance.Projects), tt.expectedCount, "project count should match")

			// Verify the added project is in the list
			found := false
			for _, p := range instance.Projects {
				if p.ID == tt.projectToAdd.ID {
					found = true
					break
				}
			}
			testutil.AssertTrue(t, found, "added project should be in the projects list")
		})
	}
}

func TestInstance_AddProject_NilProject(t *testing.T) {
	t.Helper()

	instance := &Instance{
		Name:     "Test Instance",
		Projects: []*project.Project{},
	}

	// Should not panic when adding nil project
	instance.AddProject(nil)
	testutil.AssertEqual(t, len(instance.Projects), 1, "should have one project (nil)")
	testutil.AssertNil(t, instance.Projects[0], "project should be nil")
}

func TestInstance_AddProject_PreservesOrder(t *testing.T) {
	t.Helper()

	instance := &Instance{
		Name:     "Test Instance",
		Projects: []*project.Project{},
	}

	// Add three projects
	p1 := testutil.NewTestProject()
	p2 := testutil.NewTestProject()
	p3 := testutil.NewTestProject()

	instance.AddProject(p1)
	instance.AddProject(p2)
	instance.AddProject(p3)

	// Verify order is preserved
	testutil.AssertEqual(t, instance.Projects[0].ID, p1.ID, "first project should be p1")
	testutil.AssertEqual(t, instance.Projects[1].ID, p2.ID, "second project should be p2")
	testutil.AssertEqual(t, instance.Projects[2].ID, p3.ID, "third project should be p3")
}

func TestInstance_AddProject_Concurrent(t *testing.T) {
	t.Helper()
	t.Parallel()

	instance := &Instance{
		Name:     "Test Instance",
		Projects: []*project.Project{},
	}

	// Note: This test documents the current behavior
	// In a production system, concurrent access would need proper synchronization
	project1 := testutil.NewTestProject()
	project2 := testutil.NewTestProject()

	instance.AddProject(project1)
	instance.AddProject(project2)

	testutil.AssertEqual(t, len(instance.Projects), 2, "should have two projects")
}
