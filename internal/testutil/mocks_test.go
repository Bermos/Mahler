package testutil

import (
	"testing"
	"time"
)

func TestNewMockResource(t *testing.T) {
	t.Run("creates_mock_with_defaults", func(t *testing.T) {
		mock := NewMockResource()

		AssertNotNil(t, mock, "mock resource should not be nil")
		AssertEqual(t, mock.Name(), "mock-resource", "default name should be set")
		AssertEqual(t, mock.Description(), "A mock resource for testing", "default description should be set")
		AssertNotNil(t, mock.Provides(), "provides should not be nil")
		AssertEqual(t, mock.Price(time.Hour), 10.0, "default price should be set")
		AssertEqual(t, mock.MetricsCPU(), "/metrics/cpu", "default CPU metrics should be set")
		AssertEqual(t, mock.MetricsMemory(), "/metrics/memory", "default memory metrics should be set")
	})
}

func TestMockResource_Builder(t *testing.T) {
	t.Run("builder_pattern_works", func(t *testing.T) {
		mock := NewMockResource().
			WithName("custom-resource").
			WithDescription("Custom description").
			WithPrice(25.5)

		AssertEqual(t, mock.Name(), "custom-resource", "name should be customized")
		AssertEqual(t, mock.Description(), "Custom description", "description should be customized")
		AssertEqual(t, mock.Price(time.Hour), 25.5, "price should be customized")
	})

	t.Run("with_provides", func(t *testing.T) {
		provides := []interface{}{"database", "cache"}
		mock := NewMockResource().WithProvides(provides)

		AssertNotNil(t, mock.Provides(), "provides should not be nil")
		AssertEqual(t, len(mock.Provides()), 2, "provides should have 2 items")
	})
}

func TestMockResource_Price(t *testing.T) {
	tests := []struct {
		name     string
		price    float64
		duration time.Duration
		want     float64
	}{
		{
			name:     "returns_set_price",
			price:    10.0,
			duration: time.Hour,
			want:     10.0,
		},
		{
			name:     "price_independent_of_duration",
			price:    20.0,
			duration: time.Hour * 24,
			want:     20.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := NewMockResource().WithPrice(tt.price)
			got := mock.Price(tt.duration)
			AssertEqual(t, got, tt.want, "price should match")
		})
	}
}

func TestMockResource_Metrics(t *testing.T) {
	t.Run("cpu_metrics_endpoint", func(t *testing.T) {
		mock := NewMockResource()
		mock.CPUMetrics = "/custom/cpu"

		AssertEqual(t, mock.MetricsCPU(), "/custom/cpu", "CPU metrics should be customizable")
	})

	t.Run("memory_metrics_endpoint", func(t *testing.T) {
		mock := NewMockResource()
		mock.MemoryMetrics = "/custom/memory"

		AssertEqual(t, mock.MetricsMemory(), "/custom/memory", "memory metrics should be customizable")
	})
}
