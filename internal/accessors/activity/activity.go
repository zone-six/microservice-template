package activity

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/zone-six/microservice-template/internal/accessors/activity/types"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Accessor is the store activitys
type Accessor interface {
	// ReadActivity reads activities by userID
	ReadActivities(ctx context.Context, userID string, filter types.Filter) ([]*types.Activity, error)
	// ReadActivity reads an activity by its ID
	ReadActivity(ctx context.Context, activityID string) (*types.Activity, error)
	// Create activity creates the activity returning the activities ID
	CreateActivity(ctx context.Context, activity types.Activity) (string, error)
	// Update activity updates the activity
	UpdateActivity(ctx context.Context, activity types.Activity) error
	// SetActivityInactive marks an activity as inactive
	SetActivityInactive(ctx context.Context, activityID string) error
}

type activityAccessor struct {
	db        *sqlx.DB
	config    *config.Config
	utilities *utilities.Container
}

// New returns a new activity accessor
func New(db *sqlx.DB, config *config.Config, utilities *utilities.Container) Accessor {
	return &activityAccessor{db: db, config: config, utilities: utilities}
}

// ReadActivity reads activities by userID
func (a *activityAccessor) ReadActivities(ctx context.Context, userID string, filter types.Filter) ([]*types.Activity, error) {
	panic("Not Implemented")
}

// ReadActivity reads an activity by its ID
func (a *activityAccessor) ReadActivity(ctx context.Context, activityID string) (*types.Activity, error) {
	panic("Not Implemented")
}

// Create activity creates the activity returning the activities ID
func (a *activityAccessor) CreateActivity(ctx context.Context, activity types.Activity) (string, error) {
	panic("Not Implemented")
}

// Update activity updates the activity
func (a *activityAccessor) UpdateActivity(ctx context.Context, activity types.Activity) error {
	panic("Not Implemented")
}

// SetActivityInactive sets an activity as deleted. It does not truely delete the activity
func (a *activityAccessor) SetActivityInactive(ctx context.Context, activityID string) error {
	panic("Not Implemented")
}
