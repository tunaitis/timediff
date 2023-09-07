package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cmdFormats = &cobra.Command{
	Use:   "formats",
	Short: "Display available time input formats",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Available formats:")

		for _, f := range formats {
			fmt.Printf("  %s\n", f)
		}
	},
}

func init() {
	rootCmd.AddCommand(cmdFormats)
}
