package repository

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ServerCPUMean      float64 = 0.5045931010673991
	ServerMemoryMean   float64 = 4.933911914220367
	ServerDiskMean     float64 = 26.46780679555713
	ServerCPUMedian    float64 = 0.5151895166058071
	ServerMemoryMedian float64 = 4.884661008470738
	ServerDiskMedian   float64 = 25.728163898979478
	ServerCPUMode      float64 = 0.5499999999999999
	ServerMemoryMode   float64 = 5.016666666666667
	ServerDiskMode     float64 = 32.16142857142857
)

func checkMongo(t *testing.T) bool {
	if os.Getenv("DB_HOST") == "" {
		return false
	}
	return true
}

func TestRepositoryInstanceStatistics(t *testing.T) {
	if !checkMongo(t) {
		t.Skip()
		return
	}
	repository, err := NewInstanceRepository()
	assert.Nil(t, err, "err should be nil", err)
	stats := repository.Statistics("server3", "cpu_load", "memory_usage", "disk_usage")
	assert.Equal(t, ServerCPUMean, stats.CPU.Mean, "CPU Mean is not right")
	assert.Equal(t, ServerCPUMedian, stats.CPU.Median, "CPU Median is not right")
	assert.Equal(t, ServerCPUMode, stats.CPU.Mode, "CPU Mode is not right")
	assert.Equal(t, ServerMemoryMean, stats.Memory.Mean, "Memory Mean is not right")
	assert.Equal(t, ServerMemoryMedian, stats.Memory.Median, "Memory Median is not right")
	assert.Equal(t, ServerMemoryMode, stats.Memory.Mode, "Memory Mode is not right")
	assert.Equal(t, ServerDiskMean, stats.Disk.Mean, "Disk Mean is not right")
	assert.Equal(t, ServerDiskMedian, stats.Disk.Median, "Disk Median is not right")
	assert.Equal(t, ServerDiskMode, stats.Disk.Mode, "Disk Mode is not right")
}
func TestRepositoryInstanceStatisticsUnknownInstance(t *testing.T) {
	if !checkMongo(t) {
		t.Skip()
		return
	}
	repository, err := NewInstanceRepository()
	assert.Nil(t, err, "err should be nil", err)
	stats := repository.Statistics("server3333", "cpu_load", "memory_usage", "disk_usage")
	assert.Equal(t, float64(0), stats.CPU.Mean, "CPU Mean is not right")
	assert.Equal(t, float64(0), stats.CPU.Median, "CPU Median is not right")
	assert.Equal(t, float64(0), stats.CPU.Mode, "CPU Mode is not right")
	assert.Equal(t, float64(0), stats.Memory.Mean, "Memory Mean is not right")
	assert.Equal(t, float64(0), stats.Memory.Median, "Memory Median is not right")
	assert.Equal(t, float64(0), stats.Memory.Mode, "Memory Mode is not right")
	assert.Equal(t, float64(0), stats.Disk.Mean, "Disk Mean is not right")
	assert.Equal(t, float64(0), stats.Disk.Median, "Disk Median is not right")
	assert.Equal(t, float64(0), stats.Disk.Mode, "Disk Mode is not right")
}
