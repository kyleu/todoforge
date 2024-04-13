// Code generated by qtc from "Settings.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vadmin/Settings.html:2
package vadmin

//line views/vadmin/Settings.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vadmin/Settings.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vadmin/Settings.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vadmin/Settings.html:10
type Settings struct {
	layout.Basic
}

//line views/vadmin/Settings.html:14
func (p *Settings) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:14
	qw422016.N().S(`
  <div class="card">
`)
//line views/vadmin/Settings.html:16
	if util.AppSource != "" {
//line views/vadmin/Settings.html:16
		qw422016.N().S(`    <div class="right"><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:17
		qw422016.E().S(util.AppSource)
//line views/vadmin/Settings.html:17
		qw422016.N().S(`"><button>GitHub</button></a></div>
`)
//line views/vadmin/Settings.html:18
	}
//line views/vadmin/Settings.html:18
	qw422016.N().S(`    <h3 title="github:`)
//line views/vadmin/Settings.html:19
	qw422016.E().S(as.BuildInfo.Commit)
//line views/vadmin/Settings.html:19
	qw422016.N().S(`">`)
//line views/vadmin/Settings.html:19
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:19
	qw422016.E().S(util.AppName)
//line views/vadmin/Settings.html:19
	qw422016.N().S(` `)
//line views/vadmin/Settings.html:19
	qw422016.E().S(as.BuildInfo.String())
//line views/vadmin/Settings.html:19
	qw422016.N().S(`</h3>
`)
//line views/vadmin/Settings.html:20
	if util.AppLegal != "" {
//line views/vadmin/Settings.html:20
		qw422016.N().S(`    <div class="mt">`)
//line views/vadmin/Settings.html:21
		qw422016.N().S(util.AppLegal)
//line views/vadmin/Settings.html:21
		qw422016.N().S(`</div>
`)
//line views/vadmin/Settings.html:22
	}
//line views/vadmin/Settings.html:23
	if util.AppURL != "" {
//line views/vadmin/Settings.html:23
		qw422016.N().S(`    <p><a target="_blank" rel="noopener noreferrer" href="`)
//line views/vadmin/Settings.html:24
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:24
		qw422016.N().S(`">`)
//line views/vadmin/Settings.html:24
		qw422016.N().S(util.AppURL)
//line views/vadmin/Settings.html:24
		qw422016.N().S(`</a></p>
`)
//line views/vadmin/Settings.html:25
	}
//line views/vadmin/Settings.html:25
	qw422016.N().S(`    <em>This page is for the settings of the application. To change your user preferences, such as themes, <a href="/profile">edit your profile</a>.</em>
  </div>

  <div class="card">
    <h3>Audits</h3>
    <ul class="mt">
      <li><a href="/admin/audit">`)
//line views/vadmin/Settings.html:32
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:32
	qw422016.N().S(`View audit logs</a></li>
    </ul>
  </div>

  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:37
	components.StreamSVGRefIcon(qw422016, `archive`, ps)
//line views/vadmin/Settings.html:37
	qw422016.N().S(`Admin Functions</h3>
    `)
//line views/vadmin/Settings.html:38
	streamsettingsLink(qw422016, "/admin/server", "archive", "App Information", "All sorts of info about the server and runtime", ps)
//line views/vadmin/Settings.html:38
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:39
	streamsettingsLink(qw422016, "/admin/modules", "archive", "Go Modules", "The Go modules used by "+util.AppName, ps)
//line views/vadmin/Settings.html:39
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:40
	streamsettingsLink(qw422016, "/theme", "archive", "Edit Themes", "Configure the design themes available to end users", ps)
//line views/vadmin/Settings.html:40
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:41
	streamsettingsLink(qw422016, "/admin/logs", "archive", "Recent Logs", "Displays the 100 most recent app logs", ps)
//line views/vadmin/Settings.html:41
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:45
	components.StreamSVGRefIcon(qw422016, `bolt`, ps)
//line views/vadmin/Settings.html:45
	qw422016.N().S(`HTTP Methods</h3>
    `)
//line views/vadmin/Settings.html:46
	streamsettingsLink(qw422016, "/admin/sitemap", "bolt", "Sitemap", "Displays the HTTP actions that are available, with documentation", ps)
//line views/vadmin/Settings.html:46
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:47
	streamsettingsLink(qw422016, "/admin/routes", "bolt", "HTTP routes", "Enumerates all registered HTTP routes, by method", ps)
//line views/vadmin/Settings.html:47
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:48
	streamsettingsLink(qw422016, "/admin/session", "bolt", "User Session", "View the user session, including all cookies and settings", ps)
//line views/vadmin/Settings.html:48
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:49
	streamsettingsLink(qw422016, "/admin/request", "bolt", "Debug HTTP Request", "Full debug view of an HTTP request from your browser", ps)
//line views/vadmin/Settings.html:49
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:53
	components.StreamSVGRefIcon(qw422016, `cog`, ps)
//line views/vadmin/Settings.html:53
	qw422016.N().S(`App Profiling</h3>
    `)
//line views/vadmin/Settings.html:54
	streamsettingsLink(qw422016, "/admin/memusage", "cog", "Memory Usage", "Detailed memory usage statistics", ps)
//line views/vadmin/Settings.html:54
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:55
	streamsettingsLink(qw422016, "/admin/gc", "cog", "Collect Garbage", "Runs the Go garbage collector", ps)
//line views/vadmin/Settings.html:55
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:56
	streamsettingsLink(qw422016, "/admin/heap", "cog", "Write Memory Dump", "Writes a memory dump to <em>./tmp/mem.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:56
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:57
	streamsettingsLink(qw422016, "/admin/cpu/start", "cog", "Start CPU Profile", "Profiles the CPU using <em>./tmp/cpu.pprof</em>, use script to view", ps)
//line views/vadmin/Settings.html:57
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:58
	streamsettingsLink(qw422016, "/admin/cpu/stop", "cog", "Stop CPU Profile", "Stops the active CPU profile", ps)
//line views/vadmin/Settings.html:58
	qw422016.N().S(`
    <div class="clear"></div>
  </div>
  <div class="card">
    <h3>`)
//line views/vadmin/Settings.html:62
	components.StreamSVGRefIcon(qw422016, `database`, ps)
//line views/vadmin/Settings.html:62
	qw422016.N().S(`Database Management</h3>
    `)
//line views/vadmin/Settings.html:63
	streamsettingsLink(qw422016, "/admin/database", "database", "Database Management", "Tools for exploring and manipulating your database", ps)
//line views/vadmin/Settings.html:63
	qw422016.N().S(`
    `)
//line views/vadmin/Settings.html:64
	streamsettingsLink(qw422016, "/admin/migrations", "archive", "Migrations", "Shows the full content of all database SQL migrations", ps)
//line views/vadmin/Settings.html:64
	qw422016.N().S(`
  </div>
`)
//line views/vadmin/Settings.html:66
}

//line views/vadmin/Settings.html:66
func (p *Settings) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vadmin/Settings.html:66
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:66
	p.StreamBody(qw422016, as, ps)
//line views/vadmin/Settings.html:66
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:66
}

//line views/vadmin/Settings.html:66
func (p *Settings) Body(as *app.State, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:66
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:66
	p.WriteBody(qb422016, as, ps)
//line views/vadmin/Settings.html:66
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:66
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:66
	return qs422016
//line views/vadmin/Settings.html:66
}

//line views/vadmin/Settings.html:68
func streamsettingsLink(qw422016 *qt422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:68
	qw422016.N().S(`<hr class="clear" /><div class="mts"><a href="`)
//line views/vadmin/Settings.html:71
	qw422016.E().S(href)
//line views/vadmin/Settings.html:71
	qw422016.N().S(`"><strong>`)
//line views/vadmin/Settings.html:71
	qw422016.E().S(title)
//line views/vadmin/Settings.html:71
	qw422016.N().S(`</strong></a><div><em>`)
//line views/vadmin/Settings.html:72
	qw422016.N().S(description)
//line views/vadmin/Settings.html:72
	qw422016.N().S(`</em></div></div>`)
//line views/vadmin/Settings.html:74
}

//line views/vadmin/Settings.html:74
func writesettingsLink(qq422016 qtio422016.Writer, href string, icon string, title string, description string, ps *cutil.PageState) {
//line views/vadmin/Settings.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vadmin/Settings.html:74
	streamsettingsLink(qw422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/vadmin/Settings.html:74
}

//line views/vadmin/Settings.html:74
func settingsLink(href string, icon string, title string, description string, ps *cutil.PageState) string {
//line views/vadmin/Settings.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vadmin/Settings.html:74
	writesettingsLink(qb422016, href, icon, title, description, ps)
//line views/vadmin/Settings.html:74
	qs422016 := string(qb422016.B)
//line views/vadmin/Settings.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vadmin/Settings.html:74
	return qs422016
//line views/vadmin/Settings.html:74
}
