// Package item - Content managed by Project Forge, see [projectforge.md] for details.
package item

import (
	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/lib/filter"
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["item"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("item", &filter.Ordering{Column: "name"})
}
