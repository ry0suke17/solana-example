package solana

import "github.com/gagliardetto/solana-go/rpc"

type Solana struct {
	client *rpc.Client
}

func NewSolana() *Solana {
	endpoint := rpc.DevNet_RPC
	client := rpc.New(endpoint)

	return &Solana{
		client: client,
	}
}
