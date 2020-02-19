package rest

import (
	"net/http"

	"github.com/Shodocan/InstanceInventoryApi/internal/logger"
)

type Handler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type GenericHandler struct{}

func (h *GenericHandler) ok(data []byte, w http.ResponseWriter) {
	h.write(http.StatusOK, data, w)
}

func (h *GenericHandler) write(status int, data []byte, w http.ResponseWriter) {
	writenBytes, err := w.Write(data)
	logger.Debugf("Written %d bytes on http response", writenBytes)
	if err != nil {
		h.printErr(err, http.StatusInternalServerError, "Error writing response", w)
		return
	}
	w.WriteHeader(status)
}

func (h *GenericHandler) printErr(err error, responseCode int, msg string, w http.ResponseWriter) {
	w.WriteHeader(responseCode)
	logger.Errorf(msg, err)
}
