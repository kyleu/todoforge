// Content managed by Project Forge, see [projectforge.md] for details.
package item

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Items []*Item

func (i Items) Get(id uuid.UUID) *Item {
	return lo.FindOrElse(i, nil, func(x *Item) bool {
		return x.ID == id
	})
}

func (i Items) GetByIDs(ids ...uuid.UUID) Items {
	return lo.Filter(i, func(x *Item, _ int) bool {
		return lo.Contains(ids, x.ID)
	})
}

func (i Items) IDs() []uuid.UUID {
	return lo.Map(i, func(x *Item, _ int) uuid.UUID {
		return x.ID
	})
}

func (i Items) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(i)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(i, func(x *Item, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (i Items) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(i)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(i, func(x *Item, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (i Items) Clone() Items {
	return slices.Clone(i)
}
