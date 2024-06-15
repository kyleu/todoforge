// Code generated by qtc from "Result.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vsearch/Result.html:2
package vsearch

//line views/vsearch/Result.html:2
import (
	"strings"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/search"
	"github.com/kyleu/todoforge/app/lib/search/result"
	"github.com/kyleu/todoforge/views/components"
)

//line views/vsearch/Result.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vsearch/Result.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vsearch/Result.html:12
func StreamResult(qw422016 *qt422016.Writer, res *result.Result, params *search.Params, as *app.State, ps *cutil.PageState) {
//line views/vsearch/Result.html:12
	qw422016.N().S(`
  <div class="card">
    <div class="right">
`)
//line views/vsearch/Result.html:15
	if res.Data == nil {
//line views/vsearch/Result.html:15
		qw422016.N().S(`      <em>`)
//line views/vsearch/Result.html:16
		qw422016.E().S(res.Type)
//line views/vsearch/Result.html:16
		qw422016.N().S(`</em>
`)
//line views/vsearch/Result.html:17
	} else if res.ID != "" {
//line views/vsearch/Result.html:17
		qw422016.N().S(`      <a href="#modal-`)
//line views/vsearch/Result.html:18
		qw422016.E().S(res.Type)
//line views/vsearch/Result.html:18
		qw422016.N().S(`-`)
//line views/vsearch/Result.html:18
		qw422016.E().S(res.ID)
//line views/vsearch/Result.html:18
		qw422016.N().S(`"><button type="button">`)
//line views/vsearch/Result.html:18
		qw422016.E().S(res.Type)
//line views/vsearch/Result.html:18
		qw422016.N().S(`</button></a>
      `)
//line views/vsearch/Result.html:19
		components.StreamJSONModal(qw422016, res.Type+"-"+res.ID, res.Type, res.Data, 3)
//line views/vsearch/Result.html:19
		qw422016.N().S(`
`)
//line views/vsearch/Result.html:20
	}
//line views/vsearch/Result.html:20
	qw422016.N().S(`    </div>
    <h3>`)
//line views/vsearch/Result.html:22
	if res.Icon != "" {
//line views/vsearch/Result.html:22
		qw422016.N().S(`<a href="`)
//line views/vsearch/Result.html:22
		qw422016.E().S(res.URL)
//line views/vsearch/Result.html:22
		qw422016.N().S(`">`)
//line views/vsearch/Result.html:22
		components.StreamSVGInline(qw422016, res.Icon, 18, ps)
//line views/vsearch/Result.html:22
		qw422016.N().S(`</a>`)
//line views/vsearch/Result.html:22
		qw422016.N().S(` `)
//line views/vsearch/Result.html:22
	}
//line views/vsearch/Result.html:22
	qw422016.N().S(`<a href="`)
//line views/vsearch/Result.html:22
	qw422016.E().S(res.URL)
//line views/vsearch/Result.html:22
	qw422016.N().S(`">`)
//line views/vsearch/Result.html:22
	if res.Title == "" {
//line views/vsearch/Result.html:22
		qw422016.E().S(res.URL)
//line views/vsearch/Result.html:22
	} else {
//line views/vsearch/Result.html:22
		qw422016.E().S(res.Title)
//line views/vsearch/Result.html:22
	}
//line views/vsearch/Result.html:22
	qw422016.N().S(`</a></h3>
    <p>`)
//line views/vsearch/Result.html:23
	StreamMatch(qw422016, params, res.Matches)
//line views/vsearch/Result.html:23
	qw422016.N().S(`</p>
`)
//line views/vsearch/Result.html:24
	if res.HTML != "" {
//line views/vsearch/Result.html:24
		qw422016.N().S(`    <div>`)
//line views/vsearch/Result.html:25
		qw422016.N().S(res.HTML)
//line views/vsearch/Result.html:25
		qw422016.N().S(`</div>
`)
//line views/vsearch/Result.html:26
	}
//line views/vsearch/Result.html:26
	qw422016.N().S(`  </div>
`)
//line views/vsearch/Result.html:28
}

