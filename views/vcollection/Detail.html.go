// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vcollection/Detail.html:1
package vcollection

//line views/vcollection/Detail.html:1
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/collection"
	"github.com/kyleu/todoforge/app/collection/item"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/components/view"
	"github.com/kyleu/todoforge/views/layout"
	"github.com/kyleu/todoforge/views/vcollection/vitem"
)

//line views/vcollection/Detail.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vcollection/Detail.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vcollection/Detail.html:14
type Detail struct {
	layout.Basic
	Model                  *collection.Collection
	Params                 filter.ParamSet
	RelItemsByCollectionID item.Items
}

//line views/vcollection/Detail.html:21
func (p *Detail) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/Detail.html:21
	qw422016.N().S(`
  <div class="card">
    <div class="right">
      <a href="#modal-collection"><button type="button">`)
//line views/vcollection/Detail.html:24
	components.StreamSVGButton(qw422016, "file", ps)
//line views/vcollection/Detail.html:24
	qw422016.N().S(` JSON</button></a>
      <a href="`)
//line views/vcollection/Detail.html:25
	qw422016.E().S(p.Model.WebPath())
//line views/vcollection/Detail.html:25
	qw422016.N().S(`/edit"><button>`)
//line views/vcollection/Detail.html:25
	components.StreamSVGButton(qw422016, "edit", ps)
//line views/vcollection/Detail.html:25
	qw422016.N().S(` Edit</button></a>
    </div>
    <h3>`)
//line views/vcollection/Detail.html:27
	components.StreamSVGIcon(qw422016, `archive`, ps)
//line views/vcollection/Detail.html:27
	qw422016.N().S(` `)
//line views/vcollection/Detail.html:27
	qw422016.E().S(p.Model.TitleString())
//line views/vcollection/Detail.html:27
	qw422016.N().S(`</h3>
    <div><a href="/collection"><em>Collection</em></a></div>
    <div class="mt overflow full-width">
      <table>
        <tbody>
          <tr>
            <th class="shrink" title="UUID in format (00000000-0000-0000-0000-000000000000)">ID</th>
            <td>`)
//line views/vcollection/Detail.html:34
	view.StreamUUID(qw422016, &p.Model.ID)
//line views/vcollection/Detail.html:34
	qw422016.N().S(`</td>
          </tr>
          <tr>
            <th class="shrink" title="String text">Name</th>
            <td><strong>`)
//line views/vcollection/Detail.html:38
	view.StreamString(qw422016, p.Model.Name)
//line views/vcollection/Detail.html:38
	qw422016.N().S(`</strong></td>
          </tr>
          <tr>
            <th class="shrink" title="Date and time, in almost any format">Created</th>
            <td>`)
//line views/vcollection/Detail.html:42
	view.StreamTimestamp(qw422016, &p.Model.Created)
//line views/vcollection/Detail.html:42
	qw422016.N().S(`</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
`)
//line views/vcollection/Detail.html:50
	relationHelper := collection.Collections{p.Model}

//line views/vcollection/Detail.html:50
	qw422016.N().S(`  <div class="card">
    <h3 class="mb">Relations</h3>
    <ul class="accordion">
      <li>
        <input id="accordion-ItemsByCollectionID" type="checkbox" hidden="hidden"`)
//line views/vcollection/Detail.html:55
	if p.Params.Specifies(`item`) {
//line views/vcollection/Detail.html:55
		qw422016.N().S(` checked="checked"`)
//line views/vcollection/Detail.html:55
	}
//line views/vcollection/Detail.html:55
	qw422016.N().S(` />
        <label for="accordion-ItemsByCollectionID">
          `)
//line views/vcollection/Detail.html:57
	components.StreamExpandCollapse(qw422016, 3, ps)
//line views/vcollection/Detail.html:57
	qw422016.N().S(`
          `)
//line views/vcollection/Detail.html:58
	components.StreamSVGInline(qw422016, `file`, 16, ps)
//line views/vcollection/Detail.html:58
	qw422016.N().S(`
          `)
//line views/vcollection/Detail.html:59
	qw422016.E().S(util.StringPlural(len(p.RelItemsByCollectionID), "Item"))
//line views/vcollection/Detail.html:59
	qw422016.N().S(` by [Collection ID]
        </label>
        <div class="bd"><div><div>
`)
//line views/vcollection/Detail.html:62
	if len(p.RelItemsByCollectionID) == 0 {
//line views/vcollection/Detail.html:62
		qw422016.N().S(`          <em>no related Items</em>
`)
//line views/vcollection/Detail.html:64
	} else {
//line views/vcollection/Detail.html:64
		qw422016.N().S(`          <div class="overflow clear">
            `)
//line views/vcollection/Detail.html:66
		vitem.StreamTable(qw422016, p.RelItemsByCollectionID, relationHelper, p.Params, as, ps)
//line views/vcollection/Detail.html:66
		qw422016.N().S(`
          </div>
`)
//line views/vcollection/Detail.html:68
	}
//line views/vcollection/Detail.html:68
	qw422016.N().S(`        </div></div></div>
      </li>
    </ul>
  </div>
  `)
//line views/vcollection/Detail.html:73
	components.StreamJSONModal(qw422016, "collection", "Collection JSON", p.Model, 1)
//line views/vcollection/Detail.html:73
	qw422016.N().S(`
`)
//line views/vcollection/Detail.html:74
}

//line views/vcollection/Detail.html:74
func (p *Detail) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vcollection/Detail.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vcollection/Detail.html:74
	p.StreamBody(qw422016, as, ps)
//line views/vcollection/Detail.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/vcollection/Detail.html:74
}

//line views/vcollection/Detail.html:74
func (p *Detail) Body(as *app.State, ps *cutil.PageState) string {
//line views/vcollection/Detail.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vcollection/Detail.html:74
	p.WriteBody(qb422016, as, ps)
//line views/vcollection/Detail.html:74
	qs422016 := string(qb422016.B)
//line views/vcollection/Detail.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vcollection/Detail.html:74
	return qs422016
//line views/vcollection/Detail.html:74
}
