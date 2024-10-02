package apikey

import (
	"encoding/hex"
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/db"
)

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new Yunroxy API Key",

	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}

		db, err := db.NewApiDb(config.Db)
		if err != nil {
			log.Fatal(err)
		}

		apiKey, err := db.CreateApiKey()
		if err != nil {
			log.Fatal(err)
		}

		color.Green("Created New API KEY: %s", hex.EncodeToString(apiKey))
	},
}
