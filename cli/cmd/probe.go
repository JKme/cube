package cmd

import (
	"cube/cubelib"
	"cube/log"

	//"cube/log"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
)

var probeCli *cobra.Command

func runProbe(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseProbeOptions()

	cubelib.StartProbeTask(opt, globalopts)
}

func parseProbeOptions() (*model.GlobalOptions, *model.ProbeOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}
	probeOption := model.NewProbeOptions()

	probeOption.ScanPlugin, err = probeCli.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	probeOption.Port, err = probeCli.Flags().GetString("port")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for scan port: %v", err)
	}

	probeOption.Target, err = probeCli.Flags().GetString("ip")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-ip: %w", err)
	}
	probeOption.TargetFile, err = probeCli.Flags().GetString("ip-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}
	return globalOpts, probeOption, nil
}

func init() {
	probeCli = &cobra.Command{
		Use:   "probe",
		Short: "collect pentest environment information",
		Run:   runProbe,
		Example: `cube probe -i 192.168.1.1 -x oxid
cube probe -i 192.168.1.1 -x oxid,zookeeper,ms17010
cube probe -i 192.168.1.1/24 -x ALL
		`,
	}

	probeCli.Flags().StringP("port", "p", "", "target port")
	probeCli.Flags().StringP("plugin", "x", "", "plugin to scan(e.g. OXID)")
	probeCli.Flags().StringP("ip", "i", "", "ip range to probe for(e.g. 192.168.1.1/24)")
	probeCli.Flags().StringP("ip-file", "", "", "File to probe for(e.g. ip.txt)")

	if err := crackCli.MarkFlagRequired("plugin"); err != nil {
		log.Errorf("error on marking flag as required: %v", err)
	}

	rootCmd.AddCommand(probeCli)
}
