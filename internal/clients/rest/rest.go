package rest

import (
	"embed"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zone-six/microservice-template/internal/clients/rest/controllers"
	"github.com/zone-six/microservice-template/internal/clients/rest/restapi"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/managers"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// Client is the Rest Client
type Client interface {
	RegisterHandlers(r *mux.Router)
}

type restClient struct {
	config    *config.Config
	managers  *managers.Container
	utilities *utilities.Container
}

// New returns a new rest client
func New(cfg *config.Config, managers *managers.Container, utilities *utilities.Container) Client {
	return &restClient{config: cfg, managers: managers, utilities: utilities}
}

//go:embed swaggerui
var staticFiles embed.FS

func (c *restClient) RegisterHandlers(r *mux.Router) {
	// TODO: Will need to do some work if you want to serve swagger and api under a subrouter.
	api := r.PathPrefix("/").Subrouter()

	// Initiate the http handler, with the objects that are implementing the business logic.
	h, err := restapi.Handler(restapi.Config{
		TodosAPI: controllers.NewToDoController(c.config, c.managers, c.utilities),
		Logger:   log.Printf,
	})
	if err != nil {
		log.Fatal(err)
	}
	api.Handle("/", h)

	if c.config.Stage != "prod" {
		// Note: This should be getting served by the restapi handler, but it's not for some reason
		r.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write(restapi.SwaggerJSON)
		})

		var staticFS = http.FS(staticFiles)
		r.PathPrefix("/swaggerui/").Handler(http.FileServer(staticFS))
	}
	api.Handle("/", api)
}
