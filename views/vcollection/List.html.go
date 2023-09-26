// Code generated by qtc from "List.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vcollection/List.html:2
package vcollection

//line views/vcollection/List.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vcollection/List.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcollection/List.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcollection/List.html:11
type List struct {
	layout.Basic
	Models      collection.Collections
	Params      filter.ParamSet
	SearchQuery string
}

//line views/vcollection/List.html:18
func (p *List) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/List.html:18
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vcollection/List.html:20
	components.StreamSearchForm(qw422016, "", "q", "Search collections", p.SearchQuery, ps)
//line views/vcollection/List.html:20
	qw422016.N().S(`</div>
    <h3>`)
//line views/vcollection/List.html:21
	components.StreamSVGRefIcon(qw422016, `archive`, ps)
//line views/vcollection/List.html:21
	qw422016.E().S(ps.Title)
//line views/vcollection/List.html:21
	qw422016.N().S(` <a href="/collection/new"><button>New</button></a></h3>
    <div class="clear"></div>
`)
//line views/vcollection/List.html:23
	if p.SearchQuery != "" {
//line views/vcollection/List.html:23
		qw422016.N().S(`    <em>Search results for [`)
//line views/vcollection/List.html:24
		qw422016.E().S(p.SearchQuery)
//line views/vcollection/List.html:24
		qw422016.N().S(`]</em>
`)
//line views/vcollection/List.html:25
	}
//line views/vcollection/List.html:26
	if len(p.Models) == 0 {
//line views/vcollection/List.html:26
		qw422016.N().S(`    <div class="mt"><em>No collections available</em></div>
`)
//line views/vcollection/List.html:28
	} else {
//line views/vcollection/List.html:28
		qw422016.N().S(`    <div class="overflow clear">
      `)
//line views/vcollection/List.html:30
		StreamTable(qw422016, p.Models, p.Params, as, ps)
//line views/vcollection/List.html:30
		qw422016.N().S(`
    </div>
`)
//line views/vcollection/List.html:32
	}
//line views/vcollection/List.html:32
	qw422016.N().S(`  </div>
`)
//line views/vcollection/List.html:34
}

//line views/vcollection/List.html:34
func (p *List) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/List.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcollection/List.html:34
	p.StreamBody(qw422016, as, ps)
//line views/vcollection/List.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/vcollection/List.html:34
}

//line views/vcollection/List.html:34
func (p *List) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcollection/List.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcollection/List.html:34
	p.WriteBody(qb422016, as, ps)
//line views/vcollection/List.html:34
	qs422016 := string(qb422016.B)
//line views/vcollection/List.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcollection/List.html:34
	return qs422016
//line views/vcollection/List.html:34
}