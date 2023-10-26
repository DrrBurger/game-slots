package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"

	"game-slots/internal/metrics"
	"game-slots/internal/models"
	"game-slots/internal/storage"
	"game-slots/internal/utils"
	"game-slots/pkg/calculator"
)

// Handler структура, содержащая ссылку на хранилище
type Handler struct {
	store *storage.Storage
}

// NewHandler создает и возвращает новый обработчик с указанным хранилищем
func NewHandler(s *storage.Storage) *Handler {
	return &Handler{store: s}
}

// UploadConfig обрабатывает запрос на загрузку конфигурации
func (h *Handler) UploadConfig(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to upload configuration")

	configType, name, err := utils.ValidateRequest(r)
	if err != nil {
		log.Println(err.Error())
		metrics.RecordError(r.Method, r.URL.Path)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.decodeAndSaveConfig(configType, name, r.Body); err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Successfully saved %s configuration under name: %s", configType, name)
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write([]byte("Configuration uploaded successfully"))
}

func (h *Handler) decodeAndSaveConfig(configType, name string, body io.Reader) error {
	var err error
	switch configType {
	case "reels":
		var reelsConfig models.ReelConfiguration
		err = json.NewDecoder(body).Decode(&reelsConfig)
		if err == nil {
			h.store.SaveReelConfiguration(name, reelsConfig)
		}
	case "lines":
		var linesConfig []models.LineConfiguration
		err = json.NewDecoder(body).Decode(&linesConfig)
		if err == nil {
			h.store.SaveLineConfiguration(name, linesConfig)
		}
	case "payouts":
		var payoutConfig []models.PayoutConfiguration
		err = json.NewDecoder(body).Decode(&payoutConfig)
		if err == nil {
			h.store.SavePayoutConfiguration(name, payoutConfig)
		}
	default:
		return errors.New("invalid config type")
	}

	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	return nil
}

// Calculate обрабатывает запрос на расчет
func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to calculate")

	reelsName, linesName, payoutsName := r.URL.Query().Get("reels"), r.URL.Query().Get("lines"), r.URL.Query().Get("payouts")

	result, err := calculator.Calculate(h.store, reelsName, linesName, payoutsName)
	if err != nil {
		log.Printf("Error during calculation: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Calculate endpoint hit. Reels: %s, Lines: %s, Payouts: %s", reelsName, linesName, payoutsName)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		return
	}
}
