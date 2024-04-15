// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/vcollection"
)

func CollectionList(w http.ResponseWriter, r *http.Request) {
	Act("collection.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("collection", ps.Logger)
		var ret collection.Collections
		var err error
		if q == "" {
			ret, err = as.Services.Collection.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Collection.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
		}
		ps.SetTitleAndData("Collections", ret)
		page := &vcollection.List{Models: ret, Params: ps.Params, SearchQuery: q}
		return Render(r, as, page, ps, "collection")
	})
}

func CollectionDetail(w http.ResponseWriter, r *http.Request) {
	Act("collection.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Collection)", ret)

		relItemsByCollectionIDPrms := ps.Params.Sanitized("item", ps.Logger)
		relItemsByCollectionID, err := as.Services.Item.GetByCollectionID(ps.Context, nil, ret.ID, relItemsByCollectionIDPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child items")
		}
		return Render(r, as, &vcollection.Detail{
			Model:  ret,
			Params: ps.Params,

			RelItemsByCollectionID: relItemsByCollectionID,
		}, ps, "collection", ret.TitleString()+"**archive")
	})
}

func CollectionCreateForm(w http.ResponseWriter, r *http.Request) {
	Act("collection.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &collection.Collection{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = collection.Random()
		}
		ps.SetTitleAndData("Create [Collection]", ret)
		ps.Data = ret
		return Render(r, as, &vcollection.Edit{Model: ret, IsNew: true}, ps, "collection", "Create")
	})
}

func CollectionRandom(w http.ResponseWriter, r *http.Request) {
	Act("collection.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Collection.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Collection")
		}
		return ret.WebPath(), nil
	})
}

func CollectionCreate(w http.ResponseWriter, r *http.Request) {
	Act("collection.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Collection from form")
		}
		err = as.Services.Collection.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Collection")
		}
		msg := fmt.Sprintf("Collection [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func CollectionEditForm(w http.ResponseWriter, r *http.Request) {
	Act("collection.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return Render(r, as, &vcollection.Edit{Model: ret}, ps, "collection", ret.String())
	})
}

func CollectionEdit(w http.ResponseWriter, r *http.Request) {
	Act("collection.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := collectionFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Collection from form")
		}
		frm.ID = ret.ID
		err = as.Services.Collection.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Collection [%s]", frm.String())
		}
		msg := fmt.Sprintf("Collection [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func CollectionDelete(w http.ResponseWriter, r *http.Request) {
	Act("collection.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := collectionFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Collection.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete collection [%s]", ret.String())
		}
		msg := fmt.Sprintf("Collection [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/collection", ps)
	})
}

func collectionFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*collection.Collection, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func collectionFromForm(r *http.Request, b []byte, setPK bool) (*collection.Collection, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return collection.FromMap(frm, setPK)
}
