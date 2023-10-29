// Package ccollection - Content managed by Project Forge, see [projectforge.md] for details.
package ccollection

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/vcollection/vitem"
)

func ItemList(rc *fasthttp.RequestCtx) {
	controller.Act("item.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("item", nil, ps.Logger).Sanitize("item")
		var ret item.Items
		var err error
		if q == "" {
			ret, err = as.Services.Item.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Item.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Items", ret)
		collectionIDsByCollectionID := lo.Map(ret, func(x *item.Item, _ int) uuid.UUID {
			return x.CollectionID
		})
		collectionsByCollectionID, err := as.Services.Collection.GetMultiple(ps.Context, nil, nil, ps.Logger, collectionIDsByCollectionID...)
		if err != nil {
			return "", err
		}
		page := &vitem.List{Models: ret, CollectionsByCollectionID: collectionsByCollectionID, Params: ps.Params, SearchQuery: q}
		return controller.Render(rc, as, page, ps, "collection", "item")
	})
}

//nolint:lll
func ItemDetail(rc *fasthttp.RequestCtx) {
	controller.Act("item.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Item)", ret)

		collectionByCollectionID, _ := as.Services.Collection.Get(ps.Context, nil, ret.CollectionID, ps.Logger)

		return controller.Render(rc, as, &vitem.Detail{Model: ret, CollectionByCollectionID: collectionByCollectionID}, ps, "collection", "item", ret.TitleString()+"**file")
	})
}

func ItemCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("item.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &item.Item{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = item.Random()
		}
		ps.SetTitleAndData("Create [Item]", ret)
		ps.Data = ret
		return controller.Render(rc, as, &vitem.Edit{Model: ret, IsNew: true}, ps, "collection", "item", "Create")
	})
}

func ItemRandom(rc *fasthttp.RequestCtx) {
	controller.Act("item.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Item.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Item")
		}
		return ret.WebPath(), nil
	})
}

func ItemCreate(rc *fasthttp.RequestCtx) {
	controller.Act("item.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Item from form")
		}
		err = as.Services.Item.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Item")
		}
		msg := fmt.Sprintf("Item [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func ItemEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("item.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(rc, as, &vitem.Edit{Model: ret}, ps, "collection", "item", ret.String())
	})
}

func ItemEdit(rc *fasthttp.RequestCtx) {
	controller.Act("item.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := itemFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Item from form")
		}
		frm.ID = ret.ID
		err = as.Services.Item.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Item [%s]", frm.String())
		}
		msg := fmt.Sprintf("Item [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func ItemDelete(rc *fasthttp.RequestCtx) {
	controller.Act("item.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Item.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete item [%s]", ret.String())
		}
		msg := fmt.Sprintf("Item [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/collection/item", rc, ps)
	})
}

func itemFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*item.Item, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Item.Get(ps.Context, nil, idArg, ps.Logger)
}

func itemFromForm(rc *fasthttp.RequestCtx, setPK bool) (*item.Item, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return item.FromMap(frm, setPK)
}
