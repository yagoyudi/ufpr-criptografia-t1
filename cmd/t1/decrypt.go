package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/criptografia-t1/internal/aes128"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt [key]",
	Short: "Decrypt ciphertext",
	Long:  "Decrypt ciphertext",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key, err := decodeKey(args[0])
		if err != nil {
			log.Fatal(err)
		}

		aes := aes128.NewAES(key)
		ciphertext, err := readCiphertextFromStdin()
		if err != nil {
			log.Fatal(err)
		}

		plaintext, err := aes.Decrypt([]byte(ciphertext))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x\n", plaintext)
	},
}
