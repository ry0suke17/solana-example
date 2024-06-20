package solana

import (
	"os"

	solanago "github.com/gagliardetto/solana-go"
)

func (s *Solana) GetPublicKey(privateKeyFilePath string) (string, error) {
	b, err := os.ReadFile(privateKeyFilePath)
	if err != nil {
		return "", err
	}

	privateKey, err := solanago.PrivateKeyFromBase58(string(b))
	if err != nil {
		return "", err
	}

	return privateKey.PublicKey().String(), nil
}
