package internal

import (
	"fmt"

	"github.com/ry0suke17/solana-example/pkg/solana"
	"github.com/spf13/cobra"
)

type GetPublicKeyCommand struct {
	*cobra.Command
	s                  *solana.Solana
	privateKeyFilePath string
}

func (c *GetPublicKeyCommand) run() error {
	publicKey, err := c.s.GetPublicKey(c.privateKeyFilePath)
	if err != nil {
		return err
	}
	fmt.Printf("public key: %s\n", publicKey)
	return nil
}

func NewGetPublicKeyCommand(s *solana.Solana) *GetPublicKeyCommand {
	var command = &cobra.Command{
		Use:   "get-public-key",
		Short: "Get public key",
	}

	c := &GetPublicKeyCommand{Command: command, s: s}
	c.Flags().StringVar(&c.privateKeyFilePath, "private-key-file-path", "", "file path to save private key")
	c.MarkFlagRequired("private-key-file-path")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		return c.run()
	}

	return c
}
