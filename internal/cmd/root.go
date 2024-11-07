package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "t1",
	Short: "Trabalho 1 de Criptografia",
	Long:  "Trocar a Caixa-S do AES por alguma outra cifra de substituição (exceto Cifra de Cesar)",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
