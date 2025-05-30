package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available podcasts",
	Long:  "List all downloaded podcasts from apple podcasts",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		get_files()
	},
}

// TODO: play with https://github.com/dhowden/tag
func init() {
	rootCmd.AddCommand(listCmd)
}

func get_files() {
	err := filepath.Walk("/Users/samdevlin/Library/Group Containers/243LU875E5.groups.com.apple.podcasts/Library/Cache", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}
