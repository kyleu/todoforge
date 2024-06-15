// Code generated by qtc from "Results.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsearch/Results.html:2
package vsearch

//line views/vsearch/Results.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/search"
	"github.com/kyleu/todoforge/app/lib/search/result"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vsearch/Results.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsearch/Results.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsearch/Results.html:12
type Results struct {
	layout.Basic
	Params  *search.Params
	Results result.Results
	Errors  []error
}

//line views/vsearch/Results.html:19
func (p *Results) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsearch/Results.html:19
	qw422016.N().S(`
  <div class="card">
    <h3 title="Search results for [`)
//line views/vsearch/Results.html:21
	qw422016.E().S(p.Params.Q)
//line views/vsearch/Results.html:21
	qw422016.N().S(`]">`)
//line views/vsearch/Results.html:21
	components.StreamSVGIcon(qw422016, "search", ps)
//line views/vsearch/Results.html:21
	qw422016.N().S(` Search Results</h3>
    <form class="mt expanded" action="`)
//line views/vsearch/Results.html:22
	qw422016.E().S(ps.SearchPath)
//line views/vsearch/Results.html:22
	qw422016.N().S(`">
      <input name="q" value="`)
//line views/vsearch/Results.html:23
	qw422016.E().S(p.Params.Q)
//line views/vsearch/Results.html:23
	qw422016.N().S(`" />
      <div class="mt"><button type="submit">Search Again</button></div>
    </form>
  </div>
`)
//line views/vsearch/Results.html:27
	if p.Params.Q != "" && len(p.Results) == 0 {
//line views/vsearch/Results.html:27
		qw422016.N().S(`  <div class="card">
    <h3>No results</h3>
  </div>
`)
//line views/vsearch/Results.html:31
	}
//line views/vsearch/Results.html:32
	for _, res := range p.Results {
//line views/vsearch/Results.html:32
		qw422016.N().S(`  `)
//line views/vsearch/Results.html:33
		StreamResult(qw422016, res, p.Params, as, ps)
//line views/vsearch/Results.html:33
		qw422016.N().S(`
`)
//line views/vsearch/Results.html:34
	}
//line views/vsearch/Results.html:34
	qw422016.N().S(`  `)
//line views/vsearch/Results.html:35
	if len(p.Errors) > 0 {
//line views/vsearch/Results.html:35
		qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsearch/Results.html:37
		components.StreamSVGIcon(qw422016, "error", ps)
//line views/vsearch/Results.html:37
		qw422016.N().S(` `)
//line views/vsearch/Results.html:37
		qw422016.E().S(util.StringPlural(len(p.Errors), "Error"))
//line views/vsearch/Results.html:37
		qw422016.N().S(`</h3>
    <ul class="mt">
`)
//line views/vsearch/Results.html:39
		for _, e := range p.Errors {
//line views/vsearch/Results.html:39
			qw422016.N().S(`      <li>`)
//line views/vsearch/Results.html:40
			qw422016.E().S(e.Error())
//line views/vsearch/Results.html:40
			qw422016.N().S(`</li>
`)
//line views/vsearch/Results.html:41
		}
//line views/vsearch/Results.html:41
		qw422016.N().S(`    </ul>
  </div>
  `)
//line views/vsearch/Results.html:44
	}
//line views/vsearch/Results.html:44
	qw422016.N().S(`
`)
//line views/vsearch/Results.html:45
}

//line views/vsearch/Results.html:45
func (p *Results) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsearch/Results.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsearch/Results.html:45
	p.StreamBody(qw422016, as, ps)
//line views/vsearch/Results.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/vsearch/Results.html:45
}

//line views/vsearch/Results.html:45
func (p *Results) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsearch/Results.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsearch/Results.html:45
	p.WriteBody(qb422016, as, ps)
//line views/vsearch/Results.html:45
	qs422016 := string(qb422016.B)
//line views/vsearch/Results.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsearch/Results.html:45
	return qs422016
//line views/vsearch/Results.html:45
}
