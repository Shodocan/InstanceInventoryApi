package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Shodocan/InstanceInventoryApi/internal/controller"
	"github.com/Shodocan/InstanceInventoryApi/internal/entity"
	"github.com/Shodocan/InstanceInventoryApi/internal/entity/request"
	"github.com/Shodocan/InstanceInventoryApi/internal/util"
)

func NewInstanceHandler() Handler {
	return &InstanceHandler{}
}

type InstanceHandler struct {
	*GenericHandler
}

func (h *InstanceHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.post(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *InstanceHandler) post(w http.ResponseWriter, r *http.Request) {
	instanceRequest := &request.Instance{}
	err := json.NewDecoder(r.Body).Decode(&instanceRequest)
	if err != nil {
		h.printErr(err, http.StatusBadRequest, "Failed to parse request body", w)
		return
	}
	if instanceRequest.Hostname == "" {
		h.printErr(err, http.StatusBadRequest, "Request body without hostname attr", w)
		return
	}
	h.answerStatistics(instanceRequest, w)
}

func (h *InstanceHandler) answerStatistics(instanceRequest *request.Instance, w http.ResponseWriter) {
	statistics, err := controller.NewStatistics()
	if err != nil {
		h.printErr(err, http.StatusInternalServerError, "Failed to start controller", w)
		return
	}
	instanceData := statistics.GetStatistics(instanceRequest.Hostname,
		entity.MetricCPU, entity.MetricDisk, entity.MetricMemory)
	h.ok([]byte(util.ToJSON(instanceData)), w)
}
