// Code generated by qtc from "GoSource.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vsite/GoSource.html:1
package vsite

//line views/vsite/GoSource.html:1
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vsite/GoSource.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsite/GoSource.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsite/GoSource.html:9
type GoSource struct{ layout.Basic }

//line views/vsite/GoSource.html:11
func (p *GoSource) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/GoSource.html:11
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vsite/GoSource.html:13
	components.StreamSVGIcon(qw422016, `app`, ps)
//line views/vsite/GoSource.html:13
	qw422016.N().S(` `)
//line views/vsite/GoSource.html:13
	qw422016.E().S(util.AppName)
//line views/vsite/GoSource.html:13
	qw422016.N().S(`</h3>
    <pre>go install github.com/kyleu/todoforge@latest</pre>
    <p>
      <a href="https://github.com/kyleu/todoforge"><button>Source Code</button></a>
    </p>
  </div>
`)
//line views/vsite/GoSource.html:19
}

//line views/vsite/GoSource.html:19
func (p *GoSource) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vsite/GoSource.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsite/GoSource.html:19
	p.StreamBody(qw422016, as, ps)
//line views/vsite/GoSource.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/vsite/GoSource.html:19
}

//line views/vsite/GoSource.html:19
func (p *GoSource) Body(as *app.State, ps *cutil.PageState) string {
//line views/vsite/GoSource.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsite/GoSource.html:19
	p.WriteBody(qb422016, as, ps)
//line views/vsite/GoSource.html:19
	qs422016 := string(qb422016.B)
//line views/vsite/GoSource.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsite/GoSource.html:19
	return qs422016
//line views/vsite/GoSource.html:19
}
