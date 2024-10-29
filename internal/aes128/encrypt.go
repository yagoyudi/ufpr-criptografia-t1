package aes128

func subBytes(state []byte) {
	for i := 0; i < len(state); i++ {
		state[i] = sbox[state[i]]
	}
}

// ShiftRows faz o deslocamento das linhas.
// Isso:
// s0  s4  s8  s12
// s1  s5  s9  s13
// s2  s6  s10 s14
// s3  s7  s11 s15
// Vira:
// s0  s4  s8  s12
// s5  s9  s13 s1
// s10 s14 s2  s6
// s15 s3  s7  s11
func shiftRows(state []byte) {
	state[1], state[5], state[9], state[13] = state[5], state[9], state[13], state[1]
	state[2], state[6], state[10], state[14] = state[10], state[14], state[2], state[6]
	state[3], state[7], state[11], state[15] = state[15], state[3], state[7], state[11]
}

// MixColumns mistura as colunas
func mixColumns(state []byte) {
	for c := 0; c < 4; c++ {
		var tmp [4]byte
		tmp[0] = state[c*4+0]
		tmp[1] = state[c*4+1]
		tmp[2] = state[c*4+2]
		tmp[3] = state[c*4+3]
		// Multiplicações no campo GF(2^8)
		state[c*4+0] = xtime(tmp[0]) ^ (xtime(tmp[1]) ^ tmp[1]) ^ tmp[2] ^ tmp[3]
		state[c*4+1] = tmp[0] ^ xtime(tmp[1]) ^ (xtime(tmp[2]) ^ tmp[2]) ^ tmp[3]
		state[c*4+2] = tmp[0] ^ tmp[1] ^ xtime(tmp[2]) ^ (xtime(tmp[3]) ^ tmp[3])
		state[c*4+3] = (xtime(tmp[0]) ^ tmp[0]) ^ tmp[1] ^ tmp[2] ^ xtime(tmp[3])
	}
}

func (aes *AES) EncryptBlock(input []byte) []byte {
	state := make([]byte, 16)
	copy(state, input)

	addRoundKey(state, aes.expandedKey[0])

	for i := 1; i < 10; i++ {
		subBytes(state)
		shiftRows(state)
		mixColumns(state)
		addRoundKey(state, aes.expandedKey[i])
	}

	subBytes(state)
	shiftRows(state)
	addRoundKey(state, aes.expandedKey[10])

	return state
}
