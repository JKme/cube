package cmd

import (
	"cube/cubelib"
	"cube/log"
	Plugins "cube/plugins"
	"strings"

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
		Long:  fmt.Sprintf("-x ALL will load plugins: [%s]\nAnother plugins: [%s]", strings.Join(Plugins.ProbeKeys, ","), strings.Join(Plugins.ProbeFuncExclude, ",")),
		Short: "probe pentest env",
		Run:   runProbe,
		Example: `cube probe -i 192.168.1.1 -x oxid
cube probe -i 192.168.1.1 -x oxid,zookeeper,ms17010
cube probe -i 192.168.1.1/24 -x ALL
		`,
	}

	probeCli.Flags().StringP("port", "p", "", "target port")
	probeCli.Flags().StringP("plugin", "x", "", "plugin to scan(e.g. OXID)")
	probeCli.Flags().StringP("ip", "i", "", "ip (e.g. 10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24, in the nmap format.)")
	probeCli.Flags().StringP("ip-file", "", "", "File to probe for(e.g. ip.txt)")

	if err := crackCli.MarkFlagRequired("plugin"); err != nil {
		log.Errorf("error on marking flag as required: %v", err)
	}

	rootCmd.AddCommand(probeCli)
}
