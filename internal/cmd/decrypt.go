package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/criptografia-t1/internal/myaes"
	"github.com/yagoyudi/criptografia-t1/internal/stdlib"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
	decryptCmd.Flags().BoolP("stdlib", "s", false, "Use Go's standard library AES instead of custom implementation")
}

var decryptCmd = &cobra.Command{
	Use:   "dec [key file path] [ciphertext file path]",
	Short: "Decrypt ciphertext",
	Long:  "Decrypt ciphertext",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		useStdlib, err := cmd.Flags().GetBool("stdlib")
		if err != nil {
			log.Fatal(err)
		}

		key, err := decodeKey(args[0])
		if err != nil {
			log.Fatal(err)
		}

		var initialKey [16]byte
		copy(initialKey[:], key)

		ciphertext, err := readCiphertext(args[1])
		if err != nil {
			log.Fatal(err)
		}

		aes := myaes.NewAES(initialKey)

		var plaintext []byte
		if useStdlib {
			plaintext, err = stdlib.DecryptMessage(key, ciphertext)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			plaintext, err = aes.Decrypt(ciphertext)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println(string(plaintext))
	},
}
