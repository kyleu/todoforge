// Code generated by qtc from "Index.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- $PF_IGNORE$ -->

//line views/vsite/Index.html:2
package vsite

//line views/vsite/Index.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vsite/Index.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/Index.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/Index.html:10
type Index struct{ layout.Basic }

//line views/vsite/Index.html:12
func (p *Index) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Index.html:12
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsite/Index.html:14
	components.StreamSVGIcon(qw422016, `app`, ps)
//line views/vsite/Index.html:14
	qw422016.E().S(util.AppName)
//line views/vsite/Index.html:14
	qw422016.N().S(`</h3>
    <p>This app is almost entirely generated via Project Forge. It manages collections of todo items, and not much else.</p>
    <p>
      <a href="https://tododemo.kyleu.dev"><button>Demo</button></a> <em>(please be nice)</em>
    </p>
    <p>
      <a href="/download"><button>Download</button></a>
      <a href="https://github.com/kyleu/todoforge"><button>Source Code</button></a>
    </p>
  </div>
`)
//line views/vsite/Index.html:24
}

//line views/vsite/Index.html:24
func (p *Index) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/Index.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/Index.html:24
	p.StreamBody(qw422016, as, ps)
//line views/vsite/Index.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/Index.html:24
}

//line views/vsite/Index.html:24
func (p *Index) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/Index.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/Index.html:24
	p.WriteBody(qb422016, as, ps)
//line views/vsite/Index.html:24
	qs422016 := string(qb422016.B)
//line views/vsite/Index.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/Index.html:24
	return qs422016
//line views/vsite/Index.html:24
}
