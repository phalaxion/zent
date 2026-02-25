package cmd

import (
	"context"
	"fmt"

	"github.com/phalaxion/zent/internal"
	"github.com/phalaxion/zent/ledger"
	"github.com/urfave/cli/v3"
)

func listCommand(service *ledger.Service) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "list transactions",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			transactions, err := service.List()
			if err != nil {
				return err
			}

			if len(transactions) == 0 {
				fmt.Println("No transactions recorded.")
				return nil
			}

			for _, t := range transactions {
				fmt.Printf(
					"%s | %s | %10s | %s\n",
					t.ID,
					t.Timestamp.Format("2006-01-02 15:04:05"),
					internal.FormatCurrency(t.Amount),
					t.Description,
				)
			}

			return nil
		},
	}
}
