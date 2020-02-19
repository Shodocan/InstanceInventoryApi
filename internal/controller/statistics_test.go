package controller

import (
	"testing"

	"github.com/Shodocan/InstanceInventoryApi/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestStatistics(t *testing.T) {
	controller, err := NewStatistics()
	assert.Nil(t, err, "Should have no error")
	ins := controller.GetStatistics("server4", entity.MetricCPU)
	assert.True(t, float64(0) < ins.CPU.Mean)
	assert.True(t, float64(0) < ins.CPU.Median)
	assert.True(t, float64(0) < ins.CPU.Mode)
	assert.Equal(t, "%", ins.CPU.Unit)
}
