package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/zone-six/microservice-template/internal/clients/graph/generated"
	"github.com/zone-six/microservice-template/internal/clients/graph/types"
)

func (r *queryResolver) GetAthlete(ctx context.Context) (*types.Athlete, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
