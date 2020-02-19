package controller

import (
	"github.com/Shodocan/InstanceInventoryApi/internal/entity"
	"github.com/Shodocan/InstanceInventoryApi/internal/repository"
)

type Statistics struct {
	repository repository.InstanceRepository
}

func NewStatistics() (*Statistics, error) {
	repo, err := repository.NewInstanceRepository()
	return &Statistics{repository: repo}, err
}

func (s *Statistics) GetStatistics(hostname string, fields ...entity.MetricField) *entity.Instance {
	return s.repository.Statistics(hostname, fields...)
}
