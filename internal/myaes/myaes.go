package myaes

import (
	"sync"

	"github.com/yagoyudi/ufpr-criptografia-t1/internal/padding"
)

type AES struct {
}

func (aes *AES) Encrypt(initialKey []byte, plaintext []byte) ([]byte, error) {
	plaintext = padding.Pad(plaintext, blockSize)
	ciphertext := make([]byte, len(plaintext))

	numBlocks := len(plaintext) / blockSize

	roundKey := expandKey(initialKey)

	initSbox()

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

	for b := range blockChan {
		copy(ciphertext[b.index*blockSize:], b.block[:])
	}

	return ciphertext, nil
}

func getNextBlock(i int, plaintext []byte) [blockSize]byte {
	var start, end int
	var block [blockSize]byte

	start = i * blockSize
	end = start + blockSize
	copy(block[:], plaintext[start:end])

	return block
}

func (aes *AES) Decrypt(initialKey []byte, ciphertext []byte) ([]byte, error) {
	numBlocks := len(ciphertext) / blockSize
	plaintext := make([]byte, len(ciphertext))

	roundKey := expandKey(initialKey)

	initSbox()

	decryptedBlocks := make(chan struct {
		index int
		block Block
	}, numBlocks)

	var wg sync.WaitGroup

	for i := 0; i < numBlocks; i++ {
		wg.Add(1)
		go func(blockIndex int) {
			defer wg.Done()

			var b Block
			copy(b.state[:], ciphertext[blockIndex*blockSize:(blockIndex+1)*blockSize])

			b.decrypt(roundKey)

			decryptedBlocks <- struct {
				index int
				block Block
			}{
				index: blockIndex,
				block: b,
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close(decryptedBlocks)
	}()

	orderedBlocks := make([]Block, numBlocks)
	for decryptedBlock := range decryptedBlocks {
		orderedBlocks[decryptedBlock.index] = decryptedBlock.block
	}

	for i, b := range orderedBlocks {
		copy(plaintext[i*blockSize:(i+1)*blockSize], b.state[:])
	}

	plaintext, err := padding.Unpad(plaintext, blockSize)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
