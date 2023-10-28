// Package collection - Content managed by Project Forge, see [projectforge.md] for details.
package collection

import "github.com/kyleu/todoforge/app/util"

func (c *Collection) Diff(cx *Collection) util.Diffs {
	var diffs util.Diffs
	if c.ID != cx.ID {
		diffs = append(diffs, util.NewDiff("id", c.ID.String(), cx.ID.String()))
	}
	if c.Name != cx.Name {
		diffs = append(diffs, util.NewDiff("name", c.Name, cx.Name))
	}
	if c.Created != cx.Created {
		diffs = append(diffs, util.NewDiff("created", c.Created.String(), cx.Created.String()))
	}
	return diffs
}
