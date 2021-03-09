package utilities

import (
	"github.com/zone-six/microservice-template/internal/config"
)

// Container for the utilities
type Container struct {
}

// New creates all the utilities
func New(cfg *config.Config) *Container {
	return &Container{}
}
