// Package item - Content managed by Project Forge, see [projectforge.md] for details.
package item

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/util"
)

type Item struct {
	ID           uuid.UUID `json:"id,omitempty"`
	CollectionID uuid.UUID `json:"collectionID,omitempty"`
	Name         string    `json:"name,omitempty"`
	Created      time.Time `json:"created,omitempty"`
}

func New(id uuid.UUID) *Item {
	return &Item{ID: id}
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

func Random() *Item {
	return &Item{
		ID:           util.UUID(),
		CollectionID: util.UUID(),
		Name:         util.RandomString(12),
		Created:      util.TimeCurrent(),
	}
}

func (i *Item) WebPath() string {
	return "/collection/item/" + i.ID.String()
}

func (i *Item) ToData() []any {
	return []any{i.ID, i.CollectionID, i.Name, i.Created}
}
