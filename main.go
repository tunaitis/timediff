package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var formats []string = []string{"1504", "15:04", "3:04PM"}

func usage(err error) {
	fmt.Println("timediff - calculate the difference between two time expressions")
	fmt.Println("")

	fmt.Println("Usage:")
	fmt.Println("  timediff TIME_START TIME_END")
	fmt.Println("")

	fmt.Println("Time format:")
	for _, f := range formats {
		fmt.Printf("  %s\n", f)
	}
	fmt.Println("")

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	args := os.Args

	if len(args) != 3 {
		usage(nil)
		return
	}

	var start time.Time
	var end time.Time
	var err error

	for _, f := range formats {

		if start.IsZero() {
			start, err = time.Parse(f, strings.ToUpper(args[1]))
		}

		if end.IsZero() {
			end, err = time.Parse(f, strings.ToUpper(args[2]))
		}
	}

	if err != nil {
		usage(err)
		return
	}

	diff := end.Sub(start)

	fmt.Printf("Diff: %s\n", diff)
	fmt.Printf("Hours: %.2f\n", diff.Hours())
	fmt.Printf("Minutes: %.0f\n", diff.Minutes())
}
