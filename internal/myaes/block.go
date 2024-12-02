package myaes

const (
	blockSize = 16
)

type Block struct {
	state [blockSize]byte
}

func (b *Block) encrypt(key *Key) {
	b.addRoundKey(key.round[0])

	for i := 1; i < 10; i++ {
		b.subBytes()
		b.shiftRows()
		b.mixColumns()
		b.addRoundKey(key.round[i])
	}

	b.subBytes()
	b.shiftRows()
	b.addRoundKey(key.round[10])
}

func (b *Block) decrypt(key *Key) {
	b.addRoundKey(key.round[10])

	for i := 9; i > 0; i-- {
		b.invShiftRows()
		b.invSubBytes()
		b.addRoundKey(key.round[i])
		b.invMixColumns()
	}

	b.invShiftRows()
	b.invSubBytes()
	b.addRoundKey(key.round[0])
}

func (b *Block) subBytes() {
	for i := 0; i < len(b.state); i++ {
		b.state[i] = sbox[b.state[i]]
	}
}

func (b *Block) shiftRows() {
	b.state[1], b.state[5], b.state[9], b.state[13] = b.state[5], b.state[9], b.state[13], b.state[1]
	b.state[2], b.state[6], b.state[10], b.state[14] = b.state[10], b.state[14], b.state[2], b.state[6]
	b.state[3], b.state[7], b.state[11], b.state[15] = b.state[15], b.state[3], b.state[7], b.state[11]
}

func (b *Block) mixColumns() {
	for c := 0; c < 4; c++ {
		var tmp [4]byte
		tmp[0] = b.state[c*4+0]
		tmp[1] = b.state[c*4+1]
		tmp[2] = b.state[c*4+2]
		tmp[3] = b.state[c*4+3]
		// Multiplicações no campo GF(2^8)
		b.state[c*4+0] = xtime(tmp[0]) ^ (xtime(tmp[1]) ^ tmp[1]) ^ tmp[2] ^ tmp[3]
		b.state[c*4+1] = tmp[0] ^ xtime(tmp[1]) ^ (xtime(tmp[2]) ^ tmp[2]) ^ tmp[3]
		b.state[c*4+2] = tmp[0] ^ tmp[1] ^ xtime(tmp[2]) ^ (xtime(tmp[3]) ^ tmp[3])
		b.state[c*4+3] = (xtime(tmp[0]) ^ tmp[0]) ^ tmp[1] ^ tmp[2] ^ xtime(tmp[3])
	}
}

func (b *Block) addRoundKey(key [keySize]byte) {
	for i := 0; i < len(b.state); i++ {
		b.state[i] ^= key[i]
	}
}

func (b *Block) invSubBytes() {
	for i := 0; i < len(b.state); i++ {
		b.state[i] = sboxInv[b.state[i]]
	}
}

func (b *Block) invShiftRows() {
	b.state[1], b.state[5], b.state[9], b.state[13] = b.state[13], b.state[1], b.state[5], b.state[9]
	b.state[2], b.state[6], b.state[10], b.state[14] = b.state[10], b.state[14], b.state[2], b.state[6]
	b.state[3], b.state[7], b.state[11], b.state[15] = b.state[7], b.state[11], b.state[15], b.state[3]
}

func (b *Block) invMixColumns() {
	for c := 0; c < 4; c++ {
		var tmp [4]byte
		tmp[0] = b.state[c*4+0]
		tmp[1] = b.state[c*4+1]
		tmp[2] = b.state[c*4+2]
		tmp[3] = b.state[c*4+3]

		// Multiplicações no campo GF(2^8) usando os coeficientes da matriz inversa
		b.state[c*4+0] = mul0x0e(tmp[0]) ^ mul0x0b(tmp[1]) ^ mul0x0d(tmp[2]) ^ mul0x09(tmp[3])
		b.state[c*4+1] = mul0x09(tmp[0]) ^ mul0x0e(tmp[1]) ^ mul0x0b(tmp[2]) ^ mul0x0d(tmp[3])
		b.state[c*4+2] = mul0x0d(tmp[0]) ^ mul0x09(tmp[1]) ^ mul0x0e(tmp[2]) ^ mul0x0b(tmp[3])
		b.state[c*4+3] = mul0x0b(tmp[0]) ^ mul0x0d(tmp[1]) ^ mul0x09(tmp[2]) ^ mul0x0e(tmp[3])
	}
}
