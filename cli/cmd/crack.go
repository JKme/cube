package cmd

import (
	"cube/core"
	"cube/core/crackmodule"
	"cube/gologger"
	"fmt"
	"github.com/spf13/cobra"
)

var crackCli *cobra.Command

func runCrack(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseCrackOptions()

	if len(opt.User+opt.UserFile+opt.Pass+opt.PassFile) > 0 { //当使用自定义用户密码的时候，判断是否同时指定了User和Password
		if len(opt.User)+len(opt.UserFile) == 0 || len(opt.Pass)+len(opt.PassFile) == 0 {
			gologger.Errorf("Please set login name -l/-L and password -p/-P flag)")
		}
	}

	crackmodule.StartCrack(opt, globalopts)
}

func parseCrackOptions() (*core.GlobalOption, *crackmodule.CrackOption, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}

	crackOption := crackmodule.NewCrackOptions()

	crackOption.Ip, err = crackCli.Flags().GetString("service")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	crackOption.IpFile, err = crackCli.Flags().GetString("service-file")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	crackOption.User, err = crackCli.Flags().GetString("login")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for Password: %v", err)
	}

	crackOption.UserFile, err = crackCli.Flags().GetString("login-file")
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

	crackOption.PluginName, err = crackCli.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for scan plugin: %w", err)
	}
	return globalOpts, crackOption, nil
}

func init() {
	crackCli = &cobra.Command{
		Use:   "crack",
		Long:  "Hello", //TODO
		Short: "crack service password",
		Run:   runCrack,
		Example: `cube crack -u root -p root -i 192.168.1.1 -x ssh
cube crack -u root -p root -i 192.168.1.1 -x ssh -port 2222
cube crack -u root,ubuntu -p 123,000111,root -x ssh -i 192.168.1.1
cube crack -u root -p root -i 192.168.1.1/24 -x ssh
cube crack -u root -pass-file pass.txt -i 192.168.1.1/24 -x ssh
cube crack -u root -pass-file pass.txt -i 192.168.1.1/24 -x ssh,mysql
cube crack -u root -pass-file pass.txt -i http://127.0.0.1:8080 -x httpbasic
cube crack -u root -pass-file pass.txt -i http://127.0.0.1:8080 -x phpmyadmin
		`,
	}

	crackCli.Flags().StringP("service", "s", "", "service ip(in the nmap format: 10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24)")
	crackCli.Flags().StringP("service-file", "S", "", "service ip file")
	crackCli.Flags().StringP("login", "l", "", "login user")
	crackCli.Flags().StringP("login-file", "L", "", "login user file")
	crackCli.Flags().StringP("pass", "p", "", "login password")
	crackCli.Flags().StringP("pass-file", "P", "", "login password file")
	crackCli.Flags().StringP("port", "", "", "if the service is on a different default port, define it here")
	crackCli.Flags().StringP("plugin", "x", "", fmt.Sprintf("avaliable plugin: %s", 111))
	if err := crackCli.MarkFlagRequired("plugin"); err != nil {
		gologger.Errorf("error on marking flag as required: %v", err)
	}

	rootCmd.AddCommand(crackCli)
}
