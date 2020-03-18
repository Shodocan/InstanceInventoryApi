package rest

import (
	"net/http"
	"time"
)

type HealthHandler struct {
	startTime time.Time
}

func NewHealthHandler() Handler {
	handler := new(HealthHandler)
	handler.startTime = time.Now()
	return handler
}

func (h *HealthHandler) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.get(w)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *HealthHandler) get(w http.ResponseWriter) {
	if time.Since(h.startTime) > 10*time.Second {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}
