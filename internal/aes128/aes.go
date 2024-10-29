package aes128

type AES struct {
	expandedKey [][]byte
}

func NewAES(key []byte) *AES {
	return &AES{
		expandedKey: keyExpansion(key),
	}
}
