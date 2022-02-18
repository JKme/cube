package crackmodule

import (
	"cube/core/crackmodule/plugins"
	"github.com/olekukonko/tablewriter"
	"os"
)

func CrackHelpTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Func", "Port", "Load By ALL"})
	for _, k := range plugins.CrackKeys {
		table.Append([]string{k, plugins.GetPort(k), plugins.GetLoadStatus(k)})
		table.SetRowLine(true)
	}
	table.Render()
}
