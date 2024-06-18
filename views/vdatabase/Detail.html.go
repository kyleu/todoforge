// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vdatabase/Detail.html:1
package vdatabase

//line views/vdatabase/Detail.html:1
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vdatabase/Detail.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vdatabase/Detail.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vdatabase/Detail.html:10
type Detail struct {
	layout.Basic
	Mode    string
	Svc     *database.Service
	Recent  database.DebugStatements
	Sizes   database.TableSizes
	SQL     string
	Columns []string
	Results [][]any
	Timing  int
	Commit  bool
}

//line views/vdatabase/Detail.html:23
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:23
	qw422016.N().S(`
  <div class="card">
    <div class="right"><em>`)
//line views/vdatabase/Detail.html:25
	qw422016.E().S(p.Svc.Type.Title)
//line views/vdatabase/Detail.html:25
	qw422016.N().S(`</em></div>
    <h3>`)
//line views/vdatabase/Detail.html:26
	components.StreamSVGIcon(qw422016, `database`, ps)
//line views/vdatabase/Detail.html:26
	qw422016.N().S(` `)
//line views/vdatabase/Detail.html:26
	qw422016.E().S(p.Svc.Key)
//line views/vdatabase/Detail.html:26
	qw422016.N().S(`</h3>
    <div><em>`)
//line views/vdatabase/Detail.html:27
	qw422016.E().S(p.Svc.String())
//line views/vdatabase/Detail.html:27
	qw422016.N().S(`</em></div>
    <div class="mt">
`)
//line views/vdatabase/Detail.html:29
	if p.Svc.Tracing() == "" {
//line views/vdatabase/Detail.html:29
		qw422016.N().S(`      <em>tracing is disabled</em>
`)
//line views/vdatabase/Detail.html:31
	} else {
//line views/vdatabase/Detail.html:31
		qw422016.N().S(`      <em>tracing is enabled in [`)
//line views/vdatabase/Detail.html:32
		qw422016.E().S(p.Svc.Tracing())
//line views/vdatabase/Detail.html:32
		qw422016.N().S(`] mode</em>
`)
//line views/vdatabase/Detail.html:33
	}
//line views/vdatabase/Detail.html:33
	qw422016.N().S(`    </div>
    <div class="mt">
      <a href="#modal-settings"><button>Tracing Settings</button></a>
      `)
//line views/vdatabase/Detail.html:37
	streamsettingsModal(qw422016, p.Svc)
//line views/vdatabase/Detail.html:37
	qw422016.N().S(`
`)
//line views/vdatabase/Detail.html:38
	if p.Svc.Tracing() != "" {
//line views/vdatabase/Detail.html:38
		qw422016.N().S(`      <a href="/admin/database/`)
//line views/vdatabase/Detail.html:39
		qw422016.E().S(p.Svc.Key)
//line views/vdatabase/Detail.html:39
		qw422016.N().S(`/recent"><button>Recent Activity</button></a>
`)
//line views/vdatabase/Detail.html:40
	}
//line views/vdatabase/Detail.html:40
	qw422016.N().S(`      <a href="/admin/database/`)
//line views/vdatabase/Detail.html:41
	qw422016.E().S(p.Svc.Key)
//line views/vdatabase/Detail.html:41
	qw422016.N().S(`/tables"><button>Tables</button></a>
      <a href="/admin/database/`)
//line views/vdatabase/Detail.html:42
	qw422016.E().S(p.Svc.Key)
//line views/vdatabase/Detail.html:42
	qw422016.N().S(`/analyze"><button>Analyze</button></a>
      <a href="/admin/database/`)
//line views/vdatabase/Detail.html:43
	qw422016.E().S(p.Svc.Key)
//line views/vdatabase/Detail.html:43
	qw422016.N().S(`/sql"><button>SQL</button></a>
    </div>
  </div>
`)
//line views/vdatabase/Detail.html:46
	switch p.Mode {
//line views/vdatabase/Detail.html:47
	case "recent":
//line views/vdatabase/Detail.html:47
		qw422016.N().S(`  `)
//line views/vdatabase/Detail.html:48
		streamrecentStatements(qw422016, p.Recent, p.Svc, as, ps)
//line views/vdatabase/Detail.html:48
		qw422016.N().S(`
`)
//line views/vdatabase/Detail.html:49
	case "tables":
//line views/vdatabase/Detail.html:49
		qw422016.N().S(`  `)
//line views/vdatabase/Detail.html:50
		streamtableSizes(qw422016, p.Svc.Key, p.Sizes, as, ps)
//line views/vdatabase/Detail.html:50
		qw422016.N().S(`
`)
//line views/vdatabase/Detail.html:51
	case "sql":
//line views/vdatabase/Detail.html:51
		qw422016.N().S(`  `)
//line views/vdatabase/Detail.html:52
		streamsqlEditor(qw422016, p.SQL, p.Svc, p.Commit, p.Columns, p.Results, p.Timing, as, ps)
//line views/vdatabase/Detail.html:52
		qw422016.N().S(`
`)
//line views/vdatabase/Detail.html:53
	}
