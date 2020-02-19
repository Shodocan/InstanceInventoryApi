package entity

import "fmt"

type Instance struct {
	Hostname string `json:"hostname"`
	Disk     Metric `json:"disk"`
	CPU      Metric `json:"cpu"`
	Memory   Metric `json:"memory"`
}

type Metric struct {
	Mean   float64 `json:"mean"`
	Mode   float64 `json:"mode"`
	Median float64 `json:"median"`
	Unit   string  `json:"unit"`
}

type MetricField string

func (s MetricField) MongoField() string {
	return fmt.Sprintf("$%s", s)
}

var (
	MetricCPU    MetricField = "cpu_load"
	MetricDisk   MetricField = "disk_usage"
	MetricMemory MetricField = "memory_usage"
)
