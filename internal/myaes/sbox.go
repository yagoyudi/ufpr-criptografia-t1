package myaes

import (
	"math/rand"
)

var (
	sbox    [256]byte
	sboxInv [256]byte
)

func initSbox() {
	rand.Seed(0)

	used := make(map[byte]bool)
	for i := 0; i < 256; i++ {
		var num byte
		for {
			num = byte(rand.Intn(256))
			if !used[num] {
				used[num] = true
				break
			}
		}
		sbox[i] = num
		sboxInv[num] = byte(i)
	}
}
