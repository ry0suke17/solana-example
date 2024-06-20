package solana

import (
	"context"

	solanago "github.com/gagliardetto/solana-go"
)

func (s *Solana) RequestAirdrop(ctx context.Context, address string, amount uint64) (string, error) {
	lamports := amount * solanago.LAMPORTS_PER_SOL
	publickKey := solanago.PublicKeyFromBytes([]byte(address))
	signature, err := s.client.RequestAirdrop(ctx, publickKey, lamports, "")
	if err != nil {
		return "", err
	}
	return signature.String(), nil
}
