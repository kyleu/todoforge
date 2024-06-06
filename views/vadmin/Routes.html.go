// Code generated by qtc from "Routes.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Routes.html:2
package vadmin

//line views/vadmin/Routes.html:2
import (
	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vadmin/Routes.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Routes.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Routes.html:12
type Routes struct {
	layout.Basic
	Routes map[string][]string
}

//line views/vadmin/Routes.html:17
func (p *Routes) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Routes.html:17
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/Routes.html:19
	components.StreamSVGIcon(qw422016, `code`, ps)
//line views/vadmin/Routes.html:19
	qw422016.N().S(`HTTP Routes</h3>
    <ul class="mt">
`)
//line views/vadmin/Routes.html:21
	for _, k := range util.ArraySorted(lo.Keys(p.Routes)) {
//line views/vadmin/Routes.html:21
		qw422016.N().S(`      <li>
        <strong>`)
//line views/vadmin/Routes.html:23
		qw422016.E().S(k)
//line views/vadmin/Routes.html:23
		qw422016.N().S(`</strong>
        <ul>
`)
//line views/vadmin/Routes.html:25
		for _, r := range p.Routes[k] {
//line views/vadmin/Routes.html:25
			qw422016.N().S(`          <li><code>`)
//line views/vadmin/Routes.html:26
			qw422016.E().S(r)
//line views/vadmin/Routes.html:26
			qw422016.N().S(`</code></li>
`)
//line views/vadmin/Routes.html:27
		}
//line views/vadmin/Routes.html:27
		qw422016.N().S(`        </ul>
      </li>
`)
//line views/vadmin/Routes.html:30
	}
//line views/vadmin/Routes.html:30
	qw422016.N().S(`    </ul>
  </div>
`)
//line views/vadmin/Routes.html:33
}

//line views/vadmin/Routes.html:33
func (p *Routes) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Routes.html:33
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Routes.html:33
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Routes.html:33
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Routes.html:33
}

//line views/vadmin/Routes.html:33
func (p *Routes) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Routes.html:33
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Routes.html:33
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Routes.html:33
	qs422016 := string(qb422016.B)
//line views/vadmin/Routes.html:33
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Routes.html:33
	return qs422016
//line views/vadmin/Routes.html:33
}
