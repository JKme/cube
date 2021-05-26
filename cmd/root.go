package cmd

import (
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func parseGlobalOptions()(*model.GlobalOptions, error){
	globalOpts := model.NewGlobalOptions()
	threads, _ := rootCmd.Flags().GetInt("threads")

	if threads <= 0 {
		return nil, fmt.Errorf("threads must be bigger than 0")
	}
	return globalOpts, nil
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}