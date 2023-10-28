// Package item - Content managed by Project Forge, see [projectforge.md] for details.
package item

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/util"
)

var (
	table         = "item"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "collection_id", "name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = strings.Join(columnsQuoted, ", ")
)

type row struct {
	ID           uuid.UUID `db:"id" json:"id"`
	CollectionID uuid.UUID `db:"collection_id" json:"collection_id"`
	Name         string    `db:"name" json:"name"`
	Created      time.Time `db:"created" json:"created"`
}

func (r *row) ToItem() *Item {
	if r == nil {
		return nil
	}
	return &Item{
		ID:           r.ID,
		CollectionID: r.CollectionID,
		Name:         r.Name,
		Created:      r.Created,
	}
}

type rows []*row

func (x rows) ToItems() Items {
	return lo.Map(x, func(d *row, _ int) *Item {
		return d.ToItem()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
