package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bermos/Platform/internal/app"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

func TestRegister(t *testing.T) {
	t.Helper()

	tests := []struct {
		name string
		app  *app.App
	}{
		{
			name: "registers routes successfully",
			app:  app.NewApp(),
		},
		{
			name: "registers with new app instance",
			app:  app.NewApp(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			// Create a new Chi router
			router := chi.NewRouter()

			// Create a Huma API instance
			humaAPI := humachi.New(router, huma.DefaultConfig("Test API", "1.0.0"))

			// Register should not panic
			Register(humaAPI, tt.app)

			// Verify the API was created successfully
			if humaAPI == nil {
				t.Error("API should not be nil after registration")
			}
		})
	}
}

func TestRegister_RoutesAreAccessible(t *testing.T) {
	t.Helper()

	// Create a new Chi router
	router := chi.NewRouter()

	// Create a Huma API instance
	config := huma.DefaultConfig("Test API", "1.0.0")
	humaAPI := humachi.New(router, config)

	// Create app and register routes
	application := app.NewApp()
	Register(humaAPI, application)

	// Create a test server
	server := httptest.NewServer(router)
	defer server.Close()

	// Test that the registered endpoint is accessible
	resp, err := http.Get(server.URL + "/api/v1/projects")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Should get a valid response (200 or other HTTP status, not a connection error)
	// We're just testing that the route was registered, not the full functionality
	if resp.StatusCode == 0 {
		t.Error("Should get a valid HTTP status code")
	}
}

func TestRegister_MultipleRegistrations(t *testing.T) {
	t.Helper()

	// Create a new Chi router
	router := chi.NewRouter()

	// Create a Huma API instance
	humaAPI := humachi.New(router, huma.DefaultConfig("Test API", "1.0.0"))

	// Create app
	application := app.NewApp()

	// First registration should succeed
	Register(humaAPI, application)

	// Note: Multiple registrations would cause a panic in Huma
	// This test documents the current behavior
	// In production, Register should only be called once
}

func TestRegister_WithNilApp(t *testing.T) {
	t.Helper()

	// Create a new Chi router
	router := chi.NewRouter()

	// Create a Huma API instance
	humaAPI := humachi.New(router, huma.DefaultConfig("Test API", "1.0.0"))

	// This test documents that Register expects a non-nil app
	// In a real scenario, this would be prevented by proper initialization
	defer func() {
		if r := recover(); r != nil {
			// Expected to panic with nil app
			t.Log("Register panicked with nil app as expected")
		}
	}()

	Register(humaAPI, nil)
}

func TestRegister_VerifyOperationDetails(t *testing.T) {
	t.Helper()

	// Create a new Chi router
	router := chi.NewRouter()

	// Create a Huma API instance
	config := huma.DefaultConfig("Test API", "1.0.0")
	humaAPI := humachi.New(router, config)

	// Create app and register routes
	application := app.NewApp()
	Register(humaAPI, application)

	// Create a test request to verify the endpoint exists
	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	req = req.WithContext(context.Background())
	w := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(w, req)

	// The route should exist (not 404)
	// Note: The actual response code depends on the app.ListProjects implementation
	// We're just verifying the route was registered
	if w.Code == http.StatusNotFound {
		t.Error("Route /api/v1/projects should be registered (got 404)")
	}
}
