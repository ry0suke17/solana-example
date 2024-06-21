package solana

import (
	"context"

	solanago "github.com/gagliardetto/solana-go"
)

func (s *Solana) RequestAirdrop(ctx context.Context, address string, amount uint64) (string, error) {
	publickKey, err := solanago.PublicKeyFromBase58(address)
	if err != nil {
		return "", err
	}
	signature, err := s.client.RequestAirdrop(ctx, publickKey, amount, "")
	if err != nil {
		return "", err
	}
	return signature.String(), nil
}
