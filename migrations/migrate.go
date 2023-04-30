package main

import (
	"flag"

	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
	"github.com/savi2w/dorothy/config"
)

func main() {
	flag.Parse()
	db := pg.Connect(&config.Default.PgConnOpts)
	if _, _, err := migrations.Run(db, flag.Args()...); err != nil {
		panic(err)
	}
}
