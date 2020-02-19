package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	controller := NewHealth()
	health := controller.CheckHealth()
	assert.True(t, health, "the container should be healthy if mongo is up")
}
