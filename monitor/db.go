package monitor

import (
	"context"

	"encore.dev/storage/sqldb"
)

// Define a database named 'url', using the database
// migrations  in the "./migrations" folder.
// Encore provisions, migrates, and connects to the database.
// Learn more: https://encore.dev/docs/primitives/databases
var db = sqldb.NewDatabase("events", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

type URL struct {
	ID         int
	AuthorName string
}

// insert inserts a URL into the database.
func insert(ctx context.Context, authorName string) error {
	_, err := db.Exec(ctx, `
        INSERT INTO events (author_name)
        VALUES ($1)
    `, authorName)
	return err
}
