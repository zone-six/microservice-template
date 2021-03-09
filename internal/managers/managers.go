package managers

import (
	"github.com/zone-six/microservice-template/internal/accessors"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/engines"
	"github.com/zone-six/microservice-template/internal/managers/workout"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Container for the managers
type Container struct {
	Workout workout.Manager
}

// New creates all the managers
func New(cfg *config.Config,
	engines *engines.Container,
	accessors *accessors.Container,
	utilities *utilities.Container) *Container {
	return &Container{
		Workout: workout.New(cfg, accessors, engines, utilities),
	}
}
