package collection

import (
	"net/url"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/lib/svc"
	"github.com/kyleu/todoforge/app/util"
)

const DefaultRoute = "/collection"

func Route(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(paths...)
}

var _ svc.Model = (*Collection)(nil)

type Collection struct {
	ID      uuid.UUID `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

func NewCollection(id uuid.UUID) *Collection {
	return &Collection{ID: id}
}

func (c *Collection) Clone() *Collection {
	return &Collection{c.ID, c.Name, c.Created}
}

func (c *Collection) String() string {
	return c.ID.String()
}

func (c *Collection) TitleString() string {
	if xx := c.Name; xx != "" {
		return xx
	}
	return c.String()
}

func RandomCollection() *Collection {
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
	return CollectionFieldDescs.Keys(), [][]string{c.Strings()}
}

func (c *Collection) WebPath(paths ...string) string {
	if len(paths) == 0 {
		paths = []string{DefaultRoute}
	}
	return util.StringPath(append(paths, url.QueryEscape(c.ID.String()))...)
}

func (c *Collection) Breadcrumb(extra ...string) string {
	return c.TitleString() + "||" + c.WebPath(extra...) + "**archive"
}

func (c *Collection) ToData() []any {
	return []any{c.ID, c.Name, c.Created}
}

var CollectionFieldDescs = util.FieldDescs{
	{Key: "id", Title: "ID", Description: "", Type: "uuid"},
	{Key: "name", Title: "Name", Description: "", Type: "string"},
	{Key: "created", Title: "Created", Description: "", Type: "timestamp"},
}
