package cmd

import (
	"cube/cubelib"
	"cube/log"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
)

var crackCli *cobra.Command

func runCrack(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseCrackOptions()

	cubelib.StartCrackTask(opt, globalopts)
}

func parseCrackOptions() (*model.GlobalOptions, *model.CrackOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}

	crackOption := model.NewCrackOptions()

	crackOption.Ip, err = crackCli.Flags().GetString("ip")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	crackOption.IpFile, err = crackCli.Flags().GetString("ip-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	crackOption.User, err = crackCli.Flags().GetString("user")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for Password: %v", err)
	}

	crackOption.UserFile, err = crackCli.Flags().GetString("user-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}

	crackOption.Pass, err = crackCli.Flags().GetString("pass")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}

	crackOption.PassFile, err = crackCli.Flags().GetString("pass-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}

	crackOption.Port, err = crackCli.Flags().GetString("port")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}

	crackOption.CrackPlugin, err = crackCli.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for scan plugin: %w", err)
	}

	return globalOpts, crackOption, nil
}

func init() {
	crackCli = &cobra.Command{
		Use:   "crack",
		Short: "crack service password (e.g. ssh,mssql,redis,mysql)",
		Run:   runCrack,
		Example: `cube crack -u root -p root -i 192.168.1.1 -x ssh
cube crack -u root -p root -i 192.168.1.1 -x ssh --port 2222
cube crack -u root -p root -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh,mysql
		`,
	}

	crackCli.Flags().StringP("ip", "i", "", "ip (e.g. 192.168.2.1")
	crackCli.Flags().StringP("ip-file", "", "", "login account")
	crackCli.Flags().StringP("user", "u", "", "login password")
	crackCli.Flags().StringP("user-file", "", "", "string to query or exec")
	crackCli.Flags().StringP("pass", "p", "", "login password")
	crackCli.Flags().StringP("pass-file", "", "", "string to query or exec")
	crackCli.Flags().StringP("port", "", "", "login password")
	crackCli.Flags().StringP("plugin", "x", "", "crack plugin")
	if err := crackCli.MarkFlagRequired("plugin"); err != nil {
		log.Errorf("error on marking flag as required: %v", err)
	}

	rootCmd.AddCommand(crackCli)
}
