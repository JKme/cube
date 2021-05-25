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

func parseProbeOptions() (*model.ProbeOptions, *model.GlobalOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}
	propePlugin := model.NewProbeOptions()

	propePlugin.Port, err = probeCmd.Flags().GetInt("port")
	propePlugin.Port, err = probeCmd.Flags().GetInt("func")
	propePlugin.Port, err = probeCmd.Flags().GetInt("target-ip")
	propePlugin.Port, err = probeCmd.Flags().GetInt("target-file")

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
