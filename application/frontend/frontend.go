package frontend

import (
	"github.com/tfpf/riven/application/backend"
	"github.com/tfpf/riven/config"
)

type frontend struct {
	cfg *config.Config
	be  backend.Backend
}

// Frontend handles user-facing content.
type Frontend interface {
	// MainLoop does stuff.
	MainLoop()
}

// NewFrontend returns an object to handle user-facing content.
func NewFrontend(cfg *config.Config, be backend.Backend) Frontend {
	return &frontend{
		cfg: cfg,
		be:  be,
	}
}

func (fe *frontend) MainLoop() {
}
