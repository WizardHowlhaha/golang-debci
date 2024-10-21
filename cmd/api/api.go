package api

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ApiCmd = &cobra.Command{
	Use:       "api [COMMAND] [ARGS]",
	ValidArgs: []string{"auth", "setkey"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debci api called")
	},
}

func init() {
	ApiCmd.AddCommand(setkeyCmd)
	ApiCmd.AddCommand(authCmd)
	ApiCmd.AddCommand(resetCmd)
	ApiCmd.AddCommand(deleteCmd)
}
