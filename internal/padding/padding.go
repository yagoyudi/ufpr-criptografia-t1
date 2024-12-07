package padding

import (
	"bytes"
	"fmt"
)

func Pad(data []byte, blockSize int) []byte {
	n := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(n)}, n)
	return append(data, padding...)
}

func Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("unpad: data is empty")
	}
	n := int(data[len(data)-1])
	if n <= 0 || n > blockSize || len(data) < n {
		return nil, fmt.Errorf("unpad: invalid padding")
	}
	return data[:len(data)-n], nil
}
