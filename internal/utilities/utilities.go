package utilities

import (
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/pubsub-utility"
)

// Container for the utilities
type Container struct {
	PubSub pubsub.Utility
}

// New creates all the utilities
func New(cfg *config.Config) *Container {
	pubsubOptions := pubsub.Options{
		ClusterID: cfg.NatsStreamingClusterID,
		ClientID:  cfg.NatsStreamingClientID,
		CertFile:  cfg.NatsCertFile,
		KeyFile:   cfg.NatsKeyFile,
		NatsURL:   cfg.NatsURL,
	}
	return &Container{
		PubSub: pubsub.NewPubSubUtility(pubsubOptions),
	}
}
