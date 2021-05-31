package cmd

import (
	"cube/cli"
	"cube/log"
	"cube/model"
	Plugins "cube/plugins"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var sqlCmdCli *cobra.Command

func runSqlcmd(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseProbeOptions()
	_, key := Plugins.ProbeFuncMap[opt.ScanPlugin]
	if !key {
		log.Fatalf("Available Plugins: %s", strings.Join(Plugins.ProbeKeys, ","))
		os.Exit(2)
	}
	cli.StartProbeTask(opt, globalopts)
}

func parseSqlcmdOptions() (*model.GlobalOptions, *model.SqlcmdOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}

	sqlcmdOption := model.NewSqlcmdOptions()

	sqlcmdOption.Ip, err = sqlCmdCli.Flags().GetString("ip")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	sqlcmdOption.Port, err = sqlCmdCli.Flags().GetInt("port")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for scan port: %v", err)
	}

	sqlcmdOption.User, err = sqlCmdCli.Flags().GetString("user")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for user: %v", err)
	}

	sqlcmdOption.Password, err = sqlCmdCli.Flags().GetString("password")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for Password: %v", err)
	}

	sqlcmdOption.CrackPlugin, err = sqlCmdCli.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-ip: %w", err)
	}
	sqlcmdOption.Query, err = sqlCmdCli.Flags().GetString("query")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}
	return globalOpts, sqlcmdOption, nil
}

func init() {
	sqlCmdCli = &cobra.Command{
		Use:   "sqlcmd",
		Short: "exec sql query or cmd",
		Run:   runProbe,
	}

	sqlCmdCli.Flags().StringP("ip", "i", "", "ip (e.g. 192.168.1.2)")
	sqlCmdCli.Flags().IntP("port", "p", 22, "target port")
	sqlCmdCli.Flags().StringP("user", "u", "", "login account")
	sqlCmdCli.Flags().StringP("password", "P", "", "login password")
	sqlCmdCli.Flags().StringP("plugin", "x", "", "plugin to use(e.g. SSH)")
	sqlCmdCli.Flags().StringP("query", "e", "", "string to query or exec")

	//if err := probeCmd.MarkPersistentFlagRequired("plugin"); err != nil {
	//	log.Fatalf("on marking flag as required: %v", err)
	//	//log.Fatalf("error on marking flag as required: %v", err)
	//}

	//probeCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
	//
	//}
	rootCmd.AddCommand(sqlCmdCli)
}
