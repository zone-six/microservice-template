package workout

import (
	"context"

	"github.com/zone-six/microservice-template/internal/accessors"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/engines"
	"github.com/zone-six/microservice-template/internal/managers/types"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Manager is the manager for workout
type Manager interface {
	// CreateWorkout takes in a fit file, returning the ID of the created workout
	CreateWorkout(ctx context.Context, file []byte) (string, error)
	// UpdateWorkout updates the workout detail entry and the file associated with the workout
	UpdateWorkout(ctx context.Context, file []byte, workout types.Workout) error
	// GetWorkout gets the file associated with the workoutID
	GetWorkout(ctx context.Context, workoutID string) ([]byte, error)
	// GetWorkouts returns a list of workout available workouts with their meta data
	GetWorkouts(ctx context.Context, workoutID string, filter types.WorkoutFilter) ([]*types.Workout, error)
	// DeleteWorkout sets the workout as inactive
	DeleteWorkout(ctx context.Context, workoutID string) error
}

type workoutManager struct {
	config    *config.Config
	accessors *accessors.Container
	engines   *engines.Container
	utilities *utilities.Container
}

// New returns a new workout Manager
func New(config *config.Config, accessors *accessors.Container, engines *engines.Container, utilities *utilities.Container) Manager {
	return &workoutManager{config: config, accessors: accessors, engines: engines, utilities: utilities}
}

// CreateWorkout saves the workout file and returns and ID for the given workout
func (m *workoutManager) CreateWorkout(ctx context.Context, file []byte) (string, error) {
	// TODO: Access control - Admin role?
	// strip some metadata off of the .fit file
	// Have workout accessor create the workout, returning the ID
	// Save the fit file to the workouts directory using ID as the name
	panic("not implemented")
}

// UpdateWorkout updates the workout detail entry and the file associated with the workout
func (m *workoutManager) UpdateWorkout(ctx context.Context, file []byte, workout types.Workout) error {
	// TODO: Access control - Admin role?
	// strip some metadata off of the .fit file
	// Have workout accessor update the workout
	// Save the fit file to the workouts directory using ID as the name
	panic("not implemented")
}

// GetWorkout gets the file associated with the workoutID
func (m *workoutManager) GetWorkout(ctx context.Context, workoutID string) ([]byte, error) {
	// Get the file using the file accessor
	panic("not implemented")
}

// TODO: look at using a filter here for getting the workouts based on role?

// GetWorkouts returns a list of workout available workouts with their meta data
func (m *workoutManager) GetWorkouts(ctx context.Context, workoutID string, filter types.WorkoutFilter) ([]*types.Workout, error) {
	// Get workouts using the workout accessor
	panic("not implemented")
}

// DeleteWorkout sets the workout as inactive
func (m *workoutManager) DeleteWorkout(ctx context.Context, workoutID string) error {
	// Delete (set inactive) the workout
	panic("not implemented")
}
