// Package collection - Content managed by Project Forge, see [projectforge.md] for details.
package collection

import (
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/util"
)

type Collection struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}

func New(id uuid.UUID) *Collection {
	return &Collection{ID: id}
}

func Random() *Collection {
	return &Collection{
		ID:      util.UUID(),
		Name:    util.RandomString(12),
		Created: util.TimeCurrent(),
	}
}

func FromMap(m util.ValueMap, setPK bool) (*Collection, error) {
	ret := &Collection{}
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
	ret.Name, err = m.ParseString("name", true, true)
	if err != nil {
		return nil, err
	}
	// $PF_SECTION_START(extrachecks)$
	// $PF_SECTION_END(extrachecks)$
	return ret, nil
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

func (c *Collection) WebPath() string {
	return "/collection/" + c.ID.String()
}

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

func (c *Collection) ToData() []any {
	return []any{c.ID, c.Name, c.Created}
}
