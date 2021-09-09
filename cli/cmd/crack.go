package cmd

import (
	"cube/cubelib"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var crackCli *cobra.Command

func runCrack(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseCrackOptions()

	if len(opt.User+opt.UserFile+opt.Pass+opt.PassFile) > 0 { //当使用自定义用户密码的时候，判断是否同时指定了User和Password
		if len(opt.User)+len(opt.UserFile) == 0 || len(opt.Pass)+len(opt.PassFile) == 0 {
			log.Error("-h for Help, Please set User(-u/--user-file) and Password(-p/--pass-file) flag")
		}
	}

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
		Short: fmt.Sprintf("-x ALL will load plugins: [%s]\nAnother plugins: [%s]", strings.Join(Plugins.CrackKeys, ","), strings.Join(Plugins.CrackFuncExclude, ",")),
		Run:   runCrack,
		Example: `cube crack -u root -p root -i 192.168.1.1 -x ssh
cube crack -u root -p root -i 192.168.1.1 -x ssh --port 2222
cube crack -u root,ubuntu -p 123,000111,root -x ssh -i 192.168.1.1
cube crack -u root -p root -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh
cube crack -u root --pass-file pass.txt -i 192.168.1.1/24 -x ssh,mysql
cube crack -u root --pass-file pass.txt -i http://127.0.0.1:8080 -x httpbasic
cube crack -u root --pass-file pass.txt -i http://127.0.0.1:8080 -x phpmyadmin
		`,
	}

	crackCli.Flags().StringP("ip", "i", "", "ip (e.g. 10.0.0.1, 10.0.0.5-10, 192.168.1.*, 192.168.10.0/24, in the nmap format.)")
	crackCli.Flags().StringP("ip-file", "", "", "ip file")
	crackCli.Flags().StringP("user", "u", "", "login user")
	crackCli.Flags().StringP("user-file", "", "", "login user file")
	crackCli.Flags().StringP("pass", "p", "", "login password")
	crackCli.Flags().StringP("pass-file", "", "", "login password file")
	crackCli.Flags().StringP("port", "", "", "if the service is on a different default port, define it here")
	crackCli.Flags().StringP("plugin", "x", "", fmt.Sprintf("avaliable plugin: %s", strings.Join(Plugins.CrackKeys, ",")))
	if err := crackCli.MarkFlagRequired("plugin"); err != nil {
		log.Errorf("error on marking flag as required: %v", err)
	}

	rootCmd.AddCommand(crackCli)
}
