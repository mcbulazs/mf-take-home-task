package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "CLI tool for products",
	Long:  "A cli tool to manage product stocks",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(increaseCmd)
	rootCmd.AddCommand(decreaseCmd)
	rootCmd.AddCommand(reportCmd)
}
