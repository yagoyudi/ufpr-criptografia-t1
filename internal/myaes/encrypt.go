package myaes

import (
	"sync"
)

func (aes *AES) Encrypt(initialKey []byte, plaintext []byte) ([]byte, error) {
	numBlocks := (len(plaintext) + blockSize - 1) / blockSize
	ciphertext := make([]byte, numBlocks*blockSize)

	roundKey := expandKey(initialKey)

	var wg sync.WaitGroup
	blockChan := make(chan struct {
		index int
		block [blockSize]byte
	}, numBlocks)

	for i := 0; i < numBlocks; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			block := getNextBlock(i, plaintext)
			b := Block{state: block}
			b.encrypt(roundKey)
			blockChan <- struct {
				index int
				block [blockSize]byte
			}{index: i, block: b.state}
		}(i)
	}

	go func() {
		wg.Wait()
		close(blockChan)
	}()

	for blk := range blockChan {
		copy(ciphertext[blk.index*blockSize:], blk.block[:])
	}

	return ciphertext, nil
}

func getNextBlock(i int, plaintext []byte) [blockSize]byte {
	var start, end int
	var block [blockSize]byte

	start = i * blockSize
	end = start + blockSize
	if end > len(plaintext) {
		copy(block[:], plaintext[start:])

		// Adiciona o tamanho do padding como padding.
		howMuchFilled := len(plaintext[start:])
		paddingLen := blockSize - howMuchFilled
		for j := howMuchFilled; j < blockSize; j++ {
			block[j] = byte(paddingLen)
		}
	} else {
		copy(block[:], plaintext[start:end])
	}

	return block
}