//line views/vdatabase/Detail.html:54
}

//line views/vdatabase/Detail.html:54
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Detail.html:54
	p.StreamBody(qw422016, as, ps)
//line views/vdatabase/Detail.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Detail.html:54
}

//line views/vdatabase/Detail.html:54
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/Detail.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Detail.html:54
	p.WriteBody(qb422016, as, ps)
//line views/vdatabase/Detail.html:54
	qs422016 := string(qb422016.B)
//line views/vdatabase/Detail.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Detail.html:54
	return qs422016
//line views/vdatabase/Detail.html:54
}

//line views/vdatabase/Detail.html:56
func streamrecentStatements(qw422016 *qt422016.Writer, recent database.DebugStatements, svc *database.Service, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:56
	qw422016.N().S(`
  <div class="card">
    <h3>Recent Activity</h3>
`)
//line views/vdatabase/Detail.html:59
	if len(recent) == 0 {
//line views/vdatabase/Detail.html:60
		if svc.Tracing() == "" {
//line views/vdatabase/Detail.html:60
			qw422016.N().S(`      <em>Tracing is not enabled for this database</em>
`)
//line views/vdatabase/Detail.html:62
		} else {
//line views/vdatabase/Detail.html:62
			qw422016.N().S(`      <em>No recent statements</em>
`)
//line views/vdatabase/Detail.html:64
		}
//line views/vdatabase/Detail.html:65
	} else {
//line views/vdatabase/Detail.html:65
		qw422016.N().S(`    <div class="overflow full-width">
      <table>
        <thead>
        <tr>
          <th>SQL</th>
          <th>Values</th>
          <th>Count</th>
          <th>Status</th>
          <th>Message</th>
          <th>Duration</th>
        </tr>
        </thead>
        <tbody>
`)
//line views/vdatabase/Detail.html:79
		for _, s := range recent {
//line views/vdatabase/Detail.html:79
			qw422016.N().S(`          <tr>
            <td>
              <a href="?idx=`)
//line views/vdatabase/Detail.html:82
			qw422016.N().D(s.Index)
//line views/vdatabase/Detail.html:82
			qw422016.N().S(`">`)
//line views/vdatabase/Detail.html:82
			qw422016.E().S(s.SQLTrimmed(100))
//line views/vdatabase/Detail.html:82
			qw422016.N().S(`</a>
            </td>
            <td>`)
//line views/vdatabase/Detail.html:84
			qw422016.N().D(len(s.Values))
//line views/vdatabase/Detail.html:84
			qw422016.N().S(`</td>
            <td>`)
//line views/vdatabase/Detail.html:85
			qw422016.N().D(s.Count)
//line views/vdatabase/Detail.html:85
			qw422016.N().S(`</td>
            <td>
`)
//line views/vdatabase/Detail.html:87
			if s.Error == "" {
//line views/vdatabase/Detail.html:87
				qw422016.N().S(`              OK
`)
//line views/vdatabase/Detail.html:89
			} else {
//line views/vdatabase/Detail.html:89
				qw422016.N().S(`              <span class="error">[error]: `)
//line views/vdatabase/Detail.html:90
				qw422016.E().S(s.ErrorTrimmed(100))
//line views/vdatabase/Detail.html:90
				qw422016.N().S(`</span>
`)
//line views/vdatabase/Detail.html:91
			}
//line views/vdatabase/Detail.html:91
			qw422016.N().S(`            </td>
            <td>`)
//line views/vdatabase/Detail.html:93
			qw422016.E().S(s.Message)
//line views/vdatabase/Detail.html:93
			qw422016.N().S(`</td>
            <td>`)
//line views/vdatabase/Detail.html:94
			qw422016.E().S(util.MicrosToMillis(s.Timing))
//line views/vdatabase/Detail.html:94
			qw422016.N().S(`</td>
          </tr>
`)
//line views/vdatabase/Detail.html:96
		}
//line views/vdatabase/Detail.html:96
		qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vdatabase/Detail.html:100
	}
//line views/vdatabase/Detail.html:100
	qw422016.N().S(`  </div>
`)
//line views/vdatabase/Detail.html:102
}

