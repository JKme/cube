package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runCrack(cmd *cobra.Command, args []string) error {
	fmt.Println("hello Probe")
	return nil
}

var versionCmd = &cobra.Command{
	Use:   "crack",
	Short: "Brute-force Attack",
	RunE:  runCrack,
}
