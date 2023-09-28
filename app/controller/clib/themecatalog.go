// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"
	"strings"

	"github.com/muesli/gamut"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/telemetry"
	"github.com/kyleu/todoforge/app/lib/theme"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views"
	"github.com/kyleu/todoforge/views/vtheme"
)

func ThemeColor(rc *fasthttp.RequestCtx) {
	controller.Act("theme.color", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		col, err := cutil.RCRequiredString(rc, "color", false)
		if err != nil {
			return "", err
		}
		col = strings.ToLower(col)
		if !strings.HasPrefix(col, "#") {
			col = "#" + col
		}
		ps.Title = fmt.Sprintf("[%s] Theme", col)
		x := theme.ColorTheme(col, gamut.Hex(col))
		ps.Data = x
		return controller.Render(rc, as, &vtheme.Edit{Theme: x, Icon: "app"}, ps, "Themes||/theme", col)
	})
}

func ThemeColorEdit(rc *fasthttp.RequestCtx) {
	controller.Act("theme.color.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		color := string(rc.URI().QueryArgs().Peek("color"))
		if color == "" {
			return "", errors.New("must provide color in query string")
		}
		if !strings.HasPrefix(color, "#") {
			return "", errors.New("provided color must be a hex string")
		}
		t := theme.ColorTheme(strings.TrimPrefix(color, "#"), gamut.Hex(color))
		ps.Data = t
		ps.Title = "Edit theme [" + t.Key + "]"
		page := &vtheme.Edit{Theme: t, Icon: "app", Exists: as.Themes.FileExists(t.Key)}
		return controller.Render(rc, as, page, ps, "Themes||/theme", t.Key)
	})
}

func ThemePalette(rc *fasthttp.RequestCtx) {
	controller.Act("theme.palette", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		pal, err := cutil.RCRequiredString(rc, "palette", false)
		if err != nil {
			return "", err
		}
		ps.Title = fmt.Sprintf("[%s] Themes", pal)
		_, span, _ := telemetry.StartSpan(ps.Context, "theme:load", ps.Logger)
		x, err := theme.PaletteThemes(pal)
		span.Complete()
		if err != nil {
			return "", err
		}
		if string(rc.URI().QueryArgs().Peek("t")) == "go" {
			ps.Data = strings.Join(lo.Map(x, func(t *theme.Theme, _ int) string {
				return t.ToGo()
			}), util.StringDefaultLinebreak)
			return controller.Render(rc, as, &views.Debug{}, ps, "Themes")
		}
		return controller.Render(rc, as, &vtheme.Add{Palette: pal, Themes: x}, ps, "Themes")
	})
}

func ThemePaletteEdit(rc *fasthttp.RequestCtx) {
	controller.Act("theme.palette.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		palette, err := cutil.RCRequiredString(rc, "palette", false)
		if err != nil {
			return "", err
		}
		key, err := cutil.RCRequiredString(rc, "theme", false)
		if err != nil {
			return "", err
		}
		if key == theme.Default.Key {
			return controller.FlashAndRedir(false, "Unable to edit default theme", "/theme", rc, ps)
		}
		themes, err := theme.PaletteThemes(palette)
		if err != nil {
			return "", err
		}
		t := themes.Get(key)
		if t == nil {
			return "", errors.Errorf("invalid theme [%s] for palette [%s]", key, palette)
		}
		ps.Data = t
		ps.Title = "Edit theme [" + t.Key + "]"
		return controller.Render(rc, as, &vtheme.Edit{Theme: t, Icon: "app"}, ps, "Themes||/theme", t.Key)
	})
}
