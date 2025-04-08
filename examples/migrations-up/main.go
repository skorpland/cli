package main

import (
	"context"
	"io/fs"
	"log"
	"os"

	"github.com/skorpland/cli/pkg/migration"
	"github.com/skorpland/cli/pkg/pgxv5"
)

func main() {
	if err := migrate(context.Background(), os.DirFS(".")); err != nil {
		log.Fatalln(err)
	}
}

// Applies local migrations to a remote database, and tracks the history of executed statements.
func migrate(ctx context.Context, fsys fs.FS) error {
	conn, err := pgxv5.Connect(ctx, os.Getenv("POWERBASE_POSTGRES_URL"))
	if err != nil {
		return err
	}
	defer conn.Close(ctx)
	files, err := migration.ListLocalMigrations("powerbase/migrations", fsys)
	if err != nil {
		return err
	}
	return migration.ApplyMigrations(ctx, files, conn, fsys)
}
