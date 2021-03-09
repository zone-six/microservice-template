package types

import (
	"encoding/json"
)

// Activity type
type Activity struct {
	ID string
}

// Athlete type
type Athlete struct {
	//ID is UserID
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Deleted   bool   `json:"deleted"`
}

// Settings type
type Settings struct {
	AthleteID string          `json:"athlete_id"`
	Metadata  json.RawMessage `json:"metadata"`
}

// Equipment type
type Equipment struct {
	ID string
}

// Workout type
type Workout struct {
}

// EquipmentData is the record of a given collection of data on a piece of equipment (e.g. a trainer)
type EquipmentData struct {
	Name        string
	ModelNumber string
	Link        string
	FitFile     []byte
}

// ActivityFilter type
type ActivityFilter struct {
	Take, Skip int
}

// WorkoutFilter type
type WorkoutFilter struct {
	Take, Skip int
}

// EquipmentFilter type
type EquipmentFilter struct {
	Take, Skip int
}
