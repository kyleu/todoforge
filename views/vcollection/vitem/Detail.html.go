// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vcollection/vitem/Detail.html:1
package vitem

//line views/vcollection/vitem/Detail.html:1
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/components/view"
	"github.com/kyleu/todoforge/views/layout"
)

//line views/vcollection/vitem/Detail.html:11
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcollection/vitem/Detail.html:11
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcollection/vitem/Detail.html:11
type Detail struct {
	layout.Basic
	Model                    *item.Item
	CollectionByCollectionID *collection.Collection
}

//line views/vcollection/vitem/Detail.html:17
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/vitem/Detail.html:17
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-item"><button type="button">`)
//line views/vcollection/vitem/Detail.html:20
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vcollection/vitem/Detail.html:20
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vcollection/vitem/Detail.html:21
	qw422016.E().S(p.Model.WebPath())
//line views/vcollection/vitem/Detail.html:21
	qw422016.N().S(`/edit"><button>`)
//line views/vcollection/vitem/Detail.html:21
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vcollection/vitem/Detail.html:21
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vcollection/vitem/Detail.html:23
	components.StreamSVGIcon(qw422016, `file`, ps)
//line views/vcollection/vitem/Detail.html:23
	qw422016.N().S(` `)
//line views/vcollection/vitem/Detail.html:23
	qw422016.E().S(p.Model.TitleString())
//line views/vcollection/vitem/Detail.html:23
	qw422016.N().S(`</h3>
    <div><a href="/collection/item"><em>Item</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vcollection/vitem/Detail.html:30
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vcollection/vitem/Detail.html:30
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">Collection ID</th>
            <td class="nowrap">
              `)
//line views/vcollection/vitem/Detail.html:35
	view.StreamUUID(qw422016, &p.Model.CollectionID)
//line views/vcollection/vitem/Detail.html:35
	if p.CollectionByCollectionID != nil {
//line views/vcollection/vitem/Detail.html:35
		qw422016.N().S(` (`)
//line views/vcollection/vitem/Detail.html:35
		qw422016.E().S(p.CollectionByCollectionID.TitleString())
//line views/vcollection/vitem/Detail.html:35
		qw422016.N().S(`)`)
//line views/vcollection/vitem/Detail.html:35
	}
//line views/vcollection/vitem/Detail.html:35
	qw422016.N().S(`
              <a title="Collection" href="`)
//line views/vcollection/vitem/Detail.html:36
	qw422016.E().S(`/collection` + `/` + p.Model.CollectionID.String())
//line views/vcollection/vitem/Detail.html:36
	qw422016.N().S(`">`)
//line views/vcollection/vitem/Detail.html:36
	components.StreamSVGLink(qw422016, `archive`, ps)
//line views/vcollection/vitem/Detail.html:36
	qw422016.N().S(`</a>
            </td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Name</th>
            <td><strong>`)
//line views/vcollection/vitem/Detail.html:41
	view.StreamString(qw422016, p.Model.Name)
//line views/vcollection/vitem/Detail.html:41
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vcollection/vitem/Detail.html:45
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vcollection/vitem/Detail.html:45
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vcollection/vitem/Detail.html:52
	qw422016.N().S(`  `)
//line views/vcollection/vitem/Detail.html:53
	components.StreamJSONModal(qw422016, "item", "Item JSON", p.Model, 1)
//line views/vcollection/vitem/Detail.html:53
	qw422016.N().S(`
`)
//line views/vcollection/vitem/Detail.html:54
}

//line views/vcollection/vitem/Detail.html:54
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/vitem/Detail.html:54
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcollection/vitem/Detail.html:54
	p.StreamBody(qw422016, as, ps)
//line views/vcollection/vitem/Detail.html:54
	qt422016.ReleaseWriter(qw422016)
//line views/vcollection/vitem/Detail.html:54
}

//line views/vcollection/vitem/Detail.html:54
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcollection/vitem/Detail.html:54
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcollection/vitem/Detail.html:54
	p.WriteBody(qb422016, as, ps)
//line views/vcollection/vitem/Detail.html:54
	qs422016 := string(qb422016.B)
//line views/vcollection/vitem/Detail.html:54
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcollection/vitem/Detail.html:54
	return qs422016
//line views/vcollection/vitem/Detail.html:54
}
