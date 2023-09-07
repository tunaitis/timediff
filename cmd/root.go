package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var formats []string = []string{"1504", "15:04", "3:04PM"}

var rootCmd = &cobra.Command{
	Use:   "timediff TIME_START TIME_END",
	Short: "calculate the difference between two time expressions",
	Args:  cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {

		var start time.Time
		var end time.Time
		var err error

		for _, f := range formats {

			if start.IsZero() {
				start, err = time.Parse(f, strings.ToUpper(args[0]))
			}

			if end.IsZero() {
				end, err = time.Parse(f, strings.ToUpper(args[1]))
			}
		}

		if err != nil {
			return err
		}

		diff := end.Sub(start)

		fmt.Printf("Diff: %s\n", diff)
		fmt.Printf("Hours: %.2f\n", diff.Hours())
		fmt.Printf("Minutes: %.0f\n", diff.Minutes())

		return nil
	},
}

func Execute(version string) {
	rootCmd.Version = version

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
