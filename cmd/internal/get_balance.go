package internal

import (
	"context"

	"github.com/ry0suke17/solana-example/pkg/solana"
	"github.com/spf13/cobra"
)

type getBalanceCommand struct {
	*cobra.Command
	s       *solana.Solana
	address string
}

func (c *getBalanceCommand) run() error {
	resp, err := c.s.GetBalance(context.Background(), &solana.GetBalanceRequest{
		Address: c.address,
	})
	if err != nil {
		return err
	}
	c.Printf("balance: %d\n", resp.Value)
	return nil
}

func NewGetBalanceCommand(s *solana.Solana) *getBalanceCommand {
	var command = &cobra.Command{
		Use:   "get-balance",
		Short: "Get balance",
	}

	c := &getBalanceCommand{Command: command, s: s}
	c.Flags().StringVar(&c.address, "address", "", "address")
	c.MarkFlagRequired("address")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		return c.run()
	}

	return c
}
