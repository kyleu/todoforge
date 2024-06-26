package search

import (
	"context"

	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/search/result"
	"github.com/kyleu/todoforge/app/util"
)

func generatedSearch() []Provider {
	collectionFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		prm := params.PS.Sanitized("collection", logger).WithLimit(5)
		models, err := as.Services.Collection.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *collection.Collection, _ int) *result.Result {
			return result.NewResult("Collection", m.String(), m.WebPath(), m.TitleString(), "archive", m, m, params.Q)
		}), nil
	}
	itemFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		prm := params.PS.Sanitized("item", logger).WithLimit(5)
		models, err := as.Services.Item.Search(ctx, params.Q, nil, prm, logger)
		if err != nil {
			return nil, err
		}
		return lo.Map(models, func(m *item.Item, _ int) *result.Result {
			return result.NewResult("Item", m.String(), m.WebPath(), m.TitleString(), "file", m, m, params.Q)
		}), nil
	}
	return []Provider{collectionFunc, itemFunc}
}
