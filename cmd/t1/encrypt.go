package cmd

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/criptografia-t1/internal/aes128"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [key]",
	Short: "Encrypt plaintext",
	Long:  "Encrypt plaintext",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key, err := decodeKey(args[0])
		if err != nil {
			log.Fatal(err)
		}

		aes := aes128.NewAES(key)
		plaintext, err := readPlaintextFromStdin()
		if err != nil {
			log.Fatal(err)
		}

		ciphertext, err := aes.Encrypt([]byte(plaintext))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(ciphertext))
	},
}
