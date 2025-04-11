package database

import (
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

func RunMigration() {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(DB, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Migration error:", err)
	}

	fmt.Println("Migration success, applied", n, "migrations!")
}
