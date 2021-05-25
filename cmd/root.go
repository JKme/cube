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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func parseGlobalOptions() (*model.GlobalOptions, error) {
	globalopts := model.NewGlobalOptions()
	threads, _ := rootCmd.Flags().GetInt("threads")
	if threads <= 0 {
		return nil, fmt.Errorf("threads must be bigger than 0")
	}
	globalopts.Threads = threads

	timeout, _ := rootCmd.Flags().GetInt("timeout")

	return globalopts, nil
}

func init() {
	rootCmd.PersistentFlags().IntP("threads", "n", 20, "Number of concurrent threads")
	rootCmd.PersistentFlags().DurationP("timeout", "t", 5, "Timeout each thread waits")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output file to write results to (defaults to stdout)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output (errors)")
	rootCmd.PersistentFlags().StringP("pattern", "p", "", "File containing replacement patterns")
}
