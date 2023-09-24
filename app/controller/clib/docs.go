// Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/doc"
	"github.com/kyleu/todoforge/views/vdoc"
)

func Docs(rc *fasthttp.RequestCtx) {
	controller.Act("docs", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		pth, _ := cutil.RCRequiredString(rc, "path", false)
		if pth == "" {
			return "", errors.New("invalid path")
		}

		bc := []string{"docs"}
		bc = append(bc, util.StringSplitAndTrim(pth, "/")...)

		title, x, err := doc.HTML("doc:"+pth, pth+".md", func(s string) (string, string, error) {
			return cutil.FormatCleanMarkup(s, "file")
		})
		if err != nil {
			return "", errors.Wrapf(err, "unable to load documentation from [%s]", pth)
		}
		ps.Title = title
		ps.Data, _ = doc.Content(pth + ".md")
		return controller.Render(rc, as, &vdoc.MarkdownPage{Title: pth, HTML: x}, ps, bc...)
	})
}
