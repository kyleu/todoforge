package item

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/util"
)

func (s *Service) Get(ctx context.Context, tx *sqlx.Tx, id uuid.UUID, logger util.Logger) (*Item, error) {
	wc := defaultWC(0)
	ret := &row{}
	q := database.SQLSelectSimple(columnsString, tableQuoted, s.db.Type, wc)
	err := s.db.Get(ctx, ret, q, tx, logger, id)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get item by id [%v]", id)
	}
	return ret.ToItem(), nil
}

func (s *Service) GetMultiple(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, ids ...uuid.UUID) (Items, error) {
	if len(ids) == 0 {
		return Items{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("id", len(ids), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(ids)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Items for [%d] ids", len(ids))
	}
	return ret.ToItems(), nil
}

func (s *Service) GetByCollectionID(ctx context.Context, tx *sqlx.Tx, collectionID uuid.UUID, params *filter.Params, logger util.Logger) (Items, error) {
	params = filters(params)
	wc := "\"collection_id\" = $1"
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, collectionID)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Items by collectionID [%v]", collectionID)
	}
	return ret.ToItems(), nil
}

func (s *Service) GetByCollectionIDs(ctx context.Context, tx *sqlx.Tx, params *filter.Params, logger util.Logger, collectionIDs ...uuid.UUID) (Items, error) {
	if len(collectionIDs) == 0 {
		return Items{}, nil
	}
	params = filters(params)
	wc := database.SQLInClause("collection_id", len(collectionIDs), 0, s.db.Type)
	q := database.SQLSelect(columnsString, tableQuoted, wc, params.OrderByString(), params.Limit, params.Offset, s.db.Type)
	ret := rows{}
	err := s.db.Select(ctx, &ret, q, tx, logger, lo.ToAnySlice(collectionIDs)...)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to get Items for [%d] collectionIDs", len(collectionIDs))
	}
	return ret.ToItems(), nil
}

func (s *Service) Random(ctx context.Context, tx *sqlx.Tx, logger util.Logger) (*Item, error) {
	ret := &row{}
	q := database.SQLSelect(columnsString, tableQuoted, "", "random()", 1, 0, s.db.Type)
	err := s.db.Get(ctx, ret, q, tx, logger)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get random items")
	}
	return ret.ToItem(), nil
}
