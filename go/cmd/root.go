package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "podsync",
	Short: "podsync is a cli tool for syncing your apple podcasts",
	Long:  "podsync is a cli tool for organising & syncing your downloaded apple podcasts to non iPod MP3 players",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing podsync '%s'\n", err)
		os.Exit(1)
	}
}
