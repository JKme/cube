package cmd

import (
	"context"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
)

var rootCmd = &cobra.Command{
	Use:          "cube",
	SilenceUsage: true,
}

var mainContext context.Context

// Execute is the main cobra method

func Execute() {
	var cancel context.CancelFunc
	mainContext, cancel = context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan:
			// caught CTRL+C
			fmt.Println("\n[!] Keyboard interrupt detected, terminating.")
			cancel()
		case <-mainContext.Done():
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		// Leaving this in results in the same error appearing twice
		// Once before and once after the help output. Not sure if
		// this is going to be needed to output other errors that
		// aren't automatically outputted.
		// fmt.Println(err)
		os.Exit(1)
	}
}

func configureGlobalOptions() {
	if err := rootCmd.MarkPersistentFlagRequired("plugin"); err != nil {
		log.Fatalf("error on marking flag as required: %v", err)
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
	if globalopts.Timeout < 3 {
		return nil, fmt.Errorf("timeout must be bigger than default vaule(default: %v)", model.TIMEOUT)
	}

	verbose, err := rootCmd.Flags().GetBool("verbose")
	if err != nil {
		return nil, fmt.Errorf("invalid value for verbose: %w", err)
	}
	globalopts.Verbose = verbose

	plugin, err := rootCmd.Flags().GetString("plugin")
	if err != nil {
		return nil, fmt.Errorf("invalid value for verbose: %w", err)
	}
	globalopts.Plugin = plugin

	return globalopts, nil
}

func init() {
	rootCmd.PersistentFlags().IntP("threads", "n", 20, "Number of concurrent threads")
	rootCmd.PersistentFlags().DurationP("timeout", "t", 5, "Timeout each thread waits")
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output file to write results to (defaults to stdout)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output (errors)")
	rootCmd.PersistentFlags().BoolP("plugin", "x", false, "plugin to scan")
}
