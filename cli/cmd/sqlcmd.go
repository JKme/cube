package cmd

import (
	"cube/core"
	"cube/core/sqlcmdmodule"
	"cube/gologger"
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
	sqlcmdmodule.StartSqlcmd(opt, globalopts)
}

func parseSqlcmdOptions() (*core.GlobalOption, *sqlcmdmodule.SqlcmdOption, error) {
	globalOpts, err := parseGlobalOptions()
	if err != nil {
		return nil, nil, err
	}

	sqlcmdOption := sqlcmdmodule.NewSqlcmdOption()

	sqlcmdOption.Ip, err = sqlCli.Flags().GetString("service")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}
	sqlcmdOption.Port, err = sqlCli.Flags().GetString("port")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	sqlcmdOption.User, err = sqlCli.Flags().GetString("login")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
	}

	sqlcmdOption.Password, err = sqlCli.Flags().GetString("password")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for Password: %v", err)
	}

	sqlcmdOption.Name, err = sqlCli.Flags().GetString("plugin")
	if err != nil {
		return nil, nil, fmt.Errorf("invalid value for plugin: %v", err)
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
		Short: "exec sql or cmd",
		Long:  "Hello",
		Run:   runSqlcmd,
		Example: `cube sqlcmd -x mysql -l root -p root -e "whoami"
cube sqlcmd -x mysql -l root -p root -e "clear"
cube sqlcmd -x mssql -l sa -p sa -e "whoami" --port 4134
		`,
	}

	sqlCli.Flags().StringP("service", "s", "", "service ip")
	sqlCli.Flags().StringP("login", "l", "", "login user")
	sqlCli.Flags().StringP("password", "p", "", "login password")
	sqlCli.Flags().StringP("query", "e", "", "string to query or exec")
	sqlCli.Flags().StringP("plugin", "x", "", fmt.Sprintf("avaliable plugin: %s", 111))
	sqlCli.Flags().StringP("port", "", "", "if the service is on a different default port, define it here")

	if err := sqlCli.MarkFlagRequired("plugin"); err != nil {
		gologger.Errorf("error on marking flag as required: %v", err)
	}
	if err := sqlCli.MarkFlagRequired("service"); err != nil {
		gologger.Errorf("error on marking flag as required: %v", err)
	}
	if err := sqlCli.MarkFlagRequired("login"); err != nil {
		gologger.Errorf("error on marking flag as required: %v", err)
	}
	if err := sqlCli.MarkFlagRequired("password"); err != nil {
		gologger.Errorf("error on marking flag as required: %v", err)
	}

	rootCmd.AddCommand(sqlCli)
}
