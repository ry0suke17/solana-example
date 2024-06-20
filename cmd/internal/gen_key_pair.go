package internal

import (
	"fmt"

	"github.com/ry0suke17/solana-example/pkg/solana"
	"github.com/spf13/cobra"
)

type genKeyPairCommand struct {
	*cobra.Command
	s                  *solana.Solana
	privateKeyFilePath string
}

func (c *genKeyPairCommand) run() error {
	publicKey, err := c.s.GenKeyPair(c.privateKeyFilePath)
	if err != nil {
		return err
	}
	fmt.Printf("public key: %s\n", publicKey)
	return nil
}

func NewGenKeyPairCommand(s *solana.Solana) *genKeyPairCommand {
	var command = &cobra.Command{
		Use:   "gen-key-pair",
		Short: "Generate key pair",
	}

	c := &genKeyPairCommand{Command: command, s: s}
	c.Flags().StringVar(&c.privateKeyFilePath, "private-key-file-path", "", "file path to save private key")
	c.MarkFlagRequired("private-key-file-path")
	c.RunE = func(cmd *cobra.Command, args []string) error {
		return c.run()
	}

	return c
}
