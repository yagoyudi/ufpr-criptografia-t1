package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Tabelas de substituição para cifragem e decifragem
var sBox [256]byte
var invSBox [256]byte

// Gera uma tabela de substituição aleatória e sua inversa
func generateSBox() {
	// Preenche sBox com uma permutação aleatória dos valores 0-255
	available := make([]byte, 256)
	for i := 0; i < 256; i++ {
		available[i] = byte(i)
	}

	for i := 0; i < 256; i++ {
		// Seleciona um índice aleatório
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(available))))
		sBox[i] = available[idx.Int64()]
		available = append(available[:idx.Int64()], available[idx.Int64()+1:]...)
	}

	// Cria a tabela inversa (invSBox) para decifragem
	for i, val := range sBox {
		invSBox[val] = byte(i)
	}
}

// Cifra um byte usando a sBox
func substituteByte(input byte) byte {
	return sBox[input]
}

// Decifra um byte usando a invSBox
func inverseSubstituteByte(input byte) byte {
	return invSBox[input]
}

func main() {
	// Gera a sBox e invSBox
	generateSBox()

	// Exemplo de cifragem e decifragem de um byte
	originalByte := byte(0x53)
	encryptedByte := substituteByte(originalByte)
	decryptedByte := inverseSubstituteByte(encryptedByte)

	fmt.Printf("Original: 0x%X\n", originalByte)
	fmt.Printf("Cifrado: 0x%X\n", encryptedByte)
	fmt.Printf("Decifrado: 0x%X\n", decryptedByte)

	// Verifica se o byte decifrado é igual ao original
	if decryptedByte == originalByte {
		fmt.Println("Decifragem bem-sucedida!")
	} else {
		fmt.Println("Erro na decifragem.")
	}
}
