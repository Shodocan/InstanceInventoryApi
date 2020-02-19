package main

import (
	"log"
	"net/http"

	"github.com/Shodocan/InstanceInventoryApi/internal/handler/rest"
	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/instance", rest.NewInstanceHandler().Handle)
	r.Get("/health", rest.NewHealthHandler().Handle)
	log.Panic(http.ListenAndServe(":80", r))
}
