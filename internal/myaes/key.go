package myaes

const (
	keySize        = 16
	numberOfRounds = 10
)

func expandKey(initialKey []byte) [][]byte {
	roundKeys := make([][]byte, numberOfRounds+1)
	for i := range roundKeys {
		roundKeys[i] = make([]byte, keySize)
	}

	copy(roundKeys[0], initialKey)
	for i := 1; i <= numberOfRounds; i++ {
		tmp := roundKeys[i-1][12:16]

		// RotWord, SubWord e XOR com rcon
		tmp = append(tmp[1:], tmp[0]) // RotWord

		// SubWord
		tmp[0] = sbox[tmp[0]]
		tmp[1] = sbox[tmp[1]]
		tmp[2] = sbox[tmp[2]]
		tmp[3] = sbox[tmp[3]]

		tmp[0] ^= rcon[i-1]

		for j := 0; j < keySize; j++ {
			if j < 4 {
				roundKeys[i][j] = roundKeys[i-1][j] ^ tmp[j]
			} else {
				roundKeys[i][j] = roundKeys[i-1][j] ^ roundKeys[i][j-4]
			}
		}
	}
	return roundKeys
}
