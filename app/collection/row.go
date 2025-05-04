package collection

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/util"
)

var (
	table         = "collection"
	tableQuoted   = fmt.Sprintf("%q", table)
	columns       = []string{"id", "name", "created"}
	columnsQuoted = util.StringArrayQuoted(columns)
	columnsString = util.StringJoin(columnsQuoted, ", ")
)

type row struct {
	ID      uuid.UUID `db:"id" json:"id"`
	Name    string    `db:"name" json:"name"`
	Created time.Time `db:"created" json:"created"`
}

func (r *row) ToCollection() *Collection {
	if r == nil {
		return nil
	}
	return &Collection{
		ID:      r.ID,
		Name:    r.Name,
		Created: r.Created,
	}
}

type rows []*row

func (x rows) ToCollections() Collections {
	return lo.Map(x, func(d *row, _ int) *Collection {
		return d.ToCollection()
	})
}

func defaultWC(idx int) string {
	return fmt.Sprintf("\"id\" = $%d", idx+1)
}
