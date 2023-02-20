package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/tabo-syu/bookmarks/infrastructures"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("seeding failed: %s", err.Error())
	}

	log.Print("seeding succeeded!")
}

func run() error {
	db, err := infrastructures.NewSQLHandler()
	if err != nil {
		return err
	}
	defer db.Close()

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	seed, err := os.Open(filepath.Join(wd, "cmd", "seeder", "seed.sql"))
	if err != nil {
		return err
	}
	defer seed.Close()

	seedQuery, err := io.ReadAll(seed)
	if err != nil {
		return err
	}

	_, err = db.ExecContext(context.Background(), string(seedQuery))
	if err != nil {
		return err
	}

	return nil
}
