package rest

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Shodocan/InstanceInventoryApi/internal/entity"
	"github.com/Shodocan/InstanceInventoryApi/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestPostRequest(t *testing.T) {
	handler := NewInstanceHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	jsonPayload := util.ToJSON(map[string]interface{}{"hostname": "server"})

	resp, err := http.Post(server.URL+"/instances", "application/json", strings.NewReader(jsonPayload))
	assert.Nil(t, err, "Should sucessfully request server")
	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err, "Should sucessfully decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response should be 200")
	assert.JSONEq(t, util.ToJSON(new(entity.Instance)), string(responseData), "Json should be empty")
}

func TestPostRequestValidInstance(t *testing.T) {
	handler := NewInstanceHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	jsonPayload := util.ToJSON(map[string]interface{}{"hostname": "server5"})

	resp, err := http.Post(server.URL+"/instances", "application/json", strings.NewReader(jsonPayload))
	assert.Nil(t, err, "Should sucessfully request server")
	responseData, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err, "Should sucessfully decode response body")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response should be 200")
	assert.NotEqual(t, util.ToJSON(new(entity.Instance)), string(responseData), "Json should be empty")
}

func TestPostRequestInvalidRequest(t *testing.T) {
	handler := NewInstanceHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	jsonPayload := util.ToJSON(map[string]interface{}{"hostname": "server5"})

	resp, err := http.Post(server.URL+"/instances", "application/json", strings.NewReader(jsonPayload[3:]))
	assert.Nil(t, err, "Should sucessfully request server")
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Response should be 400")
}

func TestPostRequestInvalidRequest2(t *testing.T) {
	handler := NewInstanceHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	jsonPayload := util.ToJSON(map[string]interface{}{"hostname": ""})

	resp, err := http.Post(server.URL+"/instances", "application/json", strings.NewReader(jsonPayload))
	assert.Nil(t, err, "Should sucessfully request server")
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Response should be 400")
}

func TestPostRequestInvalidRequest3(t *testing.T) {
	handler := NewInstanceHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	jsonPayload := util.ToJSON(map[string]interface{}{})

	resp, err := http.Post(server.URL+"/instances", "application/json", strings.NewReader(jsonPayload))
	assert.Nil(t, err, "Should sucessfully request server")
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode, "Response should be 400")
}

func TestPostRequestInvalidMethod(t *testing.T) {
	handler := NewInstanceHandler()
	server := httptest.NewServer(http.HandlerFunc(handler.Handle))
	defer server.Close()
	resp, err := http.Get(server.URL + "/instances")
	assert.Nil(t, err, "Should sucessfully request server")
	assert.Equal(t, http.StatusNotFound, resp.StatusCode, "Response should be 404")
}
