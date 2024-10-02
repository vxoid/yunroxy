package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/vxoid/yunroxy/cli/command/apikey"
)

var rootCmd = &cobra.Command{
	Short: "Yunroxy CLI tool",
	Long:  `Yunroxy CLI tool for easy setup and management`,

	Run: nil,
}

func main() {
	apikey.ApiKeyCmd.AddCommand(apikey.NewCmd)
	apikey.ApiKeyCmd.AddCommand(apikey.RemoveCmd)
	rootCmd.AddCommand(apikey.ApiKeyCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
