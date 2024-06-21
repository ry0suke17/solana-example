package solana

import (
	"context"

	solanago "github.com/gagliardetto/solana-go"
)

type GetBalanceRequest struct {
	Address string
}

type GetBalanceResponse struct {
	Value uint64
}

func (s *Solana) GetBalance(ctx context.Context, req *GetBalanceRequest) (*GetBalanceResponse, error) {
	publickKey, err := solanago.PublicKeyFromBase58(req.Address)
	if err != nil {
		return nil, err
	}
	result, err := s.client.GetBalance(ctx, publickKey, "")
	if err != nil {
		return nil, err
	}
	return &GetBalanceResponse{
		Value: result.Value,
	}, nil
}
