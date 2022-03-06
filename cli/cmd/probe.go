package cmd

import (
	"cube/core"
	"cube/core/probemodule"
	"cube/gologger"
	"fmt"
	"github.com/spf13/cobra"
)

var probeCli *cobra.Command

func runProbe(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseProbeOptions()

	probemodule.StartProbe(opt, globalopts)
}

func parseProbeOptions() (*core.GlobalOption, *probemodule.ProbeOption, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}
	probeOption := probemodule.NewProbeOption()

	probeOption.PluginName, err = probeCli.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	probeOption.Port, err = probeCli.Flags().GetString("port")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for scan port: %v", err)
	}

	probeOption.Ip, err = probeCli.Flags().GetString("service")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-ip: %w", err)
	}
	probeOption.IpFile, err = probeCli.Flags().GetString("service-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}
	return globalOpts, probeOption, nil
}

func init() {
	probeCli = &cobra.Command{
		Use:   "probe",
		Long:  probemodule.ProbeHelpTable(),
		Short: "probe pentest env",
		Run:   runProbe,
		Example: `cube probe -s 192.168.1.1 -x oxid
cube probe -s 192.168.1.1 -x oxid,zookeeper,ms17010
cube probe -s 192.168.1.1/24 -x X
		`,
	}

	probeCli.Flags().StringP("port", "", "", "target port")
	probeCli.Flags().StringP("plugin", "x", "", "plugin to scan(e.g. oxid,ms17010)")
	probeCli.Flags().StringP("service", "s", "", "service ip(in the nmap format: 10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24)")
	probeCli.Flags().StringP("service-file", "S", "", "File to probe for(e.g. ip.txt)")

	if err := probeCli.MarkFlagRequired("plugin"); err != nil {
		gologger.Errorf("error on marking flag as required: %v", err)
	}

	rootCmd.AddCommand(probeCli)
}
