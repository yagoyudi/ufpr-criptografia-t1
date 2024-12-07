package stdlib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"

	"github.com/yagoyudi/ufpr-criptografia-t1/internal/padding"
)

type AES struct {
}

func (a *AES) Encrypt(key, plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	iv := make([]byte, aes.BlockSize)
	_, err = rand.Read(iv)
	if err != nil {
		return nil, fmt.Errorf("failed to generate IV: %w", err)
	}

	cbc := cipher.NewCBCEncrypter(block, iv)
	plaintext = padding.Pad(plaintext, aes.BlockSize)
	ciphertext := make([]byte, len(plaintext)+aes.BlockSize)

	// Inclui o IV no início do ciphertext
	copy(ciphertext[:aes.BlockSize], iv)
	cbc.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func (a *AES) Decrypt(key, ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:] // Remove o IV

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}

	plaintext := make([]byte, len(ciphertext))
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(plaintext, ciphertext)

	plaintext, err = padding.Unpad(plaintext, aes.BlockSize)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
