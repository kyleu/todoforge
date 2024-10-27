package ccollection

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/vcollection/vitem"
)

func ItemList(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.list", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		q := strings.TrimSpace(r.URL.Query().Get("q"))
		prms := ps.Params.Sanitized("item", ps.Logger)
		var ret item.Items
		var err error
		if q == "" {
			ret, err = as.Services.Item.List(ps.Context, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
		} else {
			ret, err = as.Services.Item.Search(ps.Context, q, nil, prms, ps.Logger)
			if err != nil {
				return "", err
			}
			if len(ret) == 1 {
				return controller.FlashAndRedir(true, "single result found", ret[0].WebPath(), ps)
			}
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
		return controller.Render(r, as, page, ps, "collection", "item")
	})
}

//nolint:lll
func ItemDetail(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Item)", ret)

		collectionByCollectionID, _ := as.Services.Collection.Get(ps.Context, nil, ret.CollectionID, ps.Logger)

		return controller.Render(r, as, &vitem.Detail{Model: ret, CollectionByCollectionID: collectionByCollectionID}, ps, "collection", "item", ret.TitleString()+"**file")
	})
}

func ItemCreateForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.create.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &item.Item{}
		if r.URL.Query().Get("prototype") == util.KeyRandom {
			ret = item.RandomItem()
			randomCollection, err := as.Services.Collection.Random(ps.Context, nil, ps.Logger)
			if err == nil && randomCollection != nil {
				ret.CollectionID = randomCollection.ID
			}
		}
		ps.SetTitleAndData("Create [Item]", ret)
		ps.Data = ret
		return controller.Render(r, as, &vitem.Edit{Model: ret, IsNew: true}, ps, "collection", "item", "Create")
	})
}

func ItemRandom(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.random", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.Item.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random Item")
		}
		return ret.WebPath(), nil
	})
}

func ItemCreate(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.create", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromForm(r, ps.RequestBody, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Item from form")
		}
		err = as.Services.Item.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Item")
		}
		msg := fmt.Sprintf("Item [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), ps)
	})
}

func ItemEditForm(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.edit.form", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(r, as, &vitem.Edit{Model: ret}, ps, "collection", "item", ret.String())
	})
}

func ItemEdit(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.edit", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := itemFromForm(r, ps.RequestBody, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Item from form")
		}
		frm.ID = ret.ID
		err = as.Services.Item.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Item [%s]", frm.String())
		}
		msg := fmt.Sprintf("Item [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), ps)
	})
}

func ItemDelete(w http.ResponseWriter, r *http.Request) {
	controller.Act("item.delete", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := itemFromPath(r, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Item.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete item [%s]", ret.String())
		}
		msg := fmt.Sprintf("Item [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/collection/item", ps)
	})
}

func itemFromPath(r *http.Request, as *app.State, ps *cutil.PageState) (*item.Item, error) {
	idArgStr, err := cutil.PathString(r, "id", false)
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

func itemFromForm(r *http.Request, b []byte, setPK bool) (*item.Item, error) {
	frm, err := cutil.ParseForm(r, b)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	ret, _, err := item.ItemFromMap(frm, setPK)
	return ret, err
}
