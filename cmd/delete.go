package cmd

import (
	"context"
	"fmt"

	"github.com/phalaxion/zent/ledger"
	"github.com/urfave/cli/v3"
)

func deleteCommand(service *ledger.Service) *cli.Command {
	return &cli.Command{
		Name:  "delete",
		Usage: "delete a transaction by ID",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			id := cmd.Args().First()
			if id == "" {
				return fmt.Errorf("transaction ID is required")
			}

			err := service.Delete(id)
			if err != nil {
				return err
			}

			fmt.Printf("Deleted transaction with ID %s\n", id)

			return nil
		},
	}
}
