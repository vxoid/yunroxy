package apikey

import "github.com/spf13/cobra"

var ApiKeyCmd = &cobra.Command{
	Use:   "api-key",
	Short: "Yunroxy API Key",
	Long:  `Yunroxy API Key Management commands`,

	Run: nil,
}
