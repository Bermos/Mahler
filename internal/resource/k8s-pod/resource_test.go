package k8s_pod

import (
	"testing"
	"time"

	"github.com/Bermos/Platform/internal/resource"
)

func TestSetup(t *testing.T) {
	t.Helper()

	tests := []struct {
		name string
	}{
		{name: "creates a new Pod resource"},
		{name: "returns resource.Resource interface"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			result := Setup()

			if result == nil {
				t.Error("Setup() should return non-nil resource")
			}

			// Verify it implements the Resource interface
			var _ resource.Resource = result

			// Verify it's actually a Pod
			_, ok := result.(*Pod)
			if !ok {
				t.Error("Setup() should return a *Pod")
			}
		})
	}
}

func TestSetup_ReturnsNewInstance(t *testing.T) {
	t.Helper()

	r1 := Setup()
	r2 := Setup()

	if r1 == r2 {
		t.Error("Setup() should return different instances on each call")
	}
}

func TestPod_Name(t *testing.T) {
	t.Helper()

	pod := &Pod{}
	name := pod.Name()

	if name == "" {
		t.Error("Name() should return non-empty string")
	}

	expectedName := "Kubernetes Pod"
	if name != expectedName {
		t.Errorf("Name() = %q, want %q", name, expectedName)
	}
}

func TestPod_Description(t *testing.T) {
	t.Helper()

	pod := &Pod{}
	description := pod.Description()

	if description == "" {
		t.Error("Description() should return non-empty string")
	}

	// Check that it contains key terms
	if len(description) < 10 {
		t.Errorf("Description() seems too short: %q", description)
	}
}

func TestPod_Provides(t *testing.T) {
	t.Helper()

	tests := []struct {
		name string
		pod  *Pod
	}{
		{
			name: "default pod",
			pod:  &Pod{},
		},
		{
			name: "pod with price",
			pod:  &Pod{pricePerHour: 1.5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			provides := tt.pod.Provides()

			if provides == nil {
				t.Error("Provides() should return non-nil slice")
			}

			// Currently returns empty slice
			if len(provides) != 0 {
				t.Errorf("Provides() = %v, want empty slice", provides)
			}
		})
	}
}

func TestPod_Price(t *testing.T) {
	t.Helper()

	tests := []struct {
		name         string
		pricePerHour float64
		interval     time.Duration
		want         float64
	}{
		{
			name:         "zero price",
			pricePerHour: 0,
			interval:     1 * time.Hour,
			want:         0,
		},
		{
			name:         "one hour at $1/hour",
			pricePerHour: 1.0,
			interval:     1 * time.Hour,
			want:         1.0,
		},
		{
			name:         "two hours at $1/hour",
			pricePerHour: 1.0,
			interval:     2 * time.Hour,
			want:         2.0,
		},
		{
			name:         "30 minutes at $2/hour",
			pricePerHour: 2.0,
			interval:     30 * time.Minute,
			want:         1.0,
		},
		{
			name:         "24 hours at $0.50/hour",
			pricePerHour: 0.50,
			interval:     24 * time.Hour,
			want:         12.0,
		},
		{
			name:         "one day at $10/hour",
			pricePerHour: 10.0,
			interval:     24 * time.Hour,
			want:         240.0,
		},
		{
			name:         "15 minutes at $4/hour",
			pricePerHour: 4.0,
			interval:     15 * time.Minute,
			want:         1.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			pod := &Pod{pricePerHour: tt.pricePerHour}
			got := pod.Price(tt.interval)

			// Use a small epsilon for floating point comparison
			epsilon := 0.0001
			if diff := got - tt.want; diff < -epsilon || diff > epsilon {
				t.Errorf("Price() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPod_Price_ZeroDuration(t *testing.T) {
	t.Helper()

	pod := &Pod{pricePerHour: 10.0}
	price := pod.Price(0)

	if price != 0 {
		t.Errorf("Price(0) = %v, want 0", price)
	}
}

func TestPod_MetricsCPU(t *testing.T) {
	t.Helper()

	tests := []struct {
		name string
		pod  *Pod
	}{
		{
			name: "default pod",
			pod:  &Pod{},
		},
		{
			name: "pod with price",
			pod:  &Pod{pricePerHour: 5.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			metrics := tt.pod.MetricsCPU()

			if metrics == "" {
				t.Error("MetricsCPU() should return non-empty string")
			}

			// Should return a Prometheus metric name
			expectedMetric := "container_cpu_usage_seconds_total"
			if metrics != expectedMetric {
				t.Errorf("MetricsCPU() = %q, want %q", metrics, expectedMetric)
			}
		})
	}
}

func TestPod_MetricsMemory(t *testing.T) {
	t.Helper()

	tests := []struct {
		name string
		pod  *Pod
	}{
		{
			name: "default pod",
			pod:  &Pod{},
		},
		{
			name: "pod with price",
			pod:  &Pod{pricePerHour: 5.0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			metrics := tt.pod.MetricsMemory()

			if metrics == "" {
				t.Error("MetricsMemory() should return non-empty string")
			}

			// Should return a Prometheus metric name
			expectedMetric := "container_memory_usage_bytes"
			if metrics != expectedMetric {
				t.Errorf("MetricsMemory() = %q, want %q", metrics, expectedMetric)
			}
		})
	}
}

func TestPod_ImplementsResourceInterface(t *testing.T) {
	t.Helper()

	var _ resource.Resource = (*Pod)(nil)
	var _ resource.Resource = &Pod{}
	var _ resource.Resource = Setup()
}

func TestPod_AllMethods(t *testing.T) {
	t.Helper()

	// Integration test to ensure all methods work together
	pod := Setup()

	// Test all methods don't panic
	_ = pod.Name()
	_ = pod.Description()
	_ = pod.Provides()
	_ = pod.Price(1 * time.Hour)
	_ = pod.MetricsCPU()
	_ = pod.MetricsMemory()
}