//line views/vsearch/Result.html:28
func WriteResult(qq422016 qtio422016.Writer, res *result.Result, params *search.Params, as *app.State, ps *cutil.PageState) {
//line views/vsearch/Result.html:28
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsearch/Result.html:28
	StreamResult(qw422016, res, params, as, ps)
//line views/vsearch/Result.html:28
	qt422016.ReleaseWriter(qw422016)
//line views/vsearch/Result.html:28
}

//line views/vsearch/Result.html:28
func Result(res *result.Result, params *search.Params, as *app.State, ps *cutil.PageState) string {
//line views/vsearch/Result.html:28
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsearch/Result.html:28
	WriteResult(qb422016, res, params, as, ps)
//line views/vsearch/Result.html:28
	qs422016 := string(qb422016.B)
//line views/vsearch/Result.html:28
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsearch/Result.html:28
	return qs422016
//line views/vsearch/Result.html:28
}

//line views/vsearch/Result.html:30
func StreamMatch(qw422016 *qt422016.Writer, params *search.Params, matches result.Matches) {
//line views/vsearch/Result.html:30
	qw422016.N().S(`
  <ul>
`)
//line views/vsearch/Result.html:32
	if len(matches) == 0 {
//line views/vsearch/Result.html:32
		qw422016.N().S(`    <li><em>`)
//line views/vsearch/Result.html:33
		qw422016.E().S(params.Q)
//line views/vsearch/Result.html:33
		qw422016.N().S(`</em></li>
`)
//line views/vsearch/Result.html:34
	}
//line views/vsearch/Result.html:35
	for _, m := range matches {
//line views/vsearch/Result.html:35
		qw422016.N().S(`    <li>
      <em>`)
//line views/vsearch/Result.html:37
		qw422016.E().S(m.Key)
//line views/vsearch/Result.html:37
		qw422016.N().S(`</em>:
`)
//line views/vsearch/Result.html:38
		split := m.ValueSplit(params.Q)

//line views/vsearch/Result.html:39
		for _, x := range split {
//line views/vsearch/Result.html:39
			if strings.EqualFold(x, params.Q) {
//line views/vsearch/Result.html:39
				qw422016.N().S(`<strong>`)
//line views/vsearch/Result.html:39
				qw422016.E().S(x)
//line views/vsearch/Result.html:39
				qw422016.N().S(`</strong>`)
//line views/vsearch/Result.html:39
			} else {
//line views/vsearch/Result.html:39
				qw422016.E().S(x)
//line views/vsearch/Result.html:39
			}
//line views/vsearch/Result.html:39
		}
//line views/vsearch/Result.html:39
		qw422016.N().S(`    </li>
`)
//line views/vsearch/Result.html:41
	}
//line views/vsearch/Result.html:41
	qw422016.N().S(`  </ul>
`)
//line views/vsearch/Result.html:43
}

//line views/vsearch/Result.html:43
func WriteMatch(qq422016 qtio422016.Writer, params *search.Params, matches result.Matches) {
//line views/vsearch/Result.html:43
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vsearch/Result.html:43
	StreamMatch(qw422016, params, matches)
//line views/vsearch/Result.html:43
	qt422016.ReleaseWriter(qw422016)
//line views/vsearch/Result.html:43
}

//line views/vsearch/Result.html:43
func Match(params *search.Params, matches result.Matches) string {
//line views/vsearch/Result.html:43
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vsearch/Result.html:43
	WriteMatch(qb422016, params, matches)
//line views/vsearch/Result.html:43
	qs422016 := string(qb422016.B)
//line views/vsearch/Result.html:43
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vsearch/Result.html:43
	return qs422016
//line views/vsearch/Result.html:43
}
