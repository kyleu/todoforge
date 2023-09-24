// Content managed by Project Forge, see [projectforge.md] for details.
package audit

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/util"
)

func (s *Service) RecordsForAudits(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, auditIDs ...uuid.UUID) (Records, error) {
	params = params.Sanitize("audit_record", &filter.Ordering{Column: "occurred"})
	wc := database.SQLInClause("audit_id", len(auditIDs), 0, "$")
	q := database.SQLSelect(recordColumnsString, recordTableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := recordRows{}
	vals := make([]any, 0, len(auditIDs))
	strs := make([]string, 0, len(auditIDs))
	lo.ForEach(auditIDs, func(auditID uuid.UUID, _ int) {
		vals = append(vals, auditID.String())
		strs = append(strs, auditID.String())
	})
	err := s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get records for audits [%s]", strings.Join(strs, ", "))
	}
	return ret.ToRecords(), nil
}

func (s *Service) RecordsForModel(ctx context.Context, tx *sqlx.Tx, t string, pk string, params *filter.Params, logger util.Logger) (Records, error) {
	params = params.Sanitize("audit_record", &filter.Ordering{Column: "occurred"})
	wc := "t = $1 and pk = $2"
	q := database.SQLSelect(recordColumnsString, recordTableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := recordRows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, t, pk)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get records for model [%s:%s]", t, pk)
	}
	return ret.ToRecords(), nil
}

func (s *Service) GetRecord(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Record, error) {
	q := database.SQLSelectSimple(recordColumnsString, recordTableQuoted, s.db.Placeholder(), "id = $1")
	ret := &recordRow{}
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get audit record by id [%s]", id.String())
	}
	return ret.ToRecord(), nil
}

func (s *Service) CreateRecords(ctx context.Context, tx *sqlx.Tx, logger util.Logger, models ...*Record) error {
	if len(models) == 0 {
		return nil
	}
	q := database.SQLInsert(recordTableQuoted, recordColumnsQuoted, len(models), s.db.Placeholder())
	vals := lo.FlatMap(models, func(arg *Record, _ int) []any {
		return arg.ToData()
	})
	return s.db.Insert(ctx, q, tx, logger, vals...)
}
