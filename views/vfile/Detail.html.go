// Code generated by qtc from "Detail.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/vfile/Detail.html:2
package vfile

//line views/vfile/Detail.html:2
import (
	"encoding/base64"
	"path/filepath"
	"unicode/utf8"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/filesystem"
	"github.com/kyleu/todoforge/app/util"
)

//line views/vfile/Detail.html:13
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vfile/Detail.html:13
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vfile/Detail.html:13
func StreamDetail(qw422016 *qt422016.Writer, path []string, b []byte, urlPrefix string, additionalLinks map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Detail.html:13
	qw422016.N().S(`
`)
//line views/vfile/Detail.html:14
	if additionalLinks != nil && len(additionalLinks) > 0 {
//line views/vfile/Detail.html:14
		qw422016.N().S(`  <div class="right">
`)
//line views/vfile/Detail.html:16
		for k, v := range additionalLinks {
//line views/vfile/Detail.html:16
			qw422016.N().S(`    <a href="`)
//line views/vfile/Detail.html:17
			qw422016.E().S(v)
//line views/vfile/Detail.html:17
			qw422016.N().S(`"><button>`)
//line views/vfile/Detail.html:17
			qw422016.E().S(k)
//line views/vfile/Detail.html:17
			qw422016.N().S(`</button></a>
`)
//line views/vfile/Detail.html:18
		}
//line views/vfile/Detail.html:18
		qw422016.N().S(`  </div>
`)
//line views/vfile/Detail.html:20
	}
//line views/vfile/Detail.html:20
	qw422016.N().S(`  <h3>
`)
//line views/vfile/Detail.html:22
	for idx, p := range path {
//line views/vfile/Detail.html:22
		qw422016.N().S(`/<a href="`)
//line views/vfile/Detail.html:22
		qw422016.E().S(urlPrefix)
//line views/vfile/Detail.html:22
		qw422016.N().S(`/`)
//line views/vfile/Detail.html:22
		qw422016.E().S(filepath.Join(path[:idx+1]...))
//line views/vfile/Detail.html:22
		qw422016.N().S(`">`)
//line views/vfile/Detail.html:22
		qw422016.E().S(p)
//line views/vfile/Detail.html:22
		qw422016.N().S(`</a>`)
//line views/vfile/Detail.html:22
	}
//line views/vfile/Detail.html:22
	qw422016.N().S(`    <em>(`)
//line views/vfile/Detail.html:23
	qw422016.E().S(util.ByteSizeSI(int64(len(b))))
//line views/vfile/Detail.html:23
	qw422016.N().S(`)</em>
  </h3>
  <div class="mt">
`)
//line views/vfile/Detail.html:26
	if len(b) > (1024 * 128) {
//line views/vfile/Detail.html:26
		qw422016.N().S(`    <em>File is `)
//line views/vfile/Detail.html:27
		qw422016.N().D(len(b))
//line views/vfile/Detail.html:27
		qw422016.N().S(` bytes, which is too large for the file viewer</em>
`)
//line views/vfile/Detail.html:28
	} else if utf8.Valid(b) {
//line views/vfile/Detail.html:29
		out, _ := cutil.FormatFilename(string(b), path[len(path)-1])

//line views/vfile/Detail.html:29
		qw422016.N().S(`    `)
//line views/vfile/Detail.html:30
		qw422016.N().S(out)
//line views/vfile/Detail.html:30
		qw422016.N().S(`
`)
//line views/vfile/Detail.html:31
	} else {
//line views/vfile/Detail.html:31
		qw422016.N().S(`
`)
//line views/vfile/Detail.html:33
		if imgType := filesystem.ImageType(path...); imgType != "" {
//line views/vfile/Detail.html:33
			qw422016.N().S(`    <img alt="Image of type [`)
//line views/vfile/Detail.html:34
			qw422016.E().S(imgType)
//line views/vfile/Detail.html:34
			qw422016.N().S(`]" src="data:image/`)
//line views/vfile/Detail.html:34
			qw422016.E().S(imgType)
//line views/vfile/Detail.html:34
			qw422016.N().S(`;base64,`)
//line views/vfile/Detail.html:34
			qw422016.E().S(base64.StdEncoding.EncodeToString(b))
//line views/vfile/Detail.html:34
			qw422016.N().S(`" />
    <hr />
`)
//line views/vfile/Detail.html:36
		}
//line views/vfile/Detail.html:36
		qw422016.N().S(`
`)
//line views/vfile/Detail.html:38
		exif, err := filesystem.ExifExtract(b)

//line views/vfile/Detail.html:39
		if err == nil {
//line views/vfile/Detail.html:39
			qw422016.N().S(`    <div class="overflow full-width">
      <table>
        <thead>
          <tr>
            <th>EXIF Name</th>
            <th>Value</th>
          </tr>
        </thead>
        <tbody>
`)
//line views/vfile/Detail.html:49
			for k, v := range exif {
//line views/vfile/Detail.html:49
				qw422016.N().S(`          <tr>
            <td>`)
//line views/vfile/Detail.html:51
				qw422016.E().S(k)
//line views/vfile/Detail.html:51
				qw422016.N().S(`</td>
            <td>`)
//line views/vfile/Detail.html:52
				qw422016.E().V(v)
//line views/vfile/Detail.html:52
				qw422016.N().S(`</td>
          </tr>
`)
//line views/vfile/Detail.html:54
			}
//line views/vfile/Detail.html:54
			qw422016.N().S(`        </tbody>
      </table>
    </div>
`)
//line views/vfile/Detail.html:58
		} else {
//line views/vfile/Detail.html:58
			qw422016.N().S(`    <em>File is binary and contains no exif header</em>
`)
//line views/vfile/Detail.html:60
		}
//line views/vfile/Detail.html:61
	}
//line views/vfile/Detail.html:61
	qw422016.N().S(`  </div>
`)
//line views/vfile/Detail.html:63
}

//line views/vfile/Detail.html:63
func WriteDetail(qq422016 qtio422016.Writer, path []string, b []byte, urlPrefix string, additionalLinks map[string]string, as *app.State, ps *cutil.PageState) {
//line views/vfile/Detail.html:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vfile/Detail.html:63
	StreamDetail(qw422016, path, b, urlPrefix, additionalLinks, as, ps)
//line views/vfile/Detail.html:63
	qt422016.ReleaseWriter(qw422016)
//line views/vfile/Detail.html:63
}

//line views/vfile/Detail.html:63
func Detail(path []string, b []byte, urlPrefix string, additionalLinks map[string]string, as *app.State, ps *cutil.PageState) string {
//line views/vfile/Detail.html:63
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vfile/Detail.html:63
	WriteDetail(qb422016, path, b, urlPrefix, additionalLinks, as, ps)
//line views/vfile/Detail.html:63
	qs422016 := string(qb422016.B)
//line views/vfile/Detail.html:63
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vfile/Detail.html:63
	return qs422016
//line views/vfile/Detail.html:63
}
