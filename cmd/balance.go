package cmd

import (
	"context"
	"fmt"

	"github.com/phalaxion/zent/internal"
	"github.com/phalaxion/zent/ledger"
	"github.com/urfave/cli/v3"
)

func balanceCommand(service *ledger.Service) *cli.Command {
	return &cli.Command{
		Name:  "balance",
		Usage: "view current balance",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			balance, err := service.Balance()
			if err != nil {
				return err
			}

			fmt.Println("Current balance:", internal.FormatCurrency(balance))
			return nil
		},
	}
}
