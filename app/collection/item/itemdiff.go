// Package item - Content managed by Project Forge, see [projectforge.md] for details.
package item

import "github.com/kyleu/todoforge/app/util"

func (i *Item) Diff(ix *Item) util.Diffs {
	var diffs util.Diffs
	if i.ID != ix.ID {
		diffs = append(diffs, util.NewDiff("id", i.ID.String(), ix.ID.String()))
	}
	if i.CollectionID != ix.CollectionID {
		diffs = append(diffs, util.NewDiff("collectionID", i.CollectionID.String(), ix.CollectionID.String()))
	}
	if i.Name != ix.Name {
		diffs = append(diffs, util.NewDiff("name", i.Name, ix.Name))
	}
	if i.Created != ix.Created {
		diffs = append(diffs, util.NewDiff("created", i.Created.String(), ix.Created.String()))
	}
	return diffs
}
