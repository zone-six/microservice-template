// Package microservice provides subjects and types for interacting with the messages that the user
// microservice publishes.
package microservice

const (
	// SubjectNewUser is the subject under which new user messages will be published.
	// Message type is microservice.NewUser.
	SubjectNewUser = "zone-six.user.new"
)

// NewUser represents a new user
type NewUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
