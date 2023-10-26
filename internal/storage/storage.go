package storage

import (
	"sync"

	"game-slots/internal/models"
)

type ReelStorer interface {
	GetReelConfiguration(name string) ([][]string, bool)
	SaveReelConfiguration(name string, config models.ReelConfiguration)
}

type LineStorer interface {
	GetLineConfiguration(name string) ([]models.LineConfiguration, bool)
	SaveLineConfiguration(name string, config []models.LineConfiguration)
}

type PayoutStorer interface {
	GetPayoutConfiguration(name string) ([]models.PayoutConfiguration, bool)
	SavePayoutConfiguration(name string, config []models.PayoutConfiguration)
}

// Storage структура представляет собой хранилище конфигураций
type Storage struct {
	Reels   map[string]models.ReelConfiguration
	Lines   map[string][]models.LineConfiguration
	Payouts map[string][]models.PayoutConfiguration
	mu      sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		Reels:   make(map[string]models.ReelConfiguration),
		Lines:   make(map[string][]models.LineConfiguration),
		Payouts: make(map[string][]models.PayoutConfiguration),
	}
}

func (s *Storage) SaveReelConfiguration(name string, config models.ReelConfiguration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Reels[name] = config
}

func (s *Storage) SaveLineConfiguration(name string, config []models.LineConfiguration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Lines[name] = config
}

func (s *Storage) SavePayoutConfiguration(name string, config []models.PayoutConfiguration) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Payouts[name] = config
}

func (s *Storage) GetReelConfiguration(name string) ([][]string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	reelConfig, ok := s.Reels[name]
	return reelConfig, ok
}

func (s *Storage) GetLineConfiguration(name string) ([]models.LineConfiguration, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	lineConfig, ok := s.Lines[name]
	return lineConfig, ok
}

func (s *Storage) GetPayoutConfiguration(name string) ([]models.PayoutConfiguration, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	payoutConfig, ok := s.Payouts[name]
	return payoutConfig, ok
}
