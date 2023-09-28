// Package item - Content managed by Project Forge, see [projectforge.md] for details.
package item

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/util"
)

type Item struct {
	ID           uuid.UUID `json:"id"`
	CollectionID uuid.UUID `json:"collectionID"`
	Name         string    `json:"name"`
	Created      time.Time `json:"created"`
}

func New(id uuid.UUID) *Item {
	return &Item{ID: id}
}

func Random() *Item {
	return &Item{
		ID:           util.UUID(),
		CollectionID: util.UUID(),
		Name:         util.RandomString(12),
		Created:      util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Item, error) {
	ret := &Item{}
	var err error
	if setPK {
		retID, e := m.ParseUUID("id", true, true)
		if e != nil {
			return nil, e
		}
		if retID != nil {
			ret.ID = *retID
		}
		// $PF_SECTION_START(pkchecks)$
		// $PF_SECTION_END(pkchecks)$
	}
	retCollectionID, e := m.ParseUUID("collectionID", true, true)
	if e != nil {
		return nil, e
	}
	if retCollectionID != nil {
		ret.CollectionID = *retCollectionID
	}
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
}

func (i *Item) Clone() *Item {
	return &Item{i.ID, i.CollectionID, i.Name, i.Created}
}

func (i *Item) String() string {
	return i.ID.String()
}

func (i *Item) TitleString() string {
	return i.Name
}

func (i *Item) WebPath() string {
	return "/collection/item/" + i.ID.String()
}

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

func (i *Item) ToData() []any {
	return []any{i.ID, i.CollectionID, i.Name, i.Created}
}
