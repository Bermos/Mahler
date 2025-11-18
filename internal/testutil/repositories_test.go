package testutil

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
)

func TestMockProjectRepository_Create(t *testing.T) {
	t.Run("creates_project_successfully", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Create(ctx, proj)
		AssertNoError(t, err, "should create project without error")

		// Verify project was created
		retrieved, err := repo.Get(ctx, proj.ID)
		AssertNoError(t, err, "should retrieve created project")
		AssertEqual(t, retrieved.ID, proj.ID, "retrieved project should have same ID")
	})

	t.Run("fails_when_duplicate_id", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Create(ctx, proj)
		AssertNoError(t, err, "first create should succeed")

		err = repo.Create(ctx, proj)
		AssertError(t, err, "duplicate create should fail")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockProjectRepository()
		repo.CreateError = errors.New("simulated error")
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Create(ctx, proj)
		AssertError(t, err, "should return configured error")
	})
}

func TestMockProjectRepository_Get(t *testing.T) {
	t.Run("retrieves_existing_project", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Create(ctx, proj)
		AssertNoError(t, err, "create should succeed")

		retrieved, err := repo.Get(ctx, proj.ID)
		AssertNoError(t, err, "get should succeed")
		AssertNotNil(t, retrieved, "retrieved project should not be nil")
		AssertEqual(t, retrieved.ID, proj.ID, "IDs should match")
	})

	t.Run("fails_for_nonexistent_project", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()
		nonexistentID := uuid.New()

		_, err := repo.Get(ctx, nonexistentID)
		AssertError(t, err, "should return error for nonexistent project")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockProjectRepository()
		repo.GetError = errors.New("simulated error")
		ctx := context.Background()

		_, err := repo.Get(ctx, uuid.New())
		AssertError(t, err, "should return configured error")
	})
}

func TestMockProjectRepository_Update(t *testing.T) {
	t.Run("updates_existing_project", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Create(ctx, proj)
		AssertNoError(t, err, "create should succeed")

		proj.Name = "updated-name"
		err = repo.Update(ctx, proj)
		AssertNoError(t, err, "update should succeed")

		retrieved, err := repo.Get(ctx, proj.ID)
		AssertNoError(t, err, "get should succeed")
		AssertEqual(t, retrieved.Name, "updated-name", "name should be updated")
	})

	t.Run("fails_for_nonexistent_project", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Update(ctx, proj)
		AssertError(t, err, "should return error for nonexistent project")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockProjectRepository()
		repo.UpdateError = errors.New("simulated error")
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Update(ctx, proj)
		AssertError(t, err, "should return configured error")
	})
}

func TestMockProjectRepository_Delete(t *testing.T) {
	t.Run("deletes_existing_project", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()
		proj := NewTestProject()

		err := repo.Create(ctx, proj)
		AssertNoError(t, err, "create should succeed")

		err = repo.Delete(ctx, proj.ID)
		AssertNoError(t, err, "delete should succeed")

		_, err = repo.Get(ctx, proj.ID)
		AssertError(t, err, "get should fail after delete")
	})

	t.Run("fails_for_nonexistent_project", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()

		err := repo.Delete(ctx, uuid.New())
		AssertError(t, err, "should return error for nonexistent project")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockProjectRepository()
		repo.DeleteError = errors.New("simulated error")
		ctx := context.Background()

		err := repo.Delete(ctx, uuid.New())
		AssertError(t, err, "should return configured error")
	})
}

func TestMockProjectRepository_List(t *testing.T) {
	t.Run("lists_all_projects", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()

		proj1 := NewProjectBuilder().WithName("project-1").Build()
		proj2 := NewProjectBuilder().WithName("project-2").Build()

		err := repo.Create(ctx, proj1)
		AssertNoError(t, err, "create project 1 should succeed")

		err = repo.Create(ctx, proj2)
		AssertNoError(t, err, "create project 2 should succeed")

		projects, err := repo.List(ctx)
		AssertNoError(t, err, "list should succeed")
		AssertEqual(t, len(projects), 2, "should list 2 projects")
	})

	t.Run("returns_empty_list_when_no_projects", func(t *testing.T) {
		repo := NewMockProjectRepository()
		ctx := context.Background()

		projects, err := repo.List(ctx)
		AssertNoError(t, err, "list should succeed")
		AssertNotNil(t, projects, "projects should not be nil")
		AssertEqual(t, len(projects), 0, "should return empty list")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockProjectRepository()
		repo.ListError = errors.New("simulated error")
		ctx := context.Background()

		_, err := repo.List(ctx)
		AssertError(t, err, "should return configured error")
	})
}

