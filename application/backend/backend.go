package backend

import (
	"github.com/tfpf/riven/config"
)

type backend struct {
	cfg *config.Config
}

// Backend handles internal logic.
type Backend interface {
}

// Backend returns an object to handle internal logic.
func NewBackend(cfg *config.Config) Backend {
	return &backend{
		cfg: cfg,
	}
}
