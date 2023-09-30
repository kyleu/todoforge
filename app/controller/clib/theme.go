// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/csession"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/theme"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/vtheme"
)

func ThemeList(rc *fasthttp.RequestCtx) {
	controller.Act("theme.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.Title = "Themes"
		th := as.Themes.All(ps.Logger)
		ps.Data = th
		return controller.Render(rc, as, &vtheme.List{Themes: th}, ps, "Themes||/theme")
	})
}

func ThemeEdit(rc *fasthttp.RequestCtx) {
	controller.Act("theme.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		if key == theme.Default.Key {
			return controller.FlashAndRedir(false, "Unable to edit default theme", "/theme", rc, ps)
		}
		var t *theme.Theme
		if key == theme.KeyNew {
			t = theme.Default.Clone(key)
		} else {
			t = as.Themes.Get(key, ps.Logger)
			if t == nil {
				if pal := string(rc.URI().QueryArgs().Peek("palette")); pal != "" {
					themes, err := theme.PaletteThemes(pal)
					if err != nil {
						return "", err
					}
					t = themes.Get(key)
				}
			}
		}
		if t == nil {
			return "", errors.Errorf("invalid theme [%s]", key)
		}
		ps.Title = "Edit theme [" + t.Key + "]"
		ps.Data = t
		page := &vtheme.Edit{Theme: t, Icon: "app", Exists: as.Themes.FileExists(t.Key)}
		return controller.Render(rc, as, page, ps, "Themes||/theme", t.Key)
	})
}

func ThemeSave(rc *fasthttp.RequestCtx) {
	controller.Act("theme.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}

		orig := as.Themes.Get(key, ps.Logger)

		newKey, err := frm.GetString("key", false)
		if err != nil {
			return "", err
		}
		if newKey == theme.KeyNew {
			newKey = util.RandomString(12)
		}

		l := orig.Light.Clone().ApplyMap(frm, "light-")
		d := orig.Dark.Clone().ApplyMap(frm, "dark-")

		t := &theme.Theme{Key: newKey, Light: l, Dark: d}

		err = as.Themes.Save(t, frm.GetStringOpt("originalKey"), ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to save theme")
		}

		ps.Profile.Theme = newKey
		err = csession.SaveProfile(ps.Profile, rc, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		return controller.ReturnToReferrer("saved changes to theme ["+newKey+"]", "/theme", rc, ps)
	})
}

func ThemeRemove(rc *fasthttp.RequestCtx) {
	controller.Act("theme.remove", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.RCRequiredString(rc, "key", false)
		if err != nil {
			return "", err
		}
		err = as.Themes.Remove(key, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to remove theme")
		}
		return controller.ReturnToReferrer("removed theme ["+key+"]", "/theme", rc, ps)
	})
}
