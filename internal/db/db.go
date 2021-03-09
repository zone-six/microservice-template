package db

import (
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/zone-six/microservice-template/internal/config"
)

// DefaultMaxOpenConnections the default value for max open connections in the
// PostgreSQL connection pool
const DefaultMaxOpenConnections = 30

// NewDB create a new DB pool and migrates the database
func NewDB(cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.PGConnectionString)

	if err != nil {
		return nil, errors.Wrap(err, "DB not available")
	}

	configureConnectionPool(db)

	log.Info().Msg("Starting Database Migrations")

	if err := DoMigrate(NewMigrate(db)); err != nil {
		return nil, errors.Wrap(err, "Failed to migrate the DB.")
	}

	log.Info().Msg("Database Migrations Complete")

	return db, nil
}

// configureConnectionPool sets reasonable sizes on the built in DB queue. By
// default the connection pool is unbounded, which leads to the error `pq:
// sorry too many clients already`.
func configureConnectionPool(db *sqlx.DB) {
	var err error
	maxOpen := DefaultMaxOpenConnections
	if e := os.Getenv("SRC_PGSQL_MAX_OPEN"); e != "" {
		maxOpen, err = strconv.Atoi(e)
		if err != nil {
			log.Fatal().Err(err).Msg("SRC_PGSQL_MAX_OPEN is not an int")
		}
	}
	db.SetMaxOpenConns(maxOpen)
	db.SetMaxIdleConns(maxOpen)
	db.SetConnMaxLifetime(time.Minute)
}
