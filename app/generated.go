// Package app - Content managed by Project Forge, see [projectforge.md] for details.
package app

import (
	"context"

	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/lib/audit"
	"github.com/kyleu/todoforge/app/util"
)

type GeneratedServices struct {
	Collection *collection.Service
	Item       *item.Service
}

func initGeneratedServices(ctx context.Context, st *State, audSvc *audit.Service, logger util.Logger) GeneratedServices {
	return GeneratedServices{
		Collection: collection.NewService(st.DB),
		Item:       item.NewService(st.DB),
	}
}
