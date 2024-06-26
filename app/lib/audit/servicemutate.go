package audit

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/util"
)

func (s *Service) Create(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Audit) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLInsert(tableQuoted, columnsQuoted, len(models), s.db.Type)
	vals := lo.FlatMap(models, func(arg *Audit, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}

func (s *Service) Update(ctx context.Context, tx *sqlx.Tx, model *Audit, logger util.Logger) error {
	wc := fmt.Sprintf("%s = %s", s.db.Type.Quoted("id"), s.db.Type.PlaceholderFor(11))
	q := database.SQLUpdate(tableQuoted, columnsQuoted, wc, s.db.Type)
	data := model.ToData()
	data = append(data, model.ID)
	_, ret := s.db.Update(ctx, q, tx, 1, logger, data...)
	return ret
}

func (s *Service) Save(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Audit) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLUpsert(tableQuoted, columnsQuoted, len(models), []string{"id"}, columns, s.db.Type)
	data := lo.FlatMap(models, func(model *Audit, _ int) []any {
		return model.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, data...)
}

func (s *Service) Delete(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) error {
	q := database.SQLDelete(tableQuoted, defaultWC, s.db.Type)
	_, err := s.db.Delete(ctx, q, tx, 1, logger, id)
	return err
}
