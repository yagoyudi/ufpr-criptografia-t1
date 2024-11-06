package cmd

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func decodeKey(key string) ([]byte, error) {
	decoded, err := hex.DecodeString(strings.ReplaceAll(key, " ", ""))
	if err != nil {
		return nil, err
	}
	if len(decoded) != 16 {
		return nil, fmt.Errorf("Key must be 16 bytes in hexadecimal format.")
	}
	return decoded, nil
}

func readPlaintextFromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	plaintext = strings.TrimSpace(plaintext)
	return plaintext, nil
}

func readCiphertextFromStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	ciphertext, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	ciphertext = strings.TrimSpace(ciphertext)
	return ciphertext, nil
}
