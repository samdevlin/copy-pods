package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/dhowden/tag"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list available podcasts",
	Long:  "List all downloaded podcasts from apple podcasts",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		list_files()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list_files() {
	path, e := get_podcast_directory()
	if e != nil {
		panic(e)
	}

	path += "/Library/Group Containers/243LU875E5.groups.com.apple.podcasts/Library/Cache"
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !info.IsDir() && strings.Contains(path, ".mp3") {
			file, err := os.Open(path)

			if err != nil {
				fmt.Println(err)
			}

			m, err := tag.ReadFrom(file)

			if err == nil {
				fmt.Printf("Title: %s \n Podcast: %s\n\n", m.Title(), m.Album())
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func get_podcast_directory() (response string, err error) {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not establish user's home directory. Error: \n %s", err)
		return "", err
	}
	return home, nil
}