func TestMockServiceRepository_Create(t *testing.T) {
	t.Run("creates_service_successfully", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Create(ctx, svc)
		AssertNoError(t, err, "should create service without error")

		retrieved, err := repo.Get(ctx, svc.ID)
		AssertNoError(t, err, "should retrieve created service")
		AssertEqual(t, retrieved.ID, svc.ID, "retrieved service should have same ID")
	})

	t.Run("fails_when_duplicate_id", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Create(ctx, svc)
		AssertNoError(t, err, "first create should succeed")

		err = repo.Create(ctx, svc)
		AssertError(t, err, "duplicate create should fail")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockServiceRepository()
		repo.CreateError = errors.New("simulated error")
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Create(ctx, svc)
		AssertError(t, err, "should return configured error")
	})
}

func TestMockServiceRepository_Get(t *testing.T) {
	t.Run("retrieves_existing_service", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Create(ctx, svc)
		AssertNoError(t, err, "create should succeed")

		retrieved, err := repo.Get(ctx, svc.ID)
		AssertNoError(t, err, "get should succeed")
		AssertNotNil(t, retrieved, "retrieved service should not be nil")
		AssertEqual(t, retrieved.ID, svc.ID, "IDs should match")
	})

	t.Run("fails_for_nonexistent_service", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()

		_, err := repo.Get(ctx, uuid.New())
		AssertError(t, err, "should return error for nonexistent service")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockServiceRepository()
		repo.GetError = errors.New("simulated error")
		ctx := context.Background()

		_, err := repo.Get(ctx, uuid.New())
		AssertError(t, err, "should return configured error")
	})
}

func TestMockServiceRepository_Update(t *testing.T) {
	t.Run("updates_existing_service", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Create(ctx, svc)
		AssertNoError(t, err, "create should succeed")

		svc.Name = "updated-service"
		err = repo.Update(ctx, svc)
		AssertNoError(t, err, "update should succeed")

		retrieved, err := repo.Get(ctx, svc.ID)
		AssertNoError(t, err, "get should succeed")
		AssertEqual(t, retrieved.Name, "updated-service", "name should be updated")
	})

	t.Run("fails_for_nonexistent_service", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Update(ctx, svc)
		AssertError(t, err, "should return error for nonexistent service")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockServiceRepository()
		repo.UpdateError = errors.New("simulated error")
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Update(ctx, svc)
		AssertError(t, err, "should return configured error")
	})
}

func TestMockServiceRepository_Delete(t *testing.T) {
	t.Run("deletes_existing_service", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()
		svc := NewTestService()

		err := repo.Create(ctx, svc)
		AssertNoError(t, err, "create should succeed")

		err = repo.Delete(ctx, svc.ID)
		AssertNoError(t, err, "delete should succeed")

		_, err = repo.Get(ctx, svc.ID)
		AssertError(t, err, "get should fail after delete")
	})

	t.Run("fails_for_nonexistent_service", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()

		err := repo.Delete(ctx, uuid.New())
		AssertError(t, err, "should return error for nonexistent service")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockServiceRepository()
		repo.DeleteError = errors.New("simulated error")
		ctx := context.Background()

		err := repo.Delete(ctx, uuid.New())
		AssertError(t, err, "should return configured error")
	})
}

func TestMockServiceRepository_ListByProject(t *testing.T) {
	t.Run("lists_services", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()
		projectID := uuid.New()

		svc1 := NewServiceBuilder().WithName("service-1").Build()
		svc2 := NewServiceBuilder().WithName("service-2").Build()

		err := repo.Create(ctx, svc1)
		AssertNoError(t, err, "create service 1 should succeed")

		err = repo.Create(ctx, svc2)
		AssertNoError(t, err, "create service 2 should succeed")

		services, err := repo.ListByProject(ctx, projectID)
		AssertNoError(t, err, "list should succeed")
		AssertEqual(t, len(services), 2, "should list 2 services")
	})

	t.Run("returns_empty_list_when_no_services", func(t *testing.T) {
		repo := NewMockServiceRepository()
		ctx := context.Background()

		services, err := repo.ListByProject(ctx, uuid.New())
		AssertNoError(t, err, "list should succeed")
		AssertNotNil(t, services, "services should not be nil")
		AssertEqual(t, len(services), 0, "should return empty list")
	})

	t.Run("returns_error_when_configured", func(t *testing.T) {
		repo := NewMockServiceRepository()
		repo.ListByProjectError = errors.New("simulated error")
		ctx := context.Background()

		_, err := repo.ListByProject(ctx, uuid.New())
		AssertError(t, err, "should return configured error")
	})
}
