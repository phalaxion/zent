package cmd

import (
	"context"

	"github.com/phalaxion/zent/ledger"
	"github.com/urfave/cli/v3"
)

func addCommand(service *ledger.Service) *cli.Command {
	return &cli.Command{
		Name:  "add",
		Usage: "add a transaction",
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:     "amount",
				Aliases:  []string{"a"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "desc",
				Aliases: []string{"d"},
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return service.Add(
				cmd.Float64("amount"),
				cmd.String("desc"),
			)
		},
	}
}
