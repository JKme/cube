package cmd

import (
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
)

func runProbe(cmd *cobra.Command, args []string) error {
	fmt.Println("hello Probe")
	return nil
}

var probeCmd *cobra.Command

//var propeOption = new(model.ProbeOptions)
//propeOption := model.ProbeOptions{}

func parseProbeOptions() (*model.GlobalOptions, *model.ProbeOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}
	propePlugin := model.NewProbeOptions()

	propePlugin.Port, err = probeCmd.Flags().GetInt("port")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for scan port: %v", err)
	}
	propePlugin.Port, err = probeCmd.Flags().GetInt("func")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for func: %v", err)
	}
	propePlugin.Port, err = probeCmd.Flags().GetInt("target-ip")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-ip: %w", err)
	}
	propePlugin.Port, err = probeCmd.Flags().GetInt("target-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}
	return globalOpts, propePlugin, nil
}

func init() {
	var probeCmd = &cobra.Command{
		Use:   "probe",
		Short: "collect pentest environment information",
		RunE:  runProbe,
	}

	probeCmd.Flags().StringP("port", "p", "", "target port")
	probeCmd.Flags().StringP("func", "f", "", "func to scan")
	probeCmd.Flags().StringP("target-ip", "i", "192.168.1.1/24", "ip range to probe for)")
	probeCmd.Flags().StringP("target-file", "I", "ip.txt", "File to probe for")

	rootCmd.AddCommand(probeCmd)
}
