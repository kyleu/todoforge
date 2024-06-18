// Code generated by qtc from "Unauthorized.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/verror/Unauthorized.html:1
package verror

//line views/verror/Unauthorized.html:1
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/verror/Unauthorized.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/verror/Unauthorized.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/verror/Unauthorized.html:7
type Unauthorized struct {
	layout.Basic
	Path    string
	Message string
}

//line views/verror/Unauthorized.html:13
func (p *Unauthorized) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Unauthorized.html:13
	qw422016.N().S(`
  <div class="card">
    <h3>Unauthorized</h3>
    <em>`)
//line views/verror/Unauthorized.html:16
	qw422016.E().S(p.Message)
//line views/verror/Unauthorized.html:16
	qw422016.N().S(`</em>
    <p>You're not authorized to view <code>`)
//line views/verror/Unauthorized.html:17
	qw422016.E().S(p.Path)
//line views/verror/Unauthorized.html:17
	qw422016.N().S(`</code></p>
  </div>
`)
//line views/verror/Unauthorized.html:19
}

//line views/verror/Unauthorized.html:19
func (p *Unauthorized) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/verror/Unauthorized.html:19
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/verror/Unauthorized.html:19
	p.StreamBody(qw422016, as, ps)
//line views/verror/Unauthorized.html:19
	qt422016.ReleaseWriter(qw422016)
//line views/verror/Unauthorized.html:19
}

//line views/verror/Unauthorized.html:19
func (p *Unauthorized) Body(as *app.State, ps *cutil.PageState) string {
//line views/verror/Unauthorized.html:19
	qb422016 := qt422016.AcquireByteBuffer()
//line views/verror/Unauthorized.html:19
	p.WriteBody(qb422016, as, ps)
//line views/verror/Unauthorized.html:19
	qs422016 := string(qb422016.B)
//line views/verror/Unauthorized.html:19
	qt422016.ReleaseByteBuffer(qb422016)
//line views/verror/Unauthorized.html:19
	return qs422016
//line views/verror/Unauthorized.html:19
}
