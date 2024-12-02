package myaes

type AES struct {
	key *Key
}

func NewAES(initialKey [keySize]byte) *AES {
	var aes *AES
	var key *Key

	key = &Key{}
	key.expand(initialKey)

	aes = &AES{
		key: key,
	}

	return aes
}
