package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(probeCmd)
}

func runProbe(cmd *cobra.Command, args []string) error {
	fmt.Println("hello Probe")
	return nil
}

var probeCmd = &cobra.Command{
	Use:   "probe",
	Short: "collect pentest environment information",
	RunE:  runProbe,
}

//var propeOption = new(model.ProbeOptions)
//propeOption := model.ProbeOptions{}
