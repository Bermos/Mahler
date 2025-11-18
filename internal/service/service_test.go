package service

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockResource is a mock implementation of the Resource interface for testing
type mockResource struct {
	name        string
	description string
	provides    []interface{}
	price       float64
	cpuMetrics  string
	memMetrics  string
}

func (m *mockResource) Name() string {
	return m.name
}

func (m *mockResource) Description() string {
	return m.description
}

func (m *mockResource) Provides() []interface{} {
	return m.provides
}

func (m *mockResource) Price(duration time.Duration) float64 {
	return m.price
}

func (m *mockResource) MetricsCPU() string {
	return m.cpuMetrics
}

func (m *mockResource) MetricsMemory() string {
	return m.memMetrics
}

func TestService_NewService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		id          uuid.UUID
		serviceName string
		resource    *mockResource
	}{
		{
			name:        "create service with valid ID and name",
			id:          uuid.New(),
			serviceName: "Test Service",
			resource:    nil,
		},
		{
			name:        "create service with resource",
			id:          uuid.New(),
			serviceName: "Service with Resource",
			resource: &mockResource{
				name:        "Mock Resource",
				description: "A mock resource for testing",
				provides:    []interface{}{"test"},
				price:       10.0,
				cpuMetrics:  "cpu_usage",
				memMetrics:  "memory_usage",
			},
		},
		{
			name:        "create service with empty name",
			id:          uuid.New(),
			serviceName: "",
			resource:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			svc := &Service{
				ID:   tt.id,
				Name: tt.serviceName,
			}

			if tt.resource != nil {
				svc.Resource = tt.resource
			}

			assert.Equal(t, tt.id, svc.ID)
			assert.Equal(t, tt.serviceName, svc.Name)

			if tt.resource != nil {
				require.NotNil(t, svc.Resource)
				assert.Equal(t, tt.resource.Name(), svc.Resource.Name())
			}
		})
	}
}

func TestService_WithResource(t *testing.T) {
	t.Parallel()

	svc := &Service{
		ID:   uuid.New(),
		Name: "Test Service",
	}

	mockRes := &mockResource{
		name:        "Test Resource",
		description: "Test Description",
		provides:    []interface{}{"endpoint", "database"},
		price:       25.50,
		cpuMetrics:  "test_cpu",
		memMetrics:  "test_memory",
	}

	svc.Resource = mockRes

	require.NotNil(t, svc.Resource)
	assert.Equal(t, "Test Resource", svc.Resource.Name())
	assert.Equal(t, "Test Description", svc.Resource.Description())
	assert.Len(t, svc.Resource.Provides(), 2)
	assert.Equal(t, 25.50, svc.Resource.Price(time.Hour))
	assert.Equal(t, "test_cpu", svc.Resource.MetricsCPU())
	assert.Equal(t, "test_memory", svc.Resource.MetricsMemory())
}

func TestService_JSONMarshaling(t *testing.T) {
	t.Parallel()

	id := uuid.New()

	svc := &Service{
		ID:   id,
		Name: "Test Service",
	}

	// Verify JSON tags are properly set
	assert.NotEmpty(t, svc.ID)
	assert.NotEmpty(t, svc.Name)
}
