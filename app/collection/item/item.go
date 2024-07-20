package item

import (
	"net/url"
	"path"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/lib/svc"
	"github.com/kyleu/todoforge/app/util"
)

const DefaultRoute = "/collection/item"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(paths...)
}

var _ svc.Model = (*Item)(nil)

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
	if xx := i.Name; xx != "" {
		return xx
	}
	return i.String()
}

func Random() *Item {
	return &Item{
		ID:           util.UUID(),
		CollectionID: util.UUID(),
		Name:         util.RandomString(12),
		Created:      util.TimeCurrent(),
	}
}

func (i *Item) Strings() []string {
	return []string{i.ID.String(), i.CollectionID.String(), i.Name, util.TimeToFull(&i.Created)}
}

func (i *Item) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{i.Strings()}
}

func (i *Item) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return path.Join(append(paths, url.QueryEscape(i.ID.String()))...)
}

func (i *Item) ToData() []any {
	return []any{i.ID, i.CollectionID, i.Name, i.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "collectionID", Title: "Collection ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
