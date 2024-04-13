// Code generated by qtc from "Migrations.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Migrations.html:2
package vadmin

//line views/vadmin/Migrations.html:2
import (
	"strings"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/database/migrate"
	"github.com/kyleu/todoforge/views/components/view"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vadmin/Migrations.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Migrations.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Migrations.html:12
type Migrations struct {
	layout.Basic
	Available migrate.MigrationFiles
	Applied   migrate.Migrations
}

//line views/vadmin/Migrations.html:18
func (p *Migrations) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Migrations.html:18
	qw422016.N().S(`
  <div class="card">
    <h3>Database Migrations</h3>
  </div>
`)
//line views/vadmin/Migrations.html:22
	for idx, mf := range p.Available {
//line views/vadmin/Migrations.html:22
		qw422016.N().S(`  <div class="card">
    <div class="right">
`)
//line views/vadmin/Migrations.html:25
		curr := p.Applied.GetByIndex(idx + 1)

//line views/vadmin/Migrations.html:26
		if curr == nil {
//line views/vadmin/Migrations.html:26
			qw422016.N().S(`      <em>Not Applied</em>
`)
//line views/vadmin/Migrations.html:28
		} else {
//line views/vadmin/Migrations.html:28
			qw422016.N().S(`      Applied `)
//line views/vadmin/Migrations.html:29
			view.StreamTimestamp(qw422016, &curr.Created)
//line views/vadmin/Migrations.html:29
			qw422016.N().S(`
`)
//line views/vadmin/Migrations.html:30
		}
//line views/vadmin/Migrations.html:30
		qw422016.N().S(`    </div>
    <h3>`)
//line views/vadmin/Migrations.html:32
		qw422016.E().S(mf.Title)
//line views/vadmin/Migrations.html:32
		qw422016.N().S(`</h3>
    `)
//line views/vadmin/Migrations.html:33
		qw422016.N().S(cutil.FormatLangIgnoreErrors(strings.TrimSpace(mf.Content), "sql"))
//line views/vadmin/Migrations.html:33
		qw422016.N().S(`
  </div>
`)
//line views/vadmin/Migrations.html:35
	}
//line views/vadmin/Migrations.html:36
}

//line views/vadmin/Migrations.html:36
func (p *Migrations) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Migrations.html:36
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Migrations.html:36
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Migrations.html:36
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Migrations.html:36
}

//line views/vadmin/Migrations.html:36
func (p *Migrations) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Migrations.html:36
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Migrations.html:36
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Migrations.html:36
	qs422016 := string(qb422016.B)
//line views/vadmin/Migrations.html:36
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Migrations.html:36
	return qs422016
//line views/vadmin/Migrations.html:36
}
