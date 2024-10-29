package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/yagoyudi/criptografia-t1/internal/aes128"
)

func main() {
	var keyString string
	flag.StringVar(&keyString, "key", "", "Hexadecimal key (16 bytes)")
	flag.Parse()

	key, err := hex.DecodeString(strings.ReplaceAll(keyString, " ", ""))
	if err != nil || len(key) != 16 {
		log.Fatal("Key must be 16 bytes in hexadecimal format.")
	}

	aes := aes128.NewAES(key)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter plaintext: ")
	plaintext, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	plaintext = strings.TrimSpace(plaintext)

	ciphertext, err := aes.Encrypt([]byte(plaintext))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Ciphertext: %x\n", ciphertext)

	decrypted, err := aes.Decrypt(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
