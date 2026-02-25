package cmd

import (
	"context"
	"fmt"

	"github.com/phalaxion/zent/internal/format"
	"github.com/phalaxion/zent/ledger"
	"github.com/urfave/cli/v3"
)

func getCommand(service *ledger.Service) *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "get a transaction by ID",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			id := cmd.Args().First()
			if id == "" {
				return fmt.Errorf("transaction ID is required")
			}

			transaction, err := service.Get(id)
			if err != nil {
				return err
			}

			fmt.Printf(
				"%s | %s | %10s | %s\n",
				transaction.ID,
				transaction.Timestamp.Format("2006-01-02 15:04:05"),
				format.Currency(transaction.Amount),
				transaction.Description,
			)

			return nil
		},
	}
}
