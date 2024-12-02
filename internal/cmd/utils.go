package cmd

import (
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

func readCiphertext(path string) ([]byte, error) {
	ciphertext, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ciphertextStr := strings.TrimSpace(string(ciphertext))
	decoded, err := base64.StdEncoding.DecodeString(ciphertextStr)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}
