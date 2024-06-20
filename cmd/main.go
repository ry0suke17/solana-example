package main

import (
	"log"

	"github.com/ry0suke17/solana-example/cmd/internal"
	"github.com/ry0suke17/solana-example/pkg/solana"
	"github.com/spf13/cobra"
)

func run() error {
	s := solana.NewSolana()

	var rootCmd = &cobra.Command{}
	rootCmd.AddCommand(
		internal.NewGenKeyPairCommand(s).Command,
		internal.NewGetPublicKeyCommand(s).Command,
		internal.NewRequestAirdropCommand(s).Command,
	)

	return rootCmd.Execute()
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
