package graph

import "github.com/zone-six/microservice-template/internal/managers"

// NOTE: May need to run go get github.com/99designs/gqlgen before you can run this command.
//go:generate go run github.com/99designs/gqlgen --verbose

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the schema resolver
type Resolver struct {
	Managers *managers.Container
}
