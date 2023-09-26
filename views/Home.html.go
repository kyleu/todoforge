// Code generated by qtc from "Home.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/Home.html:1
package views

//line views/Home.html:1
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/Home.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/Home.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/Home.html:9
type Home struct {
	layout.Basic
}

//line views/Home.html:13
func (p *Home) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Home.html:13
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/Home.html:15
	components.StreamSVGRefIcon(qw422016, `app`, ps)
//line views/Home.html:15
	qw422016.E().S(util.AppName)
//line views/Home.html:15
	qw422016.N().S(`</h3>
  </div>
`)
//line views/Home.html:17
}

//line views/Home.html:17
func (p *Home) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/Home.html:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/Home.html:17
	p.StreamBody(qw422016, as, ps)
//line views/Home.html:17
	qt422016.ReleaseWriter(qw422016)
//line views/Home.html:17
}

//line views/Home.html:17
func (p *Home) Body(as *app.State, ps *cutil.PageState) string {
//line views/Home.html:17
	qb422016 := qt422016.AcquireByteBuffer()
//line views/Home.html:17
	p.WriteBody(qb422016, as, ps)
//line views/Home.html:17
	qs422016 := string(qb422016.B)
//line views/Home.html:17
	qt422016.ReleaseByteBuffer(qb422016)
//line views/Home.html:17
	return qs422016
//line views/Home.html:17
}