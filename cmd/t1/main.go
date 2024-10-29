package main

import (
	"fmt"

	"github.com/yagoyudi/criptografia-t1/internal/aes128"
)

func main() {
	key := []byte{
		0x2b, 0x7e, 0x15, 0x16,
		0x28, 0xae, 0xd2, 0xa6,
		0xab, 0xf7, 0x97, 0x75,
		0x45, 0x21, 0x48, 0x8d,
	}
	aes := aes128.NewAES(key)
	plaintext := []byte("This is a test")
	ciphertext := aes.EncryptBlock(plaintext)
	fmt.Printf("Ciphertext: %x\n", ciphertext)
	decrypted := aes.DecryptBlock(ciphertext)
	fmt.Printf("Decrypted: %s\n", decrypted)
}
