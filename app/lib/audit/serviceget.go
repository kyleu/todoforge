// Content managed by Project Forge, see [projectforge.md] for details.
package audit

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/util"
)

func (s *Service) List(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Audits, error) {
	params = filters(params)
	q := database.SQLSelect(columnsString, tableQuoted, "", params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get audits")
	}
	return ret.ToAudits(), nil
}

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Audit, error) {
	wc := defaultWC
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Placeholder(), wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get audit by id [%v]", id)
	}
	return ret.ToAudit(), nil
}

const searchClause = `(
  lower(id) like $1 or lower(app) like $1 or lower(act) like $1 or
  lower(client) like $1 or lower(server) like $1 or lower(user) like $1 or
  lower(metadata::text) like $1 or lower(message) like $1
)`

func (s *Service) Search(ctx context.Context, query string, tx *sqlx.Tx, params *filter.Params, logger util.Logger) (Audits, error) {
	params = filters(params)
	wc := searchClause
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Placeholder())
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	return ret.ToAudits(), nil
}

func filters(orig *filter.Params) *filter.Params {
	return orig.Sanitize("audit", &filter.Ordering{Column: "started"})
}
