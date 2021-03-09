package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/zone-six/microservice-template/internal/clients/graph/generated"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/managers"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Client is the GraphQL Client
type Client interface {
	RegisterHandlers(r *mux.Router)
}

type graphClient struct {
	config    *config.Config
	managers  *managers.Container
	utilities *utilities.Container
}

// New returns a new graph client
func New(cfg *config.Config, managers *managers.Container, utilities *utilities.Container) Client {
	return &graphClient{config: cfg, managers: managers, utilities: utilities}
}

func (c *graphClient) RegisterHandlers(r *mux.Router) {
	gqlConfig := generated.Config{Resolvers: &Resolver{Managers: c.managers}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(gqlConfig))
	graph := r.PathPrefix("/graphql").Subrouter()
	if c.config.Stage != "prod" {
		graph.Handle("", playground.Handler("GraphQL playground", "/graphql/query"))
	}
	graph.Handle("/query", srv)
}
