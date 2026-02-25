package main

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/phalaxion/zent/cmd"
	"github.com/phalaxion/zent/ledger"
	"github.com/phalaxion/zent/store"
)

func main() {
	dbPath := os.Getenv("ZENT_DB")

	if dbPath == "" {
		configDir, err := os.UserConfigDir()
		if err != nil {
			log.Fatal(err)
		}

		appDir := filepath.Join(configDir, "zent")

		if err := os.MkdirAll(appDir, 0755); err != nil {
			log.Fatal(err)
		}

		dbPath = filepath.Join(appDir, "ledger.db")
	}

	sqliteStore, err := store.NewSQLiteStore(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	service := ledger.NewService(sqliteStore)

	app := cmd.NewApp(service)

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
