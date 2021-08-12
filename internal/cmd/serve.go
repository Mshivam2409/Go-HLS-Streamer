package cmd

import (
	"github.com/Mshivam2409/hls-streamer/internal/api"
	"github.com/Mshivam2409/hls-streamer/internal/db"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the SMTP Server",
	RunE: func(cmd *cobra.Command, args []string) error {

		db.InitializeCache()
		if err := api.HTTPListen(); err != nil {
			return err
		}
		return nil
	},
}
