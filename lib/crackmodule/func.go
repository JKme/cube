package crackmodule

import (
	"cube/lib/crackmodule/plugins"
	"github.com/olekukonko/tablewriter"
	"os"
)

func CrackHelpTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Func", "Port", "ALL"})

	for _, k := range plugins.CRACK_KEYS {
		c := plugins.Crack{
			Name: k,
		}
		table.Append([]string{k, c.GetPort(), c.GetLoadStatus()})
		table.SetRowLine(true)
	}
	table.Render()
}
