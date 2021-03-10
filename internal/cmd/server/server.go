package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/zone-six/microservice-template/internal/accessors"
	"github.com/zone-six/microservice-template/internal/clients"
	"github.com/zone-six/microservice-template/internal/config"
	"github.com/zone-six/microservice-template/internal/db"
	"github.com/zone-six/microservice-template/internal/engines"
	"github.com/zone-six/microservice-template/internal/managers"
	"github.com/zone-six/microservice-template/internal/utilities"
)

func main() {
	cfg, err := config.NewDefaultConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	dbconn, err := db.NewDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to db")
	}

	defer func() {
		err = dbconn.Close()
		if err != nil {
			log.Error().Err(err).Msg("failed to close the db connection")
		}
	}()

	app := New(dbconn, cfg)
	app.Run()
}

// App represents the top level application
type App struct {
	config  *config.Config
	Clients *clients.Container
}

// New creates a new instance of the application
func New(db *sqlx.DB, config *config.Config) *App {
	utilities := utilities.New(config)
	accessors := accessors.New(config, db, utilities)
	engines := engines.New(config, accessors, utilities)
	managers := managers.New(config, engines, accessors, utilities)
	clients := clients.New(config, managers, utilities)
	return &App{config: config, Clients: clients}
}

// Run starts the app
func (a *App) Run() {
	port := a.config.Port

	// Listen for shutdown signals.
	go a.listenForSignals()
	// Start any subscriptions
	a.Clients.PubSub.RegisterSubscriptions()

	r := mux.NewRouter()

	// Register middleware here

	// Register client handlers here
	a.Clients.Graph.RegisterHandlers(r)
	a.Clients.Rest.RegisterHandlers(r)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}

func (a *App) listenForSignals() {
	var captureSignal = make(chan os.Signal, 1)
	signal.Notify(captureSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	a.signalHandler(<-captureSignal)
}

func (a *App) signalHandler(signal os.Signal) {
	fmt.Printf("\nCaught signal: %+v", signal)
	fmt.Println("\nWait for 1 second to finish processing")

	switch signal {

	case syscall.SIGHUP: // kill -SIGHUP XXXX
		fmt.Println("- got hungup")

	case syscall.SIGINT: // kill -SIGINT XXXX or Ctrl+c
		fmt.Println("- got ctrl+c")

	case syscall.SIGTERM: // kill -SIGTERM XXXX
		fmt.Println("- got force stop")

	case syscall.SIGQUIT: // kill -SIGQUIT XXXX
		fmt.Println("- stop and core dump")

	default:
		fmt.Println("- unknown signal")
	}

	// Cleanup any subscriptions.
	a.Clients.PubSub.CleanUp()

	fmt.Println("\nFinishing server cleanup")
	os.Exit(0)
}
