// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"net/url"

	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/csession"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/theme"
	"github.com/kyleu/todoforge/views/vprofile"
)

func Profile(rc *fasthttp.RequestCtx) {
	controller.Act("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}

func ProfileSite(rc *fasthttp.RequestCtx) {
	controller.ActSite("profile", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		return profileAction(rc, as, ps)
	})
}

func profileAction(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (string, error) {
	ps.SetTitleAndData("Profile", ps.Profile)
	thm := as.Themes.Get(ps.Profile.Theme, ps.Logger)

	redir := "/"
	ref := string(rc.Request.Header.Peek("Referer"))
	if ref != "" {
		u, err := url.Parse(ref)
		if err == nil && u != nil && u.Path != cutil.DefaultProfilePath {
			redir = u.Path
		}
	}
	ps.DefaultNavIcon = "profile"
	page := &vprofile.Profile{Profile: ps.Profile, Theme: thm, Referrer: redir}
	return controller.Render(rc, as, page, ps, "Profile")
}

func ProfileSave(rc *fasthttp.RequestCtx) {
	controller.Act("profile.save", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		frm, err := cutil.ParseForm(rc)
		if err != nil {
			return "", err
		}

		n := ps.Profile.Clone()

		referrerDefault := frm.GetStringOpt("referrer")
		if referrerDefault == "" {
			referrerDefault = cutil.DefaultProfilePath
		}

		n.Name = frm.GetStringOpt("name")
		n.Mode = frm.GetStringOpt("mode")
		n.Theme = frm.GetStringOpt("theme")
		if n.Theme == theme.Default.Key {
			n.Theme = ""
		}

		err = csession.SaveProfile(n, rc, ps.Session, ps.Logger)
		if err != nil {
			return "", err
		}

		return controller.ReturnToReferrer("Saved profile", referrerDefault, rc, ps)
	})
}
