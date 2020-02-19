package controller

import (
	"github.com/Shodocan/InstanceInventoryApi/internal/repository"
)

func NewHealth() *Health {
	return &Health{}
}

type Health struct {
}

func (h *Health) CheckHealth() bool {
	repo, err := repository.NewInstanceRepository()
	return !(err != nil || repo.Ping() != nil)
}
