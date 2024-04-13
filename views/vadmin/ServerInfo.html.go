// Code generated by qtc from "ServerInfo.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/ServerInfo.html:2
package vadmin

//line views/vadmin/ServerInfo.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vadmin/ServerInfo.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/ServerInfo.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/ServerInfo.html:10
type ServerInfo struct {
	layout.Basic
	Info *util.DebugInfo
}

//line views/vadmin/ServerInfo.html:15
func (p *ServerInfo) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/ServerInfo.html:15
	qw422016.N().S(`
  `)
//line views/vadmin/ServerInfo.html:16
	streamrenderTags(qw422016, "Server Information", p.Info.ServerTags, "cog", ps)
//line views/vadmin/ServerInfo.html:16
	qw422016.N().S(`
  `)
//line views/vadmin/ServerInfo.html:17
	streamrenderTags(qw422016, "Runtime Information", p.Info.RuntimeTags, "desktop", ps)
//line views/vadmin/ServerInfo.html:17
	qw422016.N().S(`
  `)
//line views/vadmin/ServerInfo.html:18
	streamrenderTags(qw422016, "App Information", p.Info.AppTags, "play", ps)
//line views/vadmin/ServerInfo.html:18
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/ServerInfo.html:20
	components.StreamSVGRefIcon(qw422016, `archive`, ps)
//line views/vadmin/ServerInfo.html:20
	qw422016.N().S(`Go Modules</h3>
    `)
//line views/vadmin/ServerInfo.html:21
	streammoduleTable(qw422016, p.Info.Mods)
//line views/vadmin/ServerInfo.html:21
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/ServerInfo.html:23
}

//line views/vadmin/ServerInfo.html:23
func (p *ServerInfo) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/ServerInfo.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/ServerInfo.html:23
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/ServerInfo.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/ServerInfo.html:23
}

//line views/vadmin/ServerInfo.html:23
func (p *ServerInfo) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/ServerInfo.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/ServerInfo.html:23
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/ServerInfo.html:23
	qs422016 := string(qb422016.B)
//line views/vadmin/ServerInfo.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/ServerInfo.html:23
	return qs422016
//line views/vadmin/ServerInfo.html:23
}

//line views/vadmin/ServerInfo.html:25
func streamrenderTags(qw422016 *qt422016.Writer, title string, tags *util.OrderedMap[string], icon string, ps *cutil.PageState) {
//line views/vadmin/ServerInfo.html:25
	qw422016.N().S(`
  <div class="card">
    <h3>`)
//line views/vadmin/ServerInfo.html:27
	components.StreamSVGRefIcon(qw422016, icon, ps)
//line views/vadmin/ServerInfo.html:27
	qw422016.E().S(title)
//line views/vadmin/ServerInfo.html:27
	qw422016.N().S(`</h3>
    <div class="overflow full-width">
      <table class="mt min-200">
        <tbody>
`)
//line views/vadmin/ServerInfo.html:31
	for _, k := range tags.Order {
//line views/vadmin/ServerInfo.html:31
		qw422016.N().S(`        <tr>
          <th class="shrink">`)
//line views/vadmin/ServerInfo.html:33
		qw422016.E().S(k)
//line views/vadmin/ServerInfo.html:33
		qw422016.N().S(`</th>
          <td>`)
//line views/vadmin/ServerInfo.html:34
		qw422016.E().S(tags.GetSimple(k))
//line views/vadmin/ServerInfo.html:34
		qw422016.N().S(`</td>
        </tr>
`)
//line views/vadmin/ServerInfo.html:36
	}
//line views/vadmin/ServerInfo.html:36
	qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vadmin/ServerInfo.html:41
}

//line views/vadmin/ServerInfo.html:41
func writerenderTags(qq422016 qtio422016.Writer, title string, tags *util.OrderedMap[string], icon string, ps *cutil.PageState) {
//line views/vadmin/ServerInfo.html:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/ServerInfo.html:41
	streamrenderTags(qw422016, title, tags, icon, ps)
//line views/vadmin/ServerInfo.html:41
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/ServerInfo.html:41
}

//line views/vadmin/ServerInfo.html:41
func renderTags(title string, tags *util.OrderedMap[string], icon string, ps *cutil.PageState) string {
//line views/vadmin/ServerInfo.html:41
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/ServerInfo.html:41
	writerenderTags(qb422016, title, tags, icon, ps)
//line views/vadmin/ServerInfo.html:41
	qs422016 := string(qb422016.B)
//line views/vadmin/ServerInfo.html:41
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/ServerInfo.html:41
	return qs422016
//line views/vadmin/ServerInfo.html:41
}
