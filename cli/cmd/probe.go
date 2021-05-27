package cmd

import (
	"cube/cli"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
)

var probeCmd *cobra.Command

func runProbe(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseProbeOptions()
	cli.StartProbeTask(opt, globalopts)
}

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
	propePlugin.ScanPlugin, err = probeCmd.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	propePlugin.Target, err = probeCmd.Flags().GetString("target-ip")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-ip: %w", err)
	}
	propePlugin.TargetFile, err = probeCmd.Flags().GetString("target-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}
	return globalOpts, propePlugin, nil
}

func init() {
	var probeCmd = &cobra.Command{
		Use:   "probe",
		Short: "collect pentest environment information",
		Run:   runProbe,
	}

	probeCmd.Flags().StringP("port", "p", "", "target port")
	probeCmd.Flags().StringP("plugin", "x", "", "plugin to scan(e.g. OXID)")
	probeCmd.Flags().StringP("all-plugin", "X", "", "all plugin to scan(e.g. OXID)")
	probeCmd.Flags().StringP("target-ip", "i", "", "ip range to probe for(e.g. 192.168.1.1/24)")
	probeCmd.Flags().StringP("target-file", "I", "", "File to probe for(e.g. ip.txt)")
	probeCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		configureGlobalOptions()
	}
	rootCmd.AddCommand(probeCmd)
}
