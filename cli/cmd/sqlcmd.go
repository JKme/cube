package cmd

import (
	"cube/cubelib"
	"cube/model"
	"fmt"
	"github.com/spf13/cobra"
)

var sqlCli *cobra.Command

func runSqlcmd(cmd *cobra.Command, args []string) {
	globalopts, opt, _ := parseSqlcmdOptions()
	//_, key := Plugins.ProbeFuncMap[opt.ScanPlugin]
	//if !key {
	//	log.Fatalf("Available Plugins: %s", strings.Join(Plugins.ProbeKeys, ","))
	//	os.Exit(2)
	//}
	cubelib.StartSqlcmdTask(opt, globalopts)
}

func parseSqlcmdOptions() (*model.GlobalOptions, *model.SqlcmdOptions, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}

	sqlcmdOption := model.NewSqlcmdOptions()

	sqlcmdOption.Service, err = sqlCli.Flags().GetString("service")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	sqlcmdOption.User, err = sqlCli.Flags().GetString("user")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	sqlcmdOption.Password, err = sqlCli.Flags().GetString("password")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for Password: %v", err)
	}

	sqlcmdOption.Query, err = sqlCli.Flags().GetString("query")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for target-file: %w", err)
	}
	return globalOpts, sqlcmdOption, nil
}

func init() {
	sqlCli = &cobra.Command{
		Use:   "sqlcmd",
		Short: "exec sql query or cmd",
		Run:   runSqlcmd,
		Example: `cube sqlcmd -x ssh://192.168.0.0:2200 -uroot -proot -e "whoami"
		`,
	}

	sqlCli.Flags().StringP("service", "x", "", "ip (e.g. ssh://192.168.0.0:2200 or ssh://127.0.0.1))")
	sqlCli.Flags().StringP("user", "u", "", "login account")
	sqlCli.Flags().StringP("password", "p", "", "login password")
	sqlCli.Flags().StringP("query", "e", "", "string to query or exec")

	rootCmd.AddCommand(sqlCli)
}
