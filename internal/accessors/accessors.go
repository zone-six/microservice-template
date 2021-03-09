package accessors

import (
	"github.com/jmoiron/sqlx"
	"github.com/zone-six/microservice-template/internal/accessors/activity"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Container for the accessors
type Container struct {
	Activity activity.Accessor
}

// New creates all the accessors
func New(cfg *config.Config, db *sqlx.DB, utilities *utilities.Container) *Container {
	return &Container{
		Activity: activity.New(db, cfg, utilities),
	}
}
