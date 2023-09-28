// Code generated by qtc from "Sitemap.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Sitemap.html:2
package vadmin

//line views/vadmin/Sitemap.html:2
import (
	"slices"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cmenu"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/menu"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
	"github.com/kyleu/todoforge/views/vutil"
)

//line views/vadmin/Sitemap.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Sitemap.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Sitemap.html:14
type Sitemap struct {
	layout.Basic
}

//line views/vadmin/Sitemap.html:18
func (p *Sitemap) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sitemap.html:18
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/Sitemap.html:20
	components.StreamSVGRefIcon(qw422016, `star`, ps)
//line views/vadmin/Sitemap.html:20
	qw422016.N().S(`Sitemap</h3>
    <div class="mt">
      `)
//line views/vadmin/Sitemap.html:22
	StreamSitemapDetail(qw422016, ps.Menu, 1, ps)
//line views/vadmin/Sitemap.html:22
	qw422016.N().S(`
    </div>
  </div>
`)
//line views/vadmin/Sitemap.html:25
}

//line views/vadmin/Sitemap.html:25
func (p *Sitemap) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Sitemap.html:25
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sitemap.html:25
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Sitemap.html:25
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sitemap.html:25
}

//line views/vadmin/Sitemap.html:25
func (p *Sitemap) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Sitemap.html:25
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sitemap.html:25
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Sitemap.html:25
	qs422016 := string(qb422016.B)
//line views/vadmin/Sitemap.html:25
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sitemap.html:25
	return qs422016
//line views/vadmin/Sitemap.html:25
}

//line views/vadmin/Sitemap.html:27
func StreamSitemapDetail(qw422016 *qt422016.Writer, m menu.Items, indent int, ps *cutil.PageState) {
//line views/vadmin/Sitemap.html:28
	vutil.StreamIndent(qw422016, true, 1)
//line views/vadmin/Sitemap.html:28
	qw422016.N().S(`<div class="mt">`)
//line views/vadmin/Sitemap.html:30
	vutil.StreamIndent(qw422016, true, 2)
//line views/vadmin/Sitemap.html:30
	qw422016.N().S(`<ul class="level-0">`)
//line views/vadmin/Sitemap.html:32
	for _, i := range m {
//line views/vadmin/Sitemap.html:33
		if i.Key != "" {
//line views/vadmin/Sitemap.html:34
			streamsitemapItemDetail(qw422016, i, []string{}, ps.Breadcrumbs, 3, ps)
//line views/vadmin/Sitemap.html:35
		}
//line views/vadmin/Sitemap.html:36
	}
//line views/vadmin/Sitemap.html:37
	vutil.StreamIndent(qw422016, true, 2)
//line views/vadmin/Sitemap.html:37
	qw422016.N().S(`</ul>`)
//line views/vadmin/Sitemap.html:39
	vutil.StreamIndent(qw422016, true, 1)
//line views/vadmin/Sitemap.html:39
	qw422016.N().S(`</div>`)
//line views/vadmin/Sitemap.html:41
}

//line views/vadmin/Sitemap.html:41
func WriteSitemapDetail(qq422016 qtio422016.Writer, m menu.Items, indent int, ps *cutil.PageState) {
//line views/vadmin/Sitemap.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sitemap.html:41
	StreamSitemapDetail(qw422016, m, indent, ps)
//line views/vadmin/Sitemap.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sitemap.html:41
}

//line views/vadmin/Sitemap.html:41
func SitemapDetail(m menu.Items, indent int, ps *cutil.PageState) string {
//line views/vadmin/Sitemap.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sitemap.html:41
	WriteSitemapDetail(qb422016, m, indent, ps)
//line views/vadmin/Sitemap.html:41
	qs422016 := string(qb422016.B)
//line views/vadmin/Sitemap.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sitemap.html:41
	return qs422016
//line views/vadmin/Sitemap.html:41
}

//line views/vadmin/Sitemap.html:43
func streamsitemapItemDetail(qw422016 *qt422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/vadmin/Sitemap.html:44
	vutil.StreamIndent(qw422016, true, indent)
//line views/vadmin/Sitemap.html:44
	qw422016.N().S(`<li><div class="mts">`)
//line views/vadmin/Sitemap.html:47
	vutil.StreamIndent(qw422016, true, indent+1)
//line views/vadmin/Sitemap.html:47
	qw422016.N().S(`<a href="`)
//line views/vadmin/Sitemap.html:48
	qw422016.E().S(i.Route)
//line views/vadmin/Sitemap.html:48
	qw422016.N().S(`" title="`)
//line views/vadmin/Sitemap.html:48
	qw422016.E().S(i.Desc())
//line views/vadmin/Sitemap.html:48
	qw422016.N().S(`">`)
//line views/vadmin/Sitemap.html:49
	if i.Icon != "" {
//line views/vadmin/Sitemap.html:50
		components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/vadmin/Sitemap.html:51
	}
//line views/vadmin/Sitemap.html:52
	qw422016.E().S(i.Title)
//line views/vadmin/Sitemap.html:52
	qw422016.N().S(`</a><div><em>`)
//line views/vadmin/Sitemap.html:54
	qw422016.E().S(i.Desc())
//line views/vadmin/Sitemap.html:54
	qw422016.N().S(`</em></div>`)
//line views/vadmin/Sitemap.html:55
	if len(i.Children) > 0 {
//line views/vadmin/Sitemap.html:55
		qw422016.N().S(`<ul class="level-`)
//line views/vadmin/Sitemap.html:56
		qw422016.N().D(len(path))
//line views/vadmin/Sitemap.html:56
		qw422016.N().S(`">`)
//line views/vadmin/Sitemap.html:57
		for _, kid := range i.Children {
//line views/vadmin/Sitemap.html:58
			if kid.Key != "" {
//line views/vadmin/Sitemap.html:59
				streamsitemapItemDetail(qw422016, kid, append(slices.Clone(path), i.Key), breadcrumbs, indent+2, ps)
//line views/vadmin/Sitemap.html:60
			}
//line views/vadmin/Sitemap.html:61
		}
//line views/vadmin/Sitemap.html:61
		qw422016.N().S(`</ul>`)
//line views/vadmin/Sitemap.html:63
	}
//line views/vadmin/Sitemap.html:63
	qw422016.N().S(`</div>`)
//line views/vadmin/Sitemap.html:65
	vutil.StreamIndent(qw422016, true, indent)
//line views/vadmin/Sitemap.html:65
	qw422016.N().S(`</li>`)
//line views/vadmin/Sitemap.html:67
}

//line views/vadmin/Sitemap.html:67
func writesitemapItemDetail(qq422016 qtio422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/vadmin/Sitemap.html:67
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Sitemap.html:67
	streamsitemapItemDetail(qw422016, i, path, breadcrumbs, indent, ps)
//line views/vadmin/Sitemap.html:67
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Sitemap.html:67
}

//line views/vadmin/Sitemap.html:67
func sitemapItemDetail(i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) string {
//line views/vadmin/Sitemap.html:67
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Sitemap.html:67
	writesitemapItemDetail(qb422016, i, path, breadcrumbs, indent, ps)
//line views/vadmin/Sitemap.html:67
	qs422016 := string(qb422016.B)
//line views/vadmin/Sitemap.html:67
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Sitemap.html:67
	return qs422016
//line views/vadmin/Sitemap.html:67
}
