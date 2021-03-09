package controllers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/zone-six/microservice-template/internal/clients/rest/models"
	"github.com/zone-six/microservice-template/internal/clients/rest/restapi/operations/todos"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/managers"
	"github.com/zone-six/microservice-template/internal/utilities"
)

// NOTE: This is a flawed implementation and is only for example.
var items = make(map[int64]*models.Item)
var lastID int64 = 0

// ToDoController is the implementation for the restapi.TodosAPI interface
type ToDoController struct {
	config    *config.Config
	managers  *managers.Container
	utilities *utilities.Container
}

// NewToDoController returns a new todo controller
func NewToDoController(cfg *config.Config, managers *managers.Container, utilities *utilities.Container) *ToDoController {
	return &ToDoController{config: cfg, managers: managers, utilities: utilities}
}

// Get implementation
func (s *ToDoController) Get(ctx context.Context, params todos.GetParams) middleware.Responder {
	response := []*models.Item{}
	count := int32(0)
	for k, v := range items {
		if k >= *params.Since {
			response = append(response, v)
			count++
		}
		if count == *params.Limit {
			break
		}
	}
	return todos.NewGetOK().WithPayload(response)
}

// AddOne implementation
func (s *ToDoController) AddOne(ctx context.Context, params todos.AddOneParams) middleware.Responder {
	lastID++
	params.Body.ID = lastID
	items[lastID] = params.Body
	return todos.NewAddOneCreated().WithPayload(params.Body)
}

// DestroyOne implementation
func (s *ToDoController) DestroyOne(ctx context.Context, params todos.DestroyOneParams) middleware.Responder {
	delete(items, params.ID)
	return todos.NewDestroyOneNoContent()
}

// UpdateOne implementation
func (s *ToDoController) UpdateOne(ctx context.Context, params todos.UpdateOneParams) middleware.Responder {
	_, ok := items[params.ID]
	if !ok {
		return todos.NewUpdateOneDefault(404)
	}

	items[params.ID] = params.Body
	return todos.NewUpdateOneOK().WithPayload(params.Body)
}
