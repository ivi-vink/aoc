package main

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var templateStr string = `
`

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new boilerplate entrypoint for a day of Advent of Code",
	RunE: func(cmd *cobra.Command, args []string) error {
		year, day, packageName := "", "", ""
		if len(args) == 2 {
			parts := strings.SplitN(args[0], "/", 2)
			if len(parts) == 2 {
				year = parts[0]
				day = parts[1]
			}
		} else {
			return fmt.Errorf("Expected two argument in format %q %q", "<year>/<day>", "<packageName>")
		}
		packageName = fmt.Sprintf("%s", args[1])

		if year == "" || day == "" {
			return fmt.Errorf("Expected one argument in format %q", "<year>/<day>")
		}

		_, err := strconv.Atoi(year)
		if err != nil {
			return fmt.Errorf("Couldn't parse year integer: %q", year)
		}

		dayInt, err := strconv.Atoi(day)
		if err != nil {
			return fmt.Errorf("Couldn't parse day integer: %q", day)
		}

		if dayInt < 1 || dayInt > 25 {
			return fmt.Errorf("Day must be between 1 and 25, got %d", dayInt)
		}

		dayStr := fmt.Sprintf("%02d", dayInt)

		fmt.Println(dayStr, packageName)
		if err := ensureDir(fmt.Sprintf("%s/%s", year, dayStr)); err != nil {
			return fmt.Errorf("Couldn't create directory %s/%s: %v", year, dayStr, err)
		}
		return nil
	},
}

func ensureDir(relativePath string) error {
	_, err := os.Stat(relativePath)
	if os.IsNotExist(err) {
		return os.MkdirAll(relativePath, fs.ModePerm)
	}
	return nil
}
