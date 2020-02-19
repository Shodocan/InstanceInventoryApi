package rest

import "net/http"

type HealthHandler struct {
}

func NewHealthHandler() Handler {
	return new(HealthHandler)
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
	w.WriteHeader(200)
}
