// Code generated by qtc from "Table.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vcollection/Table.html:2
package vcollection

//line views/vcollection/Table.html:2
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/views/components"
)

//line views/vcollection/Table.html:10
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcollection/Table.html:10
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcollection/Table.html:10
func StreamTable(qw422016 *qt422016.Writer, models collection.Collections, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vcollection/Table.html:10
	qw422016.N().S(`
`)
//line views/vcollection/Table.html:11
	prms := params.Get("collection", nil, ps.Logger).Sanitize("collection")

//line views/vcollection/Table.html:11
	qw422016.N().S(`  <table>
    <thead>
      <tr>
        `)
//line views/vcollection/Table.html:15
	components.StreamTableHeaderSimple(qw422016, "collection", "id", "ID", "UUID in format (00000000-0000-0000-0000-000000000000)", prms, ps.URI, ps)
//line views/vcollection/Table.html:15
	qw422016.N().S(`
        `)
//line views/vcollection/Table.html:16
	components.StreamTableHeaderSimple(qw422016, "collection", "name", "Name", "String text", prms, ps.URI, ps)
//line views/vcollection/Table.html:16
	qw422016.N().S(`
        `)
//line views/vcollection/Table.html:17
	components.StreamTableHeaderSimple(qw422016, "collection", "created", "Created", "Date and time, in almost any format", prms, ps.URI, ps)
//line views/vcollection/Table.html:17
	qw422016.N().S(`
      </tr>
    </thead>
    <tbody>
`)
//line views/vcollection/Table.html:21
	for _, model := range models {
//line views/vcollection/Table.html:21
		qw422016.N().S(`      <tr>
        <td><a href="/collection/`)
//line views/vcollection/Table.html:23
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vcollection/Table.html:23
		qw422016.N().S(`">`)
//line views/vcollection/Table.html:23
		components.StreamDisplayUUID(qw422016, &model.ID)
//line views/vcollection/Table.html:23
		qw422016.N().S(`</a></td>
        <td><strong>`)
//line views/vcollection/Table.html:24
		qw422016.E().S(model.Name)
//line views/vcollection/Table.html:24
		qw422016.N().S(`</strong></td>
        <td>`)
//line views/vcollection/Table.html:25
		components.StreamDisplayTimestamp(qw422016, &model.Created)
//line views/vcollection/Table.html:25
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vcollection/Table.html:27
	}
//line views/vcollection/Table.html:28
	if prms.HasNextPage(len(models)+prms.Offset) || prms.HasPreviousPage() {
//line views/vcollection/Table.html:28
		qw422016.N().S(`      <tr>
        <td colspan="3">`)
//line views/vcollection/Table.html:30
		components.StreamPagination(qw422016, len(models)+prms.Offset, prms, ps.URI)
//line views/vcollection/Table.html:30
		qw422016.N().S(`</td>
      </tr>
`)
//line views/vcollection/Table.html:32
	}
//line views/vcollection/Table.html:32
	qw422016.N().S(`    </tbody>
  </table>
`)
//line views/vcollection/Table.html:35
}

//line views/vcollection/Table.html:35
func WriteTable(qq422016 qtio422016.Writer, models collection.Collections, params filter.ParamSet, as *app.State, ps *cutil.PageState) {
//line views/vcollection/Table.html:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcollection/Table.html:35
	StreamTable(qw422016, models, params, as, ps)
//line views/vcollection/Table.html:35
	qt422016.ReleaseWriter(qw422016)
//line views/vcollection/Table.html:35
}

//line views/vcollection/Table.html:35
func Table(models collection.Collections, params filter.ParamSet, as *app.State, ps *cutil.PageState) string {
//line views/vcollection/Table.html:35
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcollection/Table.html:35
	WriteTable(qb422016, models, params, as, ps)
//line views/vcollection/Table.html:35
	qs422016 := string(qb422016.B)
//line views/vcollection/Table.html:35
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcollection/Table.html:35
	return qs422016
//line views/vcollection/Table.html:35
}