//line views/vdatabase/Detail.html:102
func writerecentStatements(qq422016 qtio422016.Writer, recent database.DebugStatements, svc *database.Service, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:102
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Detail.html:102
	streamrecentStatements(qw422016, recent, svc, as, ps)
//line views/vdatabase/Detail.html:102
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Detail.html:102
}

//line views/vdatabase/Detail.html:102
func recentStatements(recent database.DebugStatements, svc *database.Service, as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/Detail.html:102
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Detail.html:102
	writerecentStatements(qb422016, recent, svc, as, ps)
//line views/vdatabase/Detail.html:102
	qs422016 := string(qb422016.B)
//line views/vdatabase/Detail.html:102
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Detail.html:102
	return qs422016
//line views/vdatabase/Detail.html:102
}

//line views/vdatabase/Detail.html:104
func streamtableSizes(qw422016 *qt422016.Writer, key string, sizes database.TableSizes, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:104
	qw422016.N().S(`
  <div class="card">
    <h3>Table Sizes</h3>
    <div class="overflow full-width">
      <table class="min-200">
        <thead>
          <tr>
            <th class="shrink">Name</th>
            <th title="(estimated)">Rows*</th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vdatabase/Detail.html:117
	for _, size := range sizes {
//line views/vdatabase/Detail.html:117
		qw422016.N().S(`          <tr>
            <td>`)
//line views/vdatabase/Detail.html:119
		qw422016.E().S(size.Name)
//line views/vdatabase/Detail.html:119
		qw422016.N().S(`</td>
            <td>`)
//line views/vdatabase/Detail.html:120
		qw422016.E().S(size.Rows)
//line views/vdatabase/Detail.html:120
		qw422016.N().S(`</td>
            <td style="white-space: nowrap;">
              <a href="/admin/database/`)
//line views/vdatabase/Detail.html:122
		qw422016.E().S(key)
//line views/vdatabase/Detail.html:122
		qw422016.N().S(`/tables/`)
//line views/vdatabase/Detail.html:122
		qw422016.E().S(size.Schema)
//line views/vdatabase/Detail.html:122
		qw422016.N().S(`/`)
//line views/vdatabase/Detail.html:122
		qw422016.E().S(size.Name)
//line views/vdatabase/Detail.html:122
		qw422016.N().S(`"><button>Data</button></a>
              <a href="/admin/database/`)
//line views/vdatabase/Detail.html:123
		qw422016.E().S(key)
//line views/vdatabase/Detail.html:123
		qw422016.N().S(`/tables/`)
//line views/vdatabase/Detail.html:123
		qw422016.E().S(size.Schema)
//line views/vdatabase/Detail.html:123
		qw422016.N().S(`/`)
//line views/vdatabase/Detail.html:123
		qw422016.E().S(size.Name)
//line views/vdatabase/Detail.html:123
		qw422016.N().S(`/stats"><button>Stats</button></a>
            </td>
          </tr>
`)
//line views/vdatabase/Detail.html:126
	}
//line views/vdatabase/Detail.html:126
	qw422016.N().S(`        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vdatabase/Detail.html:131
}

//line views/vdatabase/Detail.html:131
func writetableSizes(qq422016 qtio422016.Writer, key string, sizes database.TableSizes, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:131
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Detail.html:131
	streamtableSizes(qw422016, key, sizes, as, ps)
//line views/vdatabase/Detail.html:131
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Detail.html:131
}

//line views/vdatabase/Detail.html:131
func tableSizes(key string, sizes database.TableSizes, as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/Detail.html:131
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Detail.html:131
	writetableSizes(qb422016, key, sizes, as, ps)
//line views/vdatabase/Detail.html:131
	qs422016 := string(qb422016.B)
//line views/vdatabase/Detail.html:131
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Detail.html:131
	return qs422016
//line views/vdatabase/Detail.html:131
}

//line views/vdatabase/Detail.html:133
func streamsqlEditor(qw422016 *qt422016.Writer, sql string, svc *database.Service, commit bool, columns []string, results [][]any, timing int, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:133
	qw422016.N().S(`
  <div class="card">
    <h3>SQL Editor</h3>
    <form method="post" action="/admin/database/`)
//line views/vdatabase/Detail.html:136
	qw422016.E().S(svc.Key)
//line views/vdatabase/Detail.html:136
	qw422016.N().S(`/sql">
      <div class="mt expanded">
        <textarea name="sql" rows="12" placeholder="SQL statement">`)
//line views/vdatabase/Detail.html:138
	qw422016.E().S(sql)
//line views/vdatabase/Detail.html:138
	qw422016.N().S(`</textarea>
      </div>
`)
//line views/vdatabase/Detail.html:140
	if svc.ReadOnly {
//line views/vdatabase/Detail.html:140
		qw422016.N().S(`      <input type="hidden" name="commit" value="false" />
`)
//line views/vdatabase/Detail.html:142
	} else {
//line views/vdatabase/Detail.html:142
		qw422016.N().S(`      <div class="mt">
        <label><input type="checkbox" name="commit" value="true" `)
//line views/vdatabase/Detail.html:144
		if commit {
//line views/vdatabase/Detail.html:144
			qw422016.N().S(`checked="checked"`)
//line views/vdatabase/Detail.html:144
		}
//line views/vdatabase/Detail.html:144
		qw422016.N().S(`/> Commit Changes</label>
      </div>
`)
//line views/vdatabase/Detail.html:146
	}
//line views/vdatabase/Detail.html:146
	qw422016.N().S(`      <div class="mt">
        <button type="submit" name="action" value="run">Run</button>
        <button type="submit" name="action" value="analyze">Analyze</button>
      </div>
    </form>
  </div>
`)
//line views/vdatabase/Detail.html:153
	if results != nil {
//line views/vdatabase/Detail.html:153
		qw422016.N().S(`  <div class="card">
    <div class="right">`)
//line views/vdatabase/Detail.html:155
		qw422016.E().S(util.MicrosToMillis(timing))
//line views/vdatabase/Detail.html:155
		qw422016.N().S(`</div>
    <h3>Results</h3>
`)
//line views/vdatabase/Detail.html:157
		if len(results) == 0 {
//line views/vdatabase/Detail.html:157
			qw422016.N().S(`    <em>No rows returned</em>
`)
//line views/vdatabase/Detail.html:159
		} else {
//line views/vdatabase/Detail.html:159
			qw422016.N().S(`    <div class="overflow full-width">
      <table class="mt expanded">
        <thead>
          <tr>
`)
//line views/vdatabase/Detail.html:164
			for _, c := range columns {
//line views/vdatabase/Detail.html:164
				qw422016.N().S(`            <th>`)
//line views/vdatabase/Detail.html:165
				qw422016.E().S(c)
//line views/vdatabase/Detail.html:165
				qw422016.N().S(`</th>
`)
//line views/vdatabase/Detail.html:166
			}
//line views/vdatabase/Detail.html:166
			qw422016.N().S(`          </tr>
        </thead>
        <tbody>
`)
//line views/vdatabase/Detail.html:170
			for _, row := range results {
//line views/vdatabase/Detail.html:170
				qw422016.N().S(`            <tr>
`)
//line views/vdatabase/Detail.html:172
				for _, x := range row {
//line views/vdatabase/Detail.html:172
					qw422016.N().S(`              <td>`)
//line views/vdatabase/Detail.html:173
					qw422016.E().V(x)
//line views/vdatabase/Detail.html:173
					qw422016.N().S(`</td>
`)
//line views/vdatabase/Detail.html:174
				}
//line views/vdatabase/Detail.html:174
				qw422016.N().S(`            </tr>
`)
//line views/vdatabase/Detail.html:176
			}
//line views/vdatabase/Detail.html:176
			qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vdatabase/Detail.html:180
		}
//line views/vdatabase/Detail.html:180
		qw422016.N().S(`  </div>
`)
//line views/vdatabase/Detail.html:182
	}
//line views/vdatabase/Detail.html:183
}

//line views/vdatabase/Detail.html:183
func writesqlEditor(qq422016 qtio422016.Writer, sql string, svc *database.Service, commit bool, columns []string, results [][]any, timing int, as *app.State, ps *cutil.PageState) {
//line views/vdatabase/Detail.html:183
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Detail.html:183
	streamsqlEditor(qw422016, sql, svc, commit, columns, results, timing, as, ps)
//line views/vdatabase/Detail.html:183
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Detail.html:183
}

//line views/vdatabase/Detail.html:183
func sqlEditor(sql string, svc *database.Service, commit bool, columns []string, results [][]any, timing int, as *app.State, ps *cutil.PageState) string {
//line views/vdatabase/Detail.html:183
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Detail.html:183
	writesqlEditor(qb422016, sql, svc, commit, columns, results, timing, as, ps)
//line views/vdatabase/Detail.html:183
	qs422016 := string(qb422016.B)
//line views/vdatabase/Detail.html:183
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Detail.html:183
	return qs422016
//line views/vdatabase/Detail.html:183
}

//line views/vdatabase/Detail.html:185
func streamsettingsModal(qw422016 *qt422016.Writer, svc *database.Service) {
//line views/vdatabase/Detail.html:185
	qw422016.N().S(`
  <div id="modal-settings" class="modal" style="display: none;">
    <a class="backdrop" href="#"></a>
    <div class="modal-content">
      <div class="modal-header">
        <a href="#" class="modal-close">×</a>
        <h2>Tracing Settings</h2>
      </div>
      <div class="modal-body">
        <form action="/admin/database/`)
//line views/vdatabase/Detail.html:194
	qw422016.E().S(svc.Key)
//line views/vdatabase/Detail.html:194
	qw422016.N().S(`/enable">
          <div class="overflow full-width">
            <table>
              <tbody>
`)
//line views/vdatabase/Detail.html:198
	trc := svc.Tracing()

//line views/vdatabase/Detail.html:198
	qw422016.N().S(`
                <tr>
                  <td><label><input type="radio" name="tracing" value=""`)
//line views/vdatabase/Detail.html:200
	if trc == `` {
//line views/vdatabase/Detail.html:200
		qw422016.N().S(` checked="checked"`)
//line views/vdatabase/Detail.html:200
	}
//line views/vdatabase/Detail.html:200
	qw422016.N().S(`> No Tracing</label></td>
                  <td><em>Fastest configuration, no tracing overhead</em></td>
                </tr>
                <tr>
                  <td><label><input type="radio" name="tracing" value="statement"`)
//line views/vdatabase/Detail.html:204
	if trc == `statement` {
//line views/vdatabase/Detail.html:204
		qw422016.N().S(` checked="checked"`)
//line views/vdatabase/Detail.html:204
	}
//line views/vdatabase/Detail.html:204
	qw422016.N().S(`> Save Queries</label></td>
                  <td><em>Save most recent 100 SQL statements with timing information</em></td>
                </tr>
                <tr>
                  <td><label><input type="radio" name="tracing" value="values"`)
//line views/vdatabase/Detail.html:208
	if trc == `values` {
//line views/vdatabase/Detail.html:208
		qw422016.N().S(` checked="checked"`)
//line views/vdatabase/Detail.html:208
	}
//line views/vdatabase/Detail.html:208
	qw422016.N().S(`> Save Results</label></td>
                  <td><em>Saves SQL, timing, and the results of the query</em></td>
                </tr>
                <tr>
                  <td><label><input type="radio" name="tracing" value="analyze"`)
//line views/vdatabase/Detail.html:212
	if trc == `analyze` {
//line views/vdatabase/Detail.html:212
		qw422016.N().S(` checked="checked"`)
//line views/vdatabase/Detail.html:212
	}
//line views/vdatabase/Detail.html:212
	qw422016.N().S(`> Analyze Queries</label></td>
                  <td><em>In addition to the above, runs an explain plan on each query</em></td>
                </tr>
                <tr>
                  <td colspan="2"><button>Submit</button></td>
                </tr>
              </tbody>
            </table>
          </div>
        </form>
      </div>
    </div>
  </div>
`)
//line views/vdatabase/Detail.html:225
}

//line views/vdatabase/Detail.html:225
func writesettingsModal(qq422016 qtio422016.Writer, svc *database.Service) {
//line views/vdatabase/Detail.html:225
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vdatabase/Detail.html:225
	streamsettingsModal(qw422016, svc)
//line views/vdatabase/Detail.html:225
	qt422016.ReleaseWriter(qw422016)
//line views/vdatabase/Detail.html:225
}

//line views/vdatabase/Detail.html:225
func settingsModal(svc *database.Service) string {
//line views/vdatabase/Detail.html:225
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vdatabase/Detail.html:225
	writesettingsModal(qb422016, svc)
//line views/vdatabase/Detail.html:225
	qs422016 := string(qb422016.B)
//line views/vdatabase/Detail.html:225
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vdatabase/Detail.html:225
	return qs422016
//line views/vdatabase/Detail.html:225
}
