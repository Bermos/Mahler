package v1

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Bermos/Platform/internal/app"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegister(t *testing.T) {
	t.Parallel()

	// Create a test HTTP server
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Test API", "1.0.0"))

	// Create app instance
	testApp := app.NewApp()

	// Register routes
	Register(api, testApp)

	// Verify the API was configured (by making a request)
	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	// The endpoint should exist (not 404)
	assert.NotEqual(t, http.StatusNotFound, w.Code)
}

func TestRegister_ListProjectsEndpoint(t *testing.T) {
	t.Parallel()

	// Create a test HTTP server
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Test API", "1.0.0"))

	// Create app instance
	testApp := app.NewApp()

	// Register routes
	Register(api, testApp)

	// Make a request to the ListProjects endpoint
	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Should succeed (200 OK or similar)
	assert.True(t, resp.StatusCode >= 200 && resp.StatusCode < 300, "Expected success status code, got %d", resp.StatusCode)

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.NotNil(t, body)
}

func TestRegister_WrongHTTPMethod(t *testing.T) {
	t.Parallel()

	// Create a test HTTP server
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Test API", "1.0.0"))

	// Create app instance
	testApp := app.NewApp()

	// Register routes
	Register(api, testApp)

	// Try POST instead of GET (should not be allowed for ListProjects)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Should return method not allowed or similar error
	assert.NotEqual(t, http.StatusOK, resp.StatusCode)
}

func TestRegister_InvalidPath(t *testing.T) {
	t.Parallel()

	// Create a test HTTP server
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Test API", "1.0.0"))

	// Create app instance
	testApp := app.NewApp()

	// Register routes
	Register(api, testApp)

	// Request an invalid path
	req := httptest.NewRequest(http.MethodGet, "/api/v1/nonexistent", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	// Should return 404 Not Found
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

// mockApp is a mock implementation of the App for testing
type mockApp struct {
	listProjectsCalled bool
}

func (m *mockApp) ListProjects(ctx context.Context, input *struct{}) (*struct{}, error) {
	m.listProjectsCalled = true
	return &struct{}{}, nil
}

func TestRegister_CallsAppMethod(t *testing.T) {
	t.Parallel()

	// Create a test HTTP server
	mux := http.NewServeMux()
	api := humago.New(mux, huma.DefaultConfig("Test API", "1.0.0"))

	// Create mock app
	mockApp := &mockApp{listProjectsCalled: false}

	// Register with Huma
	huma.Register(api, huma.Operation{
		OperationID: "ListProjects",
		Description: "List all projects",
		Method:      http.MethodGet,
		Path:        "/api/v1/projects",
		Tags:        []string{"projects"},
	}, mockApp.ListProjects)

	// Make a request
	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	// Verify the app method was called
	assert.True(t, mockApp.listProjectsCalled, "Expected ListProjects to be called")
}
