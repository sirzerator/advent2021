package cmd

import (
	"github.com/sirzerator/advent2021/days/day1"
	"github.com/sirzerator/advent2021/days/day2"
	"github.com/sirzerator/advent2021/days/day3"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:       "run [day]",
	Short:     "Run a specific days challenge",
	Long:      ``,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"1", "2", "3"},
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "1":
			day1.Run(Verbose)
		case "2":
			day2.Run(Verbose)
		case "3":
			day3.Run(Verbose)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
