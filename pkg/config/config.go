package config

import (
	"database/sql"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type (
	Config struct {
		Host       string `envconfig:"host" default:"0.0.0.0"`
		APIPort    string `envconfig:"api_port" default:"3000"`
		WSPort     string `envconfig:"ws_port" default:"3001"`
		DBUserName string `envconfig:"db_user_name" default:"chatapp"`
		DBPassword string `envconfig:"db_password" default:"chatapp"`
		DBHost     string `envconfig:"db_host" default:"0.0.0.0"`
		DBPort     string `envconfig:"db_port" default:":5432"`
		DBName     string `envconfig:"db_name" default:"chatapp_dev"`
	}
)

const dbServer = "chatapp"

// New constructs a new configuration.
func New() *Config {
	var c Config
	envconfig.MustProcess("chatapp", &c)
	return &c
}

func (config *Config) GetDBConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s%s/%s?sslmode=disable",
		config.DBUserName,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
}

func (config *Config) GetDBConnection(connStr string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	sql.Register(dbServer, &pq.Driver{})
	db, err = sql.Open(dbServer, connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrapf(err,
			"failed to connect to %s",
		)
	}
	return db, nil
}
