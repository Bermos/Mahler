package testutil

import (
	"time"
)

// MockResource is a mock implementation of the Resource interface for testing
type MockResource struct {
	NameValue        string
	DescriptionValue string
	ProvidesValue    []interface{}
	PriceValue       float64
	CPUMetricsValue  string
	MemMetricsValue  string
}

// Name returns the mock resource name
func (m *MockResource) Name() string {
	return m.NameValue
}

// Description returns the mock resource description
func (m *MockResource) Description() string {
	return m.DescriptionValue
}

// Provides returns what the mock resource provides
func (m *MockResource) Provides() []interface{} {
	return m.ProvidesValue
}

// Price returns the mock resource price for a given duration
func (m *MockResource) Price(duration time.Duration) float64 {
	return m.PriceValue
}

// MetricsCPU returns the mock CPU metrics identifier
func (m *MockResource) MetricsCPU() string {
	return m.CPUMetricsValue
}

// MetricsMemory returns the mock memory metrics identifier
func (m *MockResource) MetricsMemory() string {
	return m.MemMetricsValue
}

// NewMockResource creates a new mock resource with default values
func NewMockResource() *MockResource {
	return &MockResource{
		NameValue:        "Mock Resource",
		DescriptionValue: "A mock resource for testing",
		ProvidesValue:    []interface{}{"test"},
		PriceValue:       10.0,
		CPUMetricsValue:  "mock_cpu",
		MemMetricsValue:  "mock_memory",
	}
}

// WithName sets the name of the mock resource
func (m *MockResource) WithName(name string) *MockResource {
	m.NameValue = name
	return m
}

// WithDescription sets the description of the mock resource
func (m *MockResource) WithDescription(desc string) *MockResource {
	m.DescriptionValue = desc
	return m
}

// WithProvides sets what the mock resource provides
func (m *MockResource) WithProvides(provides []interface{}) *MockResource {
	m.ProvidesValue = provides
	return m
}

// WithPrice sets the price of the mock resource
func (m *MockResource) WithPrice(price float64) *MockResource {
	m.PriceValue = price
	return m
}
