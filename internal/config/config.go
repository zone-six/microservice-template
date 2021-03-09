package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	// ErrMissingEnvironmentStage missing stage configuration
	ErrMissingEnvironmentStage = errors.New("Missing Stage ENV Variable")

	// ErrMissingEnvironmentBranch missing branch configuration
	ErrMissingEnvironmentBranch = errors.New("Missing Branch ENV Variable")
)

// Config for the environment
type Config struct {
	Debug              bool   `envconfig:"DEBUG"`
	Stage              string `envconfig:"STAGE" default:"dev"`
	Branch             string `envconfig:"BRANCH"`
	Port               string `envconfig:"PORT" default:"8080"`
	PGConnectionString string `envconfig:"PGCONNECTIONSTRING"`
	DbSecrets          string `envconfig:"DB_SECRET"`
}

// DBSecrets for the DB
type DBSecrets struct {
	Password string `json:"password,omitempty"`
	DBName   string `json:"dbname,omitempty"`
	Port     string `json:"port,omitempty"`
	Host     string `json:"host,omitempty"`
	UserName string `json:"username,omitempty"`
}

// NewDefaultConfig reads configuration from environment variables and validates it
func NewDefaultConfig() (*Config, error) {
	cfg := new(Config)

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse environment config")
	}

	err = cfg.parseDbSecrets()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse DBSecrets")
	}

	err = cfg.validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed validation of config")
	}

	err = cfg.logging()
	if err != nil {
		return nil, errors.Wrap(err, "failed setup logging based on config")
	}

	log.Info().Str("stage", cfg.Stage).Bool("debug", cfg.Debug).Msg("logging configured")
	log.Info().Str("stage", cfg.Stage).Str("branch", cfg.Branch).Msg("Configuration loaded")

	return cfg, nil
}

func (cfg *Config) logging() error {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	if cfg.Stage == "local" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	return nil
}

func (cfg *Config) parseDbSecrets() error {
	dbsecrets := &DBSecrets{}
	err := json.Unmarshal([]byte(cfg.DbSecrets), dbsecrets)

	if err != nil {
		return err
	}

	cfg.PGConnectionString = fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbsecrets.Port,
		dbsecrets.Host,
		dbsecrets.UserName,
		dbsecrets.Password,
		dbsecrets.DBName)

	log.Debug().Str("PGConnectionString", cfg.PGConnectionString).Msg("configured PG Connection String")

	return nil
}

func (cfg *Config) validate() error {
	if cfg.Stage == "" {
		return ErrMissingEnvironmentStage
	}
	if cfg.Branch == "" {
		return ErrMissingEnvironmentBranch
	}
	return nil
}
