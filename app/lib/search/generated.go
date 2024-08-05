package search

import (
	"context"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/search/result"
	"github.com/kyleu/todoforge/app/util"
)

func generatedSearch() []Provider {
	collectionFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		prm := params.PS.Sanitized("collection", logger).WithLimit(5)
		return as.Services.Collection.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	itemFunc := func(ctx context.Context, params *Params, as *app.State, page *cutil.PageState, logger util.Logger) (result.Results, error) {
		prm := params.PS.Sanitized("item", logger).WithLimit(5)
		return as.Services.Item.SearchEntries(ctx, params.Q, nil, prm, logger)
	}
	return []Provider{collectionFunc, itemFunc}
}
