package cmd

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/yagoyudi/criptografia-t1/internal/myaes"
	"github.com/yagoyudi/criptografia-t1/internal/stdlib"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
	encryptCmd.Flags().BoolP("stdlib", "s", false, "Use Go's standard library AES instead of custom implementation")
}

var encryptCmd = &cobra.Command{
	Use:   "enc [key file path] [plaintext file path]",
	Short: "Encrypt plaintext",
	Long:  "Encrypt plaintext",
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

		plaintext, err := os.ReadFile(args[1])
		if err != nil {
			log.Fatal(err)
		}

		aes := myaes.NewAES(initialKey)

		var ciphertext []byte
		if useStdlib {
			ciphertext, err = stdlib.EncryptMessage(key, plaintext)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			ciphertext, err = aes.Encrypt(plaintext)
			if err != nil {
				log.Fatal(err)
			}
		}

		fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(ciphertext))
	},
}
