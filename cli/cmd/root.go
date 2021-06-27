package cmd

import (
	"cube/log"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:          "cube",
	SilenceUsage: true,
}

// Execute is the main cobra method

func Execute() {

	if err := rootCmd.Execute(); err != nil {

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

	globalopts.Timeout, _ = rootCmd.Flags().GetInt("timeout")
	globalopts.Delay, _ = rootCmd.Flags().GetInt("delay")

	verbose, err := rootCmd.Flags().GetBool("verbose")
	if err != nil {
		return nil, fmt.Errorf("invalid value for verbose: %w", err)
	}
	if verbose {
		log.InitLog("DEBUG")
	} else {
		log.InitLog("INFO")
	}
	globalopts.Verbose = verbose

	return globalopts, nil
}

func init() {
	rootCmd.PersistentFlags().IntP("threads", "n", 30, "Number of concurrent threads")
	rootCmd.PersistentFlags().IntP("timeout", "", 5, "Timeout each thread waits")
	rootCmd.PersistentFlags().IntP("delay", "", 0, "delay for request")
	//rootCmd.PersistentFlags().StringP("output", "o", "", "Output file to write results to (defaults to stdout)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output (errors)")
}
