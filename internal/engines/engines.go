package engines

import (
	"github.com/zone-six/microservice-template/internal/accessors"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/engines/dataprocessing"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Container for the engines
type Container struct {
	DataProcessing dataprocessing.Engine
}

// New creates all the engines
func New(cfg *config.Config, accessors *accessors.Container, utilities *utilities.Container) *Container {
	return &Container{
		DataProcessing: dataprocessing.New(cfg, accessors, utilities),
	}
}
