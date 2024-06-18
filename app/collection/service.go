package collection

import (
	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/lib/svc"
)

var (
	_ svc.ServiceID[*Collection, Collections, uuid.UUID] = (*Service)(nil)
	_ svc.ServiceSearch[Collections]                     = (*Service)(nil)
)

type Service struct {
	db *database.Service
}

func NewService(db *database.Service) *Service {
	filter.AllowedColumns["collection"] = columns
	return &Service{db: db}
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("collection", &filter.Ordering{Column: "name"})
}
