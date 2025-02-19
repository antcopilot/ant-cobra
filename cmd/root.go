package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ant-cobra",
	Short: "Ant-Cobra is a CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Ant-Cobra!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
