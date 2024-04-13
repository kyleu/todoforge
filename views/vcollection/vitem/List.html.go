// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vcollection/vitem/List.html:2
package vitem

//line views/vcollection/vitem/List.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/components/edit"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vcollection/vitem/List.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcollection/vitem/List.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcollection/vitem/List.html:13
type List struct {
	layout.Basic
	Models                    item.Items
	CollectionsByCollectionID collection.Collections
	Params                    filter.ParamSet
	SearchQuery               string
}

//line views/vcollection/vitem/List.html:21
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/vitem/List.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vcollection/vitem/List.html:23
	edit.StreamSearchForm(qw422016, "", "q", "Search Items", p.SearchQuery, ps)
//line views/vcollection/vitem/List.html:23
	qw422016.N().S(`</div>
    <div class="right mrs large-buttons">
`)
//line views/vcollection/vitem/List.html:25
	if len(p.Models) > 0 {
//line views/vcollection/vitem/List.html:25
		qw422016.N().S(`<a href="/collection/item/_random"><button>Random</button></a>`)
//line views/vcollection/vitem/List.html:25
	}
//line views/vcollection/vitem/List.html:25
	qw422016.N().S(`      <a href="/collection/item/_new"><button>New</button></a>
    </div>
    <h3>`)
//line views/vcollection/vitem/List.html:28
	components.StreamSVGRefIcon(qw422016, `file`, ps)
//line views/vcollection/vitem/List.html:28
	qw422016.E().S(ps.Title)
//line views/vcollection/vitem/List.html:28
	qw422016.N().S(`</h3>
    <div class="clear"></div>
`)
//line views/vcollection/vitem/List.html:30
	if p.SearchQuery != "" {
//line views/vcollection/vitem/List.html:30
		qw422016.N().S(`    <hr />
    <em>Search results for [`)
//line views/vcollection/vitem/List.html:32
		qw422016.E().S(p.SearchQuery)
//line views/vcollection/vitem/List.html:32
		qw422016.N().S(`]</em> (<a href="?">clear</a>)
`)
//line views/vcollection/vitem/List.html:33
	}
//line views/vcollection/vitem/List.html:34
	if len(p.Models) == 0 {
//line views/vcollection/vitem/List.html:34
		qw422016.N().S(`    <div class="mt"><em>No items available</em></div>
`)
//line views/vcollection/vitem/List.html:36
	} else {
//line views/vcollection/vitem/List.html:36
		qw422016.N().S(`    <div class="mt">
      `)
//line views/vcollection/vitem/List.html:38
		StreamTable(qw422016, p.Models, p.CollectionsByCollectionID, p.Params, as, ps)
//line views/vcollection/vitem/List.html:38
		qw422016.N().S(`
    </div>
`)
//line views/vcollection/vitem/List.html:40
	}
//line views/vcollection/vitem/List.html:40
	qw422016.N().S(`  </div>
`)
//line views/vcollection/vitem/List.html:42
}

//line views/vcollection/vitem/List.html:42
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/vitem/List.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcollection/vitem/List.html:42
	p.StreamBody(qw422016, as, ps)
//line views/vcollection/vitem/List.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/vcollection/vitem/List.html:42
}

//line views/vcollection/vitem/List.html:42
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcollection/vitem/List.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcollection/vitem/List.html:42
	p.WriteBody(qb422016, as, ps)
//line views/vcollection/vitem/List.html:42
	qs422016 := string(qb422016.B)
//line views/vcollection/vitem/List.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcollection/vitem/List.html:42
	return qs422016
//line views/vcollection/vitem/List.html:42
}
