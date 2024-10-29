package main

import (
	"reflect"
	"testing"

	"github.com/yagoyudi/criptografia-t1/internal/aes128"
)

func TestEncryptionAndDecryption(t *testing.T) {
	key := []byte{
		0x2b, 0x7e, 0x15, 0x16,
		0x28, 0xae, 0xd2, 0xa6,
		0xab, 0xf7, 0x97, 0x75,
		0x45, 0x21, 0x48, 0x8d,
	}
	aes := aes128.NewAES(key)

	plaintext := []byte("This test will be using 2 blocks")
	ciphertext, err := aes.Encrypt(plaintext)
	if err != nil {
		t.Error(err)
	}
	decrypted, err := aes.Decrypt(ciphertext)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(decrypted, plaintext) {
		t.Errorf("decrypted text not equal plaintext")
	}
}
