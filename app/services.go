package app

import (
	"context"

	"github.com/pkg/errors"

	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/lib/audit"
	"github.com/kyleu/todoforge/app/lib/database/migrate"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/queries/migrations"
)

type Services struct {
	Collection *collection.Service
	Item       *item.Service
	Audit      *audit.Service
}

func NewServices(ctx context.Context, st *State, logger util.Logger) (*Services, error) {
	migrations.LoadMigrations(st.Debug)
	err := migrate.Migrate(ctx, st.DB, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to run database migrations")
	}
	return &Services{
		Collection: collection.NewService(st.DB),
		Item:       item.NewService(st.DB),
		Audit:      audit.NewService(st.DB, logger),
	}, nil
}

func (s *Services) Close(_ context.Context, logger util.Logger) error {
	return nil
}
