package dataprocessing

import (
	"context"

	"github.com/zone-six/microservice-template/internal/accessors"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Engine is the data processing engine
type Engine interface {
	// ProcessEquipmentData processes a data record for a given piece of equipment
	ProcessTrainerData(ctx context.Context, equipmentDataRecordID string) error
}

type dataProcessingEngine struct {
	config    *config.Config
	accessors *accessors.Container
	utilities *utilities.Container
}

// New returns a new data processing Engine
func New(config *config.Config, accessors *accessors.Container, utilities *utilities.Container) Engine {
	return &dataProcessingEngine{config: config, accessors: accessors, utilities: utilities}
}

// TODO: Need to do some discovery on what to do if there is more than one data record for a specific trainer.

// ProcessEquipmentData processes a data record for a given piece of equipment
func (e *dataProcessingEngine) ProcessTrainerData(ctx context.Context, equipmentDataRecordID string) error {
	// Call file accessor to pull in the fit file.
	// Read fit file records for speed and power
	// Do linear regression/curve fitting on the speed and power
	// Call the equipment accessor to save the power curve data, marking the record status in some sort of "To be reviewed" state.
	panic("not implemented")
}
