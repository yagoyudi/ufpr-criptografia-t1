package aes128

// invSubBytes substitui os bytes do estado pela Inverse S-Box
func invSubBytes(state []byte) {
	for i := 0; i < len(state); i++ {
		state[i] = sboxInv[state[i]]
	}
}

// invShiftRows inverte o deslocamento das linhas
// Isso:
// s0  s4  s8  s12
// s1  s5  s9  s13
// s2  s6  s10 s14
// s3  s7  s11 s15
// Vira:
// s0  s4  s8  s12
// s13 s1  s5  s9
// s10 s14 s2  s6
// s7  s11 s15 s3
func invShiftRows(state []byte) {
	state[1], state[5], state[9], state[13] = state[13], state[1], state[5], state[9]
	state[2], state[6], state[10], state[14] = state[10], state[14], state[2], state[6]
	state[3], state[7], state[11], state[15] = state[7], state[11], state[15], state[3]
}

// invMixColumns inverte a mistura das colunas
func invMixColumns(state []byte) {
	for c := 0; c < 4; c++ {
		var tmp [4]byte
		tmp[0] = state[c*4+0]
		tmp[1] = state[c*4+1]
		tmp[2] = state[c*4+2]
		tmp[3] = state[c*4+3]

		// Multiplicações no campo GF(2^8) usando os coeficientes da matriz inversa
		state[c*4+0] = mul0x0e(tmp[0]) ^ mul0x0b(tmp[1]) ^ mul0x0d(tmp[2]) ^ mul0x09(tmp[3])
		state[c*4+1] = mul0x09(tmp[0]) ^ mul0x0e(tmp[1]) ^ mul0x0b(tmp[2]) ^ mul0x0d(tmp[3])
		state[c*4+2] = mul0x0d(tmp[0]) ^ mul0x09(tmp[1]) ^ mul0x0e(tmp[2]) ^ mul0x0b(tmp[3])
		state[c*4+3] = mul0x0b(tmp[0]) ^ mul0x0d(tmp[1]) ^ mul0x09(tmp[2]) ^ mul0x0e(tmp[3])
	}
}

// decryptBlock decifra um bloco de 16 bytes
func (aes *AES) DecryptBlock(input []byte) []byte {
	state := make([]byte, 16)
	copy(state, input)

	addRoundKey(state, aes.expandedKey[10])

	for i := 9; i > 0; i-- {
		invShiftRows(state)
		invSubBytes(state)
		addRoundKey(state, aes.expandedKey[i])
		invMixColumns(state)
	}

	invShiftRows(state)
	invSubBytes(state)
	addRoundKey(state, aes.expandedKey[0])

	return state
}
