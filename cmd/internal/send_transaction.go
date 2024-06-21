package internal

import (
	"context"

	"github.com/ry0suke17/solana-example/pkg/solana"
	"github.com/spf13/cobra"
)

type sendTransactionCommand struct {
	*cobra.Command
	s                  *solana.Solana
	privateKeyFilePath string
	addressFrom        string
	addressTo          string
	amount             uint64
}

func (c *sendTransactionCommand) run() error {
	resp, err := c.s.SendTransaction(context.Background(), &solana.SendTransactionRequest{
		PrivateKeyFilePath: c.privateKeyFilePath,
		AddressFrom:        c.addressFrom,
		AddressTo:          c.addressTo,
		Amount:             c.amount,
	})
	if err != nil {
		return err
	}
	c.Printf("signature: %s\n", resp.Signature)
	return nil
}

func NewSendTransactionCommand(s *solana.Solana) *sendTransactionCommand {
	var command = &cobra.Command{
		Use:   "send-transaction",
		Short: "Send transaction",
	}

	c := &sendTransactionCommand{Command: command, s: s}
	c.Flags().StringVar(&c.privateKeyFilePath, "private-key-file-path", "", "private key file path")
	c.MarkFlagRequired("private-key-file-path")
	c.Flags().StringVar(&c.addressFrom, "address-from", "", "send from address")
	c.MarkFlagRequired("address-from")
	c.Flags().StringVar(&c.addressTo, "address-to", "", "send to address")
	c.MarkFlagRequired("address-to")
	c.Flags().Uint64Var(&c.amount, "amount", 0, "amount. If send 0.01 SOL, value is `0,01`. Not lamports.")
	c.MarkFlagRequired("amount")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		return c.run()
	}

	return c
}
