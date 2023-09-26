// Code generated by qtc from "Edit.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vtheme/Edit.html:2
package vtheme

//line views/vtheme/Edit.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/theme"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vtheme/Edit.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vtheme/Edit.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vtheme/Edit.html:11
type Edit struct {
	layout.Basic
	Theme  *theme.Theme
	Icon   string
	Exists bool
}

//line views/vtheme/Edit.html:18
func (p *Edit) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Edit.html:18
	qw422016.N().S(`
  <form action="/theme/save" method="post">
    <input type="hidden" name="originalKey" value="`)
//line views/vtheme/Edit.html:20
	qw422016.E().S(p.Theme.Key)
//line views/vtheme/Edit.html:20
	qw422016.N().S(`" />
    <div class="card">
      <div class="right">
`)
//line views/vtheme/Edit.html:23
	if p.Exists {
//line views/vtheme/Edit.html:23
		qw422016.N().S(`        <a href="/theme/`)
//line views/vtheme/Edit.html:24
		qw422016.E().S(p.Theme.Key)
//line views/vtheme/Edit.html:24
		qw422016.N().S(`/remove" class="link-confirm" data-message="Are you sure you wish to delete the [`)
//line views/vtheme/Edit.html:24
		qw422016.E().S(p.Theme.Key)
//line views/vtheme/Edit.html:24
		qw422016.N().S(`] theme?"><button type="button">Remove</button></a>
`)
//line views/vtheme/Edit.html:25
	} else {
//line views/vtheme/Edit.html:25
		qw422016.N().S(`        <em>built-in</em>
`)
//line views/vtheme/Edit.html:27
	}
//line views/vtheme/Edit.html:27
	qw422016.N().S(`        <a href="#modal-theme"><button type="button">JSON</button></a>
      </div>
      <h3>`)
//line views/vtheme/Edit.html:30
	if p.Theme.Key == theme.KeyNew {
//line views/vtheme/Edit.html:30
		qw422016.N().S(`New Theme`)
//line views/vtheme/Edit.html:30
	} else {
//line views/vtheme/Edit.html:30
		qw422016.N().S(`Theme Edit`)
//line views/vtheme/Edit.html:30
	}
//line views/vtheme/Edit.html:30
	qw422016.N().S(`</h3>
      <div class="overflow full-width">
        <table class="mt expanded">
          <tbody>
            `)
//line views/vtheme/Edit.html:34
	components.StreamTableInput(qw422016, "key", "", "Key", p.Theme.Key, 5)
//line views/vtheme/Edit.html:34
	qw422016.N().S(`
          </tbody>
        </table>
      </div>
    </div>
    `)
//line views/vtheme/Edit.html:39
	StreamEditor(qw422016, "Theme ["+p.Theme.Key+"]", util.AppName, p.Theme, p.Icon, as, ps)
//line views/vtheme/Edit.html:39
	qw422016.N().S(`
    <div class="card">
      <button type="submit">Save All Changes</button>
      <a href="/theme/`)
//line views/vtheme/Edit.html:42
	qw422016.N().U(p.Theme.Key)
//line views/vtheme/Edit.html:42
	qw422016.N().S(`"><button type="button">Reset</button></a>
    </div>
  </form>
  `)
//line views/vtheme/Edit.html:45
	components.StreamJSONModal(qw422016, "theme", "Theme JSON", p.Theme, 1)
//line views/vtheme/Edit.html:45
	qw422016.N().S(`
`)
//line views/vtheme/Edit.html:46
}

//line views/vtheme/Edit.html:46
func (p *Edit) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vtheme/Edit.html:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vtheme/Edit.html:46
	p.StreamBody(qw422016, as, ps)
//line views/vtheme/Edit.html:46
	qt422016.ReleaseWriter(qw422016)
//line views/vtheme/Edit.html:46
}

//line views/vtheme/Edit.html:46
func (p *Edit) Body(as *app.State, ps *cutil.PageState) string {
//line views/vtheme/Edit.html:46
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vtheme/Edit.html:46
	p.WriteBody(qb422016, as, ps)
//line views/vtheme/Edit.html:46
	qs422016 := string(qb422016.B)
//line views/vtheme/Edit.html:46
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vtheme/Edit.html:46
	return qs422016
//line views/vtheme/Edit.html:46
}