package aes128

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecrypt(t *testing.T) {
	key := []byte{
		0x2b, 0x7e, 0x15, 0x16,
		0x28, 0xae, 0xd2, 0xa6,
		0xab, 0xf7, 0x97, 0x75,
		0x45, 0x21, 0x48, 0x8d,
	}
	aes := NewAES(key)
	want := []byte("This test will be using 2 blocks")
	ciphertext, err := aes.Encrypt(want)
	assert.NoError(t, err)
	got, err := aes.Decrypt(ciphertext)
	assert.NoError(t, err)
	assert.Equal(t, got, want)
}
