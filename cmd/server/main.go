package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"game-slots/internal/api"
	"game-slots/internal/storage"
)

func main() {
	store := storage.NewStorage()
	r := api.NewRouter(store)

	log.Info("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
