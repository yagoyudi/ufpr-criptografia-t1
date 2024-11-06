package cmd

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func decodeKey(key string) ([]byte, error) {
	content, err := os.ReadFile(key)
	if err != nil {
		return nil, err
	}
	decoded, err := base64.StdEncoding.DecodeString(string(content))
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

func readCiphertextFromStdin() ([]byte, error) {
	reader := bufio.NewReader(os.Stdin)
	ciphertext, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	ciphertext = strings.TrimSpace(ciphertext)
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
