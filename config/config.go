package config

import (
	"os"

	"github.com/go-pg/pg/v9"
)

type Config struct {
	FakerAddr  string
	PgConnOpts pg.Options
}

var Default = Config{
	FakerAddr: "http://faker-http:3114",
	PgConnOpts: pg.Options{
		Addr:     "postgres:5432",
		Database: "dorothy",
		Password: os.Getenv("POSTGRES_PASSWORD"),
		User:     "postgres",
	},
}
