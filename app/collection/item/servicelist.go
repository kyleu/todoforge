package item

import (
	"context"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/lib/search/result"
	"github.com/kyleu/todoforge/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Items, error) {
	params = filters(params)
	wc := ""
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get items")
	}
	return ret.ToItems(), nil
}

func (s *Service) ListSQL(ctx context.Context, tx *sqlx.Tx, sql string, logger util.Logger, values ...any) (Items, error) {
	ret := rows{}
	err := s.db.Select(ctx, &ret, sql, tx, logger, values...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get items using custom SQL")
	}
	return ret.ToItems(), nil
}

func (s *Service) ListWhere(ctx context.Context, tx *sqlx.Tx, where string, params *filter.Params, logger util.Logger, values ...any) (Items, error) {
	params = filters(params)
	sql := database.SQLSelect(columnsString, tableQuoted, where, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	return s.ListSQL(ctx, tx, sql, logger, values...)
}

func (s *Service) Count(ctx context.Context, tx *sqlx.Tx, whereClause string, logger util.Logger, args ...any) (int, error) {
	if strings.Contains(whereClause, "'") || strings.Contains(whereClause, ";") {
		return 0, errors.Errorf("invalid where clause [%s]", whereClause)
	}
	q := database.SQLSelectSimple("count(*) as x", tableQuoted, s.db.Type, whereClause)
	ret, err := s.db.SingleInt(ctx, q, tx, logger, args...)
	if err != nil {
		return 0, errors.Wrap(err, "unable to get count of items")
	}
	return int(ret), nil
}

const searchClause = `(lower(id) like $1 or lower("name") like $1)`

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Items, error) {
	params = filters(params)
	var wc string
	var vals []any
	var err error
	if strings.Contains(query, ":") {
		wc, vals, err = database.QueryFieldDescs(ItemFieldDescs, query, 0)
	} else {
		wc = searchClause
		vals = []any{"%" + strings.ToLower(query) + "%"}
	}
	if err != nil {
		return nil, err
	}
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err = s.db.Select(ctx, &ret, q, tx, logger, vals...)
	if err != nil {
		return nil, err
	}
	return ret.ToItems(), nil
}

func (s *Service) SearchEntries(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (result.Results, error) {
	ret, err := s.Search(ctx, query, tx, params, logger)
	if err != nil {
		return nil, err
	}
	return lo.Map(ret, func(m *Item, _ int) *result.Result {
		return result.NewResult("Item", m.String(), m.WebPath(), m.TitleString(), "file", m, m, query)
	}), nil
}
