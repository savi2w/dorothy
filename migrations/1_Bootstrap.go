package main

import (
	"fmt"

	"github.com/go-pg/migrations/v7"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		fmt.Println("Migrating Dorothy 1_Bootstrap...")

		_, err := db.Exec(`
			CREATE TABLE users (
				id VARCHAR,
				role VARCHAR UNIQUE,
				PRIMARY KEY(id)
			);
		`)

		return err
	}, func(db migrations.DB) error {
		fmt.Println("Dropping Dorothy 1_Bootstrap...")

		_, err := db.Exec(("DROP TABLE users;"))
		return err
	})
}
