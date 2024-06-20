package solana

import (
	"os"

	solanago "github.com/gagliardetto/solana-go"
)

func (s *Solana) GenKeyPair(privateKeyFilePath string) (string, error) {
	privateKey, err := solanago.NewRandomPrivateKey()
	if err != nil {
		return "", err
	}

	f, err := os.Create(privateKeyFilePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.WriteString(privateKey.String())
	if err != nil {
		return "", err
	}

	return privateKey.PublicKey().String(), nil
}
