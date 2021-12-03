package cmd

import (
	"github.com/sirzerator/advent2021/days"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:       "run [day]",
	Short:     "Run a specific days challenge",
	Long:      ``,
	Args:      cobra.ExactValidArgs(1),
	ValidArgs: []string{"1"},
	Run: func(cmd *cobra.Command, args []string) {
		days.Run(args[0], Verbose)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
