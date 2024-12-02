package playfair

import "strings"

func genTable(keyword string) [5][5]byte {
	var table [5][5]byte
	alphabet := "ABCDEFGHIKLMNOPQRSTUVWXYZ" // J é substituído por I
	used := make(map[byte]bool)

	keyword = strings.ToUpper(keyword)
	keyword = strings.ReplaceAll(keyword, "J", "I")

	index := 0
	for i := 0; i < len(keyword); i++ {
		if !used[keyword[i]] && strings.Contains(alphabet, string(keyword[i])) {
			used[keyword[i]] = true
			table[index/5][index%5] = keyword[i]
			index++
		}
	}

	for i := 0; i < len(alphabet); i++ {
		if !used[alphabet[i]] {
			table[index/5][index%5] = alphabet[i]
			index++
		}
	}

	return table
}

// Substitute substitui com base na tabela Playfair.
func substitute(input byte, table [5][5]byte) byte {
	input = byte(strings.ToUpper(string(input))[0])
	if input == 'J' {
		input = 'I'
	}

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if table[row][col] == input {
				// Retorne o próximo elemento na tabela
				return table[(row+1)%5][col]
			}
		}
	}
	return input // Se não encontrar, retorne o próprio caractere
}

// SubBytes substitui os bytes com Playfair.
func SubBytes(state [4][4]byte, table [5][5]byte) [4][4]byte {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			state[i][j] = substitute(state[i][j], table)
		}
	}
	return state
}

// invSubstitute reverte a substituição com base na tabela Playfair.
func invSubstitute(input byte, table [5][5]byte) byte {
	input = byte(strings.ToUpper(string(input))[0])
	if input == 'J' {
		input = 'I'
	}

	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if table[row][col] == input {
				// Retorne o elemento anterior na tabela
				return table[(row+4)%5][col]
			}
		}
	}
	return input // Se não encontrar, retorne o próprio caractere
}

// InvSubBytes reverte a substituição dos bytes com Playfair.
func InvSubBytes(state [4][4]byte, table [5][5]byte) [4][4]byte {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			state[i][j] = invSubstitute(state[i][j], table)
		}
	}
	return state
}
