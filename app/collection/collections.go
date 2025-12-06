package collection

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/util"
)

type Collections []*Collection

func (c Collections) Get(id uuid.UUID) *Collection {
	return lo.FindOrElse(c, nil, func(x *Collection) bool {
		return x.ID == id
	})
}

func (c Collections) IDs() []uuid.UUID {
	return lo.Map(c, func(xx *Collection, _ int) uuid.UUID {
		return xx.ID
	})
}

func (c Collections) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(c)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(c, func(x *Collection, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (c Collections) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(c)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(c, func(x *Collection, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (c Collections) GetByID(id uuid.UUID) Collections {
	return lo.Filter(c, func(xx *Collection, _ int) bool {
		return xx.ID == id
	})
}

func (c Collections) GetByIDs(ids ...uuid.UUID) Collections {
	return lo.Filter(c, func(xx *Collection, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (c Collections) ToMap() map[uuid.UUID]*Collection {
	return lo.SliceToMap(c, func(xx *Collection) (uuid.UUID, *Collection) {
		return xx.ID, xx
	})
}

func (c Collections) ToMaps() []util.ValueMap {
	return lo.Map(c, func(xx *Collection, _ int) util.ValueMap {
		return xx.ToMap()
	})
}

func (c Collections) ToOrderedMaps() util.OrderedMaps[any] {
	return lo.Map(c, func(x *Collection, _ int) *util.OrderedMap[any] {
		return x.ToOrderedMap()
	})
}

func (c Collections) ToCSV() ([]string, [][]string) {
	return CollectionFieldDescs.Keys(), lo.Map(c, func(x *Collection, _ int) []string {
		return x.Strings()
	})
}

func (c Collections) Random() *Collection {
	return util.RandomElement(c)
}

func (c Collections) Clone() Collections {
	return lo.Map(c, func(xx *Collection, _ int) *Collection {
		return xx.Clone()
	})
}
