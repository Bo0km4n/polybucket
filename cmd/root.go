package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "polyb",
	Short: "PolyBucket is a very simple CLI tool to manage the version of ML models",
	Long: `PolyBucket provide the comfortable management system of ML models.
This tool is independent from Git and more versioning tool.
You can use the only two commands: "push" or "pull"`,
}

func init() {
	rootCmd.AddCommand(pushCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
