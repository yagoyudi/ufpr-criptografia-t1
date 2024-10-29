package aes128

func mul0x0e(x byte) byte {
	return xtime(xtime(xtime(x)^x) ^ x)
}

func mul0x0b(x byte) byte {
	return xtime(xtime(xtime(x))^x) ^ x
}

func mul0x0d(x byte) byte {
	return xtime(xtime(xtime(x)^x)) ^ x
}

func mul0x09(x byte) byte {
	return xtime(xtime(xtime(x))) ^ x
}

// Xtime realiza a multiplicação por 2 em GF(2^8)
func xtime(b byte) byte {
	if b&0x80 != 0 {
		return (b << 1) ^ 0x1B
	}
	return b << 1
}

// Função para rotacionar uma palavra (4 bytes)
func rotWord(word []byte) {
	word[0], word[1], word[2], word[3] = word[1], word[2], word[3], word[0]
}
