package repository

import "github.com/Shodocan/InstanceInventoryApi/internal/entity"

type InstanceRepository interface {
	Statistics(hostname string, field ...entity.MetricField) *entity.Instance
	Ping() error
}
