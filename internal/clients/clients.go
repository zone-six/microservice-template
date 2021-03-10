package clients

import (
	"github.com/zone-six/microservice-template/internal/clients/graph"
	"github.com/zone-six/microservice-template/internal/clients/pubsub"
	"github.com/zone-six/microservice-template/internal/clients/rest"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/managers"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Container for the clients
type Container struct {
	Graph  graph.Client
	Rest   rest.Client
	PubSub pubsub.Client
}

// New creates all the clients
func New(cfg *config.Config, managers *managers.Container, utilities *utilities.Container) *Container {
	return &Container{
		Graph:  graph.New(cfg, managers, utilities),
		Rest:   rest.New(cfg, managers, utilities),
		PubSub: pubsub.New(cfg, managers, utilities),
	}
}
