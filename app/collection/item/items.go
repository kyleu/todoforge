package item

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/util"
)

type Items []*Item

func (i Items) Get(id uuid.UUID) *Item {
	return lo.FindOrElse(i, nil, func(x *Item) bool {
		return x.ID == id
	})
}

func (i Items) IDs() []uuid.UUID {
	return lo.Map(i, func(xx *Item, _ int) uuid.UUID {
		return xx.ID
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

func (i Items) GetByID(id uuid.UUID) Items {
	return lo.Filter(i, func(xx *Item, _ int) bool {
		return xx.ID == id
	})
}

func (i Items) GetByIDs(ids ...uuid.UUID) Items {
	return lo.Filter(i, func(xx *Item, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (i Items) CollectionIDs() []uuid.UUID {
	return lo.Map(i, func(xx *Item, _ int) uuid.UUID {
		return xx.CollectionID
	})
}

func (i Items) GetByCollectionID(collectionID uuid.UUID) Items {
	return lo.Filter(i, func(xx *Item, _ int) bool {
		return xx.CollectionID == collectionID
	})
}

func (i Items) GetByCollectionIDs(collectionIDs ...uuid.UUID) Items {
	return lo.Filter(i, func(xx *Item, _ int) bool {
		return lo.Contains(collectionIDs, xx.CollectionID)
	})
}

func (i Items) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), lo.Map(i, func(x *Item, _ int) []string {
		return x.Strings()
	})
}

func (i Items) Random() *Item {
	return util.RandomElement(i)
}

func (i Items) Clone() Items {
	return slices.Clone(i)
}
