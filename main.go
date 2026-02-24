package main

import (
	"context"
	"log"
	"os"

	"github.com/phalaxion/zent/cmd"
	"github.com/phalaxion/zent/ledger"
	"github.com/phalaxion/zent/store"
)

func main() {
	jsonStore := &store.JSONStore{
		FilePath: "ledger.json",
	}

	service := ledger.NewService(jsonStore)

	app := cmd.NewApp(service)

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
