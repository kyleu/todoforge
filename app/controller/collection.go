// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/vcollection"
)

func CollectionList(rc *fasthttp.RequestCtx) {
	Act("collection.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(string(rc.URI().QueryArgs().Peek("q")))
		prms := ps.Params.Get("collection", nil, ps.Logger).Sanitize("collection")
		var ret collection.Collections
		var err error
		if q == "" {
			ret, err = as.Services.Collection.List(ps.Context, nil, prms, ps.Logger)
		} else {
			ret, err = as.Services.Collection.Search(ps.Context, q, nil, prms, ps.Logger)
		}
		if err != nil {
			return "", err
		}
		ps.Title = "Collections"
		ps.Data = ret
		page := &vcollection.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(rc, as, page, ps, "collection")
	})
}

func CollectionDetail(rc *fasthttp.RequestCtx) {
	Act("collection.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Collection)"
		ps.Data = ret

		relItemsByCollectionIDPrms := ps.Params.Get("item", nil, ps.Logger).Sanitize("item")
		relItemsByCollectionID, err := as.Services.Item.GetByCollectionID(ps.Context, nil, ret.ID, relItemsByCollectionIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child items")
		}
		return Render(rc, as, &vcollection.Detail{
			Model:  ret,
			Params: ps.Params,

			RelItemsByCollectionID: relItemsByCollectionID,
		}, ps, "collection", ret.String())
	})
}

func CollectionCreateForm(rc *fasthttp.RequestCtx) {
	Act("collection.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &collection.Collection{}
		ps.Title = "Create [Collection]"
		ps.Data = ret
		return Render(rc, as, &vcollection.Edit{Model: ret, IsNew: true}, ps, "collection", "Create")
	})
}

func CollectionCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("collection.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := collection.Random()
		ps.Title = "Create Random Collection"
		ps.Data = ret
		return Render(rc, as, &vcollection.Edit{Model: ret, IsNew: true}, ps, "collection", "Create")
	})
}

func CollectionCreate(rc *fasthttp.RequestCtx) {
	Act("collection.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Collection from form")
		}
		err = as.Services.Collection.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Collection")
		}
		msg := fmt.Sprintf("Collection [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func CollectionEditForm(rc *fasthttp.RequestCtx) {
	Act("collection.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vcollection.Edit{Model: ret}, ps, "collection", ret.String())
	})
}

func CollectionEdit(rc *fasthttp.RequestCtx) {
	Act("collection.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := collectionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Collection from form")
		}
		frm.ID = ret.ID
		err = as.Services.Collection.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Collection [%s]", frm.String())
		}
		msg := fmt.Sprintf("Collection [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func CollectionDelete(rc *fasthttp.RequestCtx) {
	Act("collection.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Collection.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete collection [%s]", ret.String())
		}
		msg := fmt.Sprintf("Collection [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/collection", rc, ps)
	})
}

func collectionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*collection.Collection, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Collection.Get(ps.Context, nil, idArg, ps.Logger)
}

func collectionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*collection.Collection, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return collection.FromMap(frm, setPK)
}
