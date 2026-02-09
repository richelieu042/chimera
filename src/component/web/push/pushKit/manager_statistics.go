package pushKit

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/richelieu042/chimera/v3/src/prettyKit"
)

// GetStatistics
/*
e.g.
+----------+-------------------+
| TYPE     | COUNT OF CHANNELS |
+----------+-------------------+
| idMap    |                 0 |
| bsidMap  |                 0 |
| userMap  |                 0 |
| groupMap |                 0 |
+----------+-------------------+
*/
func GetStatistics() string {
	writer := prettyKit.NewTableWriter()

	writer.AppendHeader(table.Row{"Type", "Count of channels"})
	writer.AppendRow(table.Row{"idMap", GetCountOfIdMap()})
	writer.AppendRow(table.Row{"bsidMap", GetCountOfBsidMap()})
	writer.AppendRow(table.Row{"userMap", GetCountOfUserMap()})
	writer.AppendRow(table.Row{"groupMap", GetCountOfGroupMap()})

	return writer.Render()
}
