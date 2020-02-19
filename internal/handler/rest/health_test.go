package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	handler := NewHealthHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	resp, err := http.Get(server.URL + "/health")
	assert.Nil(t, err, "Should sucessfully request server")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response should be 200")
}

func TestHealthInvalidMethod(t *testing.T) {
	handler := NewHealthHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	resp, err := http.Post(server.URL+"/health", "text/plain", nil)
	assert.Nil(t, err, "Should sucessfully request server")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Response should be 404")
}
