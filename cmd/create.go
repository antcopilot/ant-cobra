package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)
// dummy commit
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new resource",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create command executed")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
