package cmd

import (
	"cube/core"
	"cube/gologger"
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

func parseGlobalOptions() (*core.GlobalOption, error) {
	globalopts := core.NewGlobalOptions()
	threads, _ := rootCmd.Flags().GetInt("threads")
	if threads <= 0 {
		return nil, fmt.Errorf("threads must be greater than 0")
	}
	globalopts.Threads = threads

	globalopts.Timeout, _ = rootCmd.Flags().GetInt("timeout")
	globalopts.Delay, _ = rootCmd.Flags().GetFloat64("delay")
	globalopts.Output, _ = rootCmd.Flags().GetString("output")

	verbose, err := rootCmd.Flags().GetBool("verbose")
	if err != nil {
		return nil, fmt.Errorf("invalid value for verbose: %w", err)
	}
	if verbose {
		gologger.InitLog("DEBUG")
	} else {
		gologger.InitLog("INFO")
	}
	globalopts.Verbose = verbose

	return globalopts, nil
}

func init() {
	rootCmd.PersistentFlags().IntP("threads", "n", 30, "Number of concurrent requests")
	rootCmd.PersistentFlags().IntP("timeout", "", 5, "Seconds to wait before timeout connection")
	rootCmd.PersistentFlags().Float64P("delay", "d", 0, "Delay in random seconds between each TCP/UDP request")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output file to write results to (eg. pwn.xlsx)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose (Default error)")
}
