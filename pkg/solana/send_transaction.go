package solana

import (
	"context"
	"os"

	solanago "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
)

type SendTransactionRequest struct {
	PrivateKeyFilePath string
	AddressTo          string
	AddressFrom        string
	Amount             uint64
}

type SendTransactionResponse struct {
	Signature string
}

func (s *Solana) SendTransaction(ctx context.Context, req *SendTransactionRequest) (*SendTransactionResponse, error) {
	b, err := os.ReadFile(req.PrivateKeyFilePath)
	if err != nil {
		return nil, err
	}

	privateKeyFrom := solanago.MustPrivateKeyFromBase58(string(b))
	publicKeyFrom := solanago.MustPublicKeyFromBase58(req.AddressFrom)
	publicKeyTo := solanago.MustPublicKeyFromBase58(req.AddressTo)

	recent, err := s.client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	tx, err := solanago.NewTransaction(
		[]solanago.Instruction{
			system.NewTransferInstruction(
				req.Amount,
				publicKeyFrom,
				publicKeyTo,
			).Build(),
		},
		recent.Value.Blockhash,
		solanago.TransactionPayer(publicKeyFrom),
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Sign(
		func(key solanago.PublicKey) *solanago.PrivateKey {
			if publicKeyFrom.Equals(key) {
				return &privateKeyFrom
			}
			return nil
		},
	)
	if err != nil {
		return nil, err

	}

	// spew.Dump(tx)
	// fmt.Println("----------------------------------------------")
	// // Pretty print the transaction:
	// tx.EncodeTree(text.NewTreeEncoder(os.Stdout, "Transfer SOL"))

	signature, err := s.client.SendTransaction(ctx, tx)
	if err != nil {
		return nil, err
	}

	return &SendTransactionResponse{
		Signature: signature.String(),
	}, nil
}
