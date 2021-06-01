package cmd

import (
	"cube/cli"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
)

var sqlCmdCli *cobra.Command

func runSqlcmd(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseSqlcmdOptions()
	//_, key := Plugins.ProbeFuncMap[opt.ScanPlugin]
	//if !key {
	//	log.Fatalf("Available Plugins: %s", strings.Join(Plugins.ProbeKeys, ","))
	//	os.Exit(2)
	//}
	cli.StartSqlcmdTask(opt, globalopts)
}

func parseSqlcmdOptions() (*model.GlobalOptions, *model.SqlcmdOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}

	sqlcmdOption := model.NewSqlcmdOptions()

	sqlcmdOption.Service, err = sqlCmdCli.Flags().GetString("service")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	sqlcmdOption.User, err = sqlCmdCli.Flags().GetString("user")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	sqlcmdOption.Password, err = sqlCmdCli.Flags().GetString("password")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for Password: %v", err)
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
		Run:   runSqlcmd,
	}

	sqlCmdCli.Flags().StringP("service", "x", "", "ip (e.g. ssh://192.168.0.0:2200)")
	sqlCmdCli.Flags().StringP("user", "u", "", "login account")
	sqlCmdCli.Flags().StringP("password", "p", "", "login password")
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
