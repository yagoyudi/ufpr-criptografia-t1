package myaes

import "sync"

func (aes *AES) Decrypt(ciphertext []byte) ([]byte, error) {
	numBlocks := len(ciphertext) / blockSize
	plaintext := make([]byte, len(ciphertext))

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

			b.decrypt(aes.key)

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

	// Remove o padding
	paddingLen := int(plaintext[len(plaintext)-1])
	if paddingLen > 0 && paddingLen <= blockSize {
		plaintext = plaintext[:len(plaintext)-paddingLen]
	}

	return plaintext, nil
}
