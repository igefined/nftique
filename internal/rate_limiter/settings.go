package rate_limiter

import (
	"errors"
	"sync"
)

type Settings struct {
	mu       sync.RWMutex
	settings map[RequestType]*settings
}

type settings struct {
	Rate      float64
	MaxTokens float64
}

func NewSettings() *Settings {
	return &Settings{
		mu:       sync.RWMutex{},
		settings: make(map[RequestType]*settings),
	}
}

func (s *Settings) Set(reqType RequestType, rate, maxTokens float64) error {
	if rate < 1 {
		return errors.New("rate should be greater than 1")
	}

	if maxTokens < 1 {
		return errors.New("maxTokens should be greater than 1")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.settings[reqType] = &settings{rate, maxTokens}

	return nil
}

func (s *Settings) get(requestType RequestType) *settings {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.settings[requestType]
}
