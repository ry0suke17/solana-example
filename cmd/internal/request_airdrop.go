package internal

import (
	"context"

	"github.com/ry0suke17/solana-example/pkg/solana"
	"github.com/spf13/cobra"
)

type requestAirdropCommand struct {
	*cobra.Command
	s       *solana.Solana
	address string
	amount  uint64
}

func (c *requestAirdropCommand) run() error {
	signature, err := c.s.RequestAirdrop(context.Background(), c.address, c.amount)
	if err != nil {
		return err
	}
	c.Printf("signature: %s\n", signature)
	return nil
}

func NewRequestAirdropCommand(s *solana.Solana) *requestAirdropCommand {
	var command = &cobra.Command{
		Use:   "request-airdrop",
		Short: "Request airdrop",
	}

	c := &requestAirdropCommand{Command: command, s: s}
	c.Flags().StringVar(&c.address, "address", "", "send to address")
	c.MarkFlagRequired("address")
	c.Flags().Uint64Var(&c.amount, "amount of lamports", 0, "amount. If send 1 SOL, value is `1000000000`.")
	c.MarkFlagRequired("amount")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		return c.run()
	}

	return c
}
