// Package database - Content managed by Project Forge, see [projectforge.md] for details.
package database

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/todoforge/app/util"
)

func (s *Service) Explain(ctx context.Context, q string, values []any, _ util.Logger) ([]util.ValueMap, error) {
	q = strings.TrimSpace(q)
	explainPrefix := "explain "
	if !strings.HasPrefix(q, explainPrefix) {
		if s.Type.Key == TypeSQLite.Key {
			explainPrefix += "query plan "
		}
		q = explainPrefix + q
	}
	res, err := s.db.QueryxContext(ctx, q, values...)
	if err != nil {
		return nil, errors.Wrap(err, "invalid explain result")
	}
	defer func() { _ = res.Close() }()
	var ret []util.ValueMap
	for res.Next() {
		x := util.ValueMap{}
		err = res.MapScan(x)
		if err != nil {
			return nil, errors.Wrap(err, "can't read results")
		}
		ret = append(ret, x)
	}

	return ret, nil
}
