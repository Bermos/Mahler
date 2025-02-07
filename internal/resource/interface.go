package resource

import "time"

type Resource interface {
	Name() string
	Description() string
	Provides() []interface{}
	Price(duration time.Duration) float64
	MetricsCPU() string
	MetricsMemory() string
}
