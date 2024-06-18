package collection

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/lib/svc"
	"github.com/kyleu/todoforge/app/util"
)

var _ svc.Model = (*Collection)(nil)

type Collection struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func New(id uuid.UUID) *Collection {
	return &Collection{ID: id}
}

func (c *Collection) Clone() *Collection {
	return &Collection{c.ID, c.Name, c.Created}
}

func (c *Collection) String() string {
	return c.ID.String()
}

func (c *Collection) TitleString() string {
	return c.Name
}

func Random() *Collection {
	return &Collection{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func (c *Collection) Strings() []string {
	return []string{c.ID.String(), c.Name, util.TimeToFull(&c.Created)}
}

func (c *Collection) ToCSV() ([]string, [][]string) {
	return FieldDescs.Keys(), [][]string{c.Strings()}
}

func (c *Collection) WebPath() string {
	return "/collection/" + c.ID.String()
}

func (c *Collection) ToData() []any {
	return []any{c.ID, c.Name, c.Created}
}

var FieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
