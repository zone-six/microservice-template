package types

// Activity Type
type Activity struct {
	ID     string
	UserID string
}

// Filter Type
type Filter struct {
	Take, Skip int
}
