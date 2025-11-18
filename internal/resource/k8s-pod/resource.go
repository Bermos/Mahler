package k8s_pod

import (
	"github.com/Bermos/Platform/internal/resource"
	"time"
)

func Setup() resource.Resource {
	return &Pod{}
}

type Pod struct {
	pricePerHour float64
}

func (p *Pod) Name() string {
	return "Kubernetes Pod"
}

func (p *Pod) Description() string {
	return "A Kubernetes Pod is a group of one or more containers, with shared storage/network resources, and a specification for how to run the containers."
}

func (p *Pod) Provides() []interface{} {
	return []interface{}{}
}

func (p *Pod) Price(interval time.Duration) float64 {
	return p.pricePerHour * interval.Hours()
}

func (p *Pod) MetricsCPU() string {
	// TODO: implement proper Kubernetes metrics integration
	return "container_cpu_usage_seconds_total"
}

func (p *Pod) MetricsMemory() string {
	// TODO: implement proper Kubernetes metrics integration
	return "container_memory_working_set_bytes"
}
