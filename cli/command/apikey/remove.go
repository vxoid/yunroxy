package apikey

import (
	"log"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/vxoid/yunroxy/config"
	"github.com/vxoid/yunroxy/db"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Yunroxy API Key",

	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.GetConfig()
		if err != nil {
			color.Red("Could not open config file: '%s'", err)
			return
		}

		database, err := db.NewApiDb(config.Db)
		if err != nil {
			color.Red("Could not open the database: '%s'", err)
			return
		}

		apiKey, err := db.ParseApiKey(args[0])
		if err != nil {
			color.Red("Could not parse the API Key: '%s'", err)
			return
		}
		err = database.RemoveApiKey(apiKey)
		if err != nil {
			log.Fatal(err)
		}

		color.Green("Succesfully removed the API Key.")
	},
}
