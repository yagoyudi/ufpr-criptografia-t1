package myaes

const (
	keySize        = 16
	numberOfRounds = 10
)

type Key struct {
	// numberOfRounds + 1, para que round[0] seja a chave inicial.
	round [numberOfRounds + 1][keySize]byte
}

func (k *Key) expand(initialKey [keySize]byte) {
	copy(k.round[0][:], initialKey[:])
	for i := 1; i <= numberOfRounds; i++ {
		// últimos 4 bytes da rodada anterior
		tmp := k.round[i-1][12:16]

		// RotWord, SubWord e XOR com rcon
		tmp = append(tmp[1:], tmp[0]) // RotWord

		// SubWord
		tmp[0] = sbox[tmp[0]]
		tmp[1] = sbox[tmp[1]]
		tmp[2] = sbox[tmp[2]]
		tmp[3] = sbox[tmp[3]]

		tmp[0] ^= rcon[i-1]

		// Gera os 16 bytes da chave de rodada
		for j := 0; j < 16; j++ {
			if j < 4 {
				k.round[i][j] = k.round[i-1][j] ^ tmp[j]
			} else {
				k.round[i][j] = k.round[i-1][j] ^ k.round[i][j-4]
			}
		}
	}
}
