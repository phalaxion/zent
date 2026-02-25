package cmd

import (
	"github.com/phalaxion/zent/ledger"
	"github.com/urfave/cli/v3"
)

func NewApp(service *ledger.Service) *cli.Command {
	return &cli.Command{
		Commands: []*cli.Command{
			addCommand(service),
			balanceCommand(service),
			listCommand(service),
			getCommand(service),
		},
	}
}
