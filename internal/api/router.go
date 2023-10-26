package api

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"game-slots/internal/metrics"
	"game-slots/internal/storage"
)

// NewRouter создает и возвращает новый роутер с заданными обработчиками
func NewRouter(store *storage.Storage) *mux.Router {
	handler := NewHandler(store)
	r := mux.NewRouter()

	r.Use(metrics.InstrumentMiddleware)
	r.HandleFunc("/upload", handler.UploadConfig).Methods("POST")
	r.HandleFunc("/calculate", handler.Calculate).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	return r
}
