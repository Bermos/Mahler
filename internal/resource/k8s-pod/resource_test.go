package k8s_pod

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSetup(t *testing.T) {
	t.Parallel()

	resource := Setup()

	require.NotNil(t, resource)
	assert.IsType(t, &Pod{}, resource)
}

func TestPod_Name(t *testing.T) {
	t.Parallel()

	pod := &Pod{}
	name := pod.Name()

	assert.Equal(t, "Kubernetes Pod", name)
	assert.NotEmpty(t, name)
}

func TestPod_Description(t *testing.T) {
	t.Parallel()

	pod := &Pod{}
	desc := pod.Description()

	assert.NotEmpty(t, desc)
	assert.Contains(t, desc, "Kubernetes")
	assert.Contains(t, desc, "Pod")
}

func TestPod_Provides(t *testing.T) {
	t.Parallel()

	pod := &Pod{}
	provides := pod.Provides()

	require.NotNil(t, provides)
	assert.Empty(t, provides, "Pod should currently provide empty list")
}

func TestPod_Price(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name         string
		pricePerHour float64
		duration     time.Duration
		expectedCost float64
	}{
		{
			name:         "zero price per hour",
			pricePerHour: 0.0,
			duration:     time.Hour,
			expectedCost: 0.0,
		},
		{
			name:         "one hour duration",
			pricePerHour: 10.0,
			duration:     time.Hour,
			expectedCost: 10.0,
		},
		{
			name:         "half hour duration",
			pricePerHour: 10.0,
			duration:     30 * time.Minute,
			expectedCost: 5.0,
		},
		{
			name:         "24 hours duration",
			pricePerHour: 5.0,
			duration:     24 * time.Hour,
			expectedCost: 120.0,
		},
		{
			name:         "fractional price per hour",
			pricePerHour: 2.5,
			duration:     2 * time.Hour,
			expectedCost: 5.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			pod := &Pod{
				pricePerHour: tt.pricePerHour,
			}

			cost := pod.Price(tt.duration)
			assert.Equal(t, tt.expectedCost, cost)
		})
	}
}

func TestPod_MetricsCPU(t *testing.T) {
	t.Parallel()

	pod := &Pod{}
	metric := pod.MetricsCPU()

	assert.NotEmpty(t, metric)
	assert.Equal(t, "container_cpu_usage_seconds_total", metric)
}

func TestPod_MetricsMemory(t *testing.T) {
	t.Parallel()

	pod := &Pod{}
	metric := pod.MetricsMemory()

	assert.NotEmpty(t, metric)
	assert.Equal(t, "container_memory_working_set_bytes", metric)
}

func TestPod_InterfaceCompliance(t *testing.T) {
	t.Parallel()

	// This test verifies that Pod implements the Resource interface
	pod := Setup()

	assert.NotEmpty(t, pod.Name())
	assert.NotEmpty(t, pod.Description())
	assert.NotNil(t, pod.Provides())
	assert.GreaterOrEqual(t, pod.Price(time.Hour), 0.0)
	assert.NotEmpty(t, pod.MetricsCPU())
	assert.NotEmpty(t, pod.MetricsMemory())
}
