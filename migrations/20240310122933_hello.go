package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upHello, downHello)
}

func upHello(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downHello(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
