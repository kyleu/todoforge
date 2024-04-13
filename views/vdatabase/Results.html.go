// Code generated by qtc from "Results.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vdatabase/Results.html:2
package vdatabase

//line views/vdatabase/Results.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/components/view"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vdatabase/Results.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vdatabase/Results.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vdatabase/Results.html:13
type Results struct {
	layout.Basic
	Svc     *database.Service
	Schema  string
	Table   string
	Results []util.ValueMap
	Params  *filter.Params
	Timing  int
	Error   error
}

//line views/vdatabase/Results.html:24
func (p *Results) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Results.html:24
	qw422016.N().S(`
  <div class="card">
    <div class="right">`)
//line views/vdatabase/Results.html:26
	qw422016.E().S(util.MicrosToMillis(p.Timing))
//line views/vdatabase/Results.html:26
	qw422016.N().S(` elapsed</div>
    <h3>`)
//line views/vdatabase/Results.html:27
	components.StreamSVGRefIcon(qw422016, `database`, ps)
//line views/vdatabase/Results.html:27
	qw422016.N().S(`Table [`)
//line views/vdatabase/Results.html:27
	if p.Schema != "default" {
//line views/vdatabase/Results.html:27
		qw422016.E().S(p.Schema)
//line views/vdatabase/Results.html:27
		qw422016.N().S(`:`)
//line views/vdatabase/Results.html:27
	}
//line views/vdatabase/Results.html:27
	qw422016.E().S(p.Table)
//line views/vdatabase/Results.html:27
	qw422016.N().S(`]</h3>
    <div><em>`)
//line views/vdatabase/Results.html:28
	qw422016.N().D(len(p.Results))
//line views/vdatabase/Results.html:28
	qw422016.N().S(` rows returned</em></div>
`)
//line views/vdatabase/Results.html:29
	if p.Error != nil {
//line views/vdatabase/Results.html:29
		qw422016.N().S(`    <div class="mt error">`)
//line views/vdatabase/Results.html:30
		qw422016.E().S(p.Error.Error())
//line views/vdatabase/Results.html:30
		qw422016.N().S(`</div>
`)
//line views/vdatabase/Results.html:31
	}
//line views/vdatabase/Results.html:31
	qw422016.N().S(`    <div class="mt overflow">`)
//line views/vdatabase/Results.html:32
	view.StreamMapArray(qw422016, true, p.Results...)
//line views/vdatabase/Results.html:32
	qw422016.N().S(`</div>
  </div>
`)
//line views/vdatabase/Results.html:34
}

//line views/vdatabase/Results.html:34
func (p *Results) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Results.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Results.html:34
	p.StreamBody(qw422016, as, ps)
//line views/vdatabase/Results.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Results.html:34
}

//line views/vdatabase/Results.html:34
func (p *Results) Body(as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/Results.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Results.html:34
	p.WriteBody(qb422016, as, ps)
//line views/vdatabase/Results.html:34
	qs422016 := string(qb422016.B)
//line views/vdatabase/Results.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Results.html:34
	return qs422016
//line views/vdatabase/Results.html:34
}
