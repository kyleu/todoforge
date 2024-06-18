// Code generated by qtc from "Int.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/edit/Int.html:1
package edit

//line views/components/edit/Int.html:1
import (
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/views/components"
)

//line views/components/edit/Int.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Int.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Int.html:6
func StreamInt(qw422016 *qt422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/edit/Int.html:7
	if id == "" {
//line views/components/edit/Int.html:7
		qw422016.N().S(`<input name="`)
//line views/components/edit/Int.html:8
		qw422016.E().S(key)
//line views/components/edit/Int.html:8
		qw422016.N().S(`" type="number" value="`)
//line views/components/edit/Int.html:8
		qw422016.E().V(value)
//line views/components/edit/Int.html:8
		qw422016.N().S(`"`)
//line views/components/edit/Int.html:8
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Int.html:8
		qw422016.N().S(`/>`)
//line views/components/edit/Int.html:9
	} else {
//line views/components/edit/Int.html:9
		qw422016.N().S(`<input id="`)
//line views/components/edit/Int.html:10
		qw422016.E().S(id)
//line views/components/edit/Int.html:10
		qw422016.N().S(`" name="`)
//line views/components/edit/Int.html:10
		qw422016.E().S(key)
//line views/components/edit/Int.html:10
		qw422016.N().S(`" type="number" value="`)
//line views/components/edit/Int.html:10
		qw422016.E().V(value)
//line views/components/edit/Int.html:10
		qw422016.N().S(`"`)
//line views/components/edit/Int.html:10
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Int.html:10
		qw422016.N().S(`/>`)
//line views/components/edit/Int.html:11
	}
//line views/components/edit/Int.html:12
}

//line views/components/edit/Int.html:12
func WriteInt(qq422016 qtio422016.Writer, key string, id string, value any, placeholder ...string) {
//line views/components/edit/Int.html:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Int.html:12
	StreamInt(qw422016, key, id, value, placeholder...)
//line views/components/edit/Int.html:12
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Int.html:12
}

//line views/components/edit/Int.html:12
func Int(key string, id string, value any, placeholder ...string) string {
//line views/components/edit/Int.html:12
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Int.html:12
	WriteInt(qb422016, key, id, value, placeholder...)
//line views/components/edit/Int.html:12
	qs422016 := string(qb422016.B)
//line views/components/edit/Int.html:12
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Int.html:12
	return qs422016
//line views/components/edit/Int.html:12
}

//line views/components/edit/Int.html:14
func StreamIntVertical(qw422016 *qt422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/edit/Int.html:15
	id = cutil.CleanID(key, id)

//line views/components/edit/Int.html:15
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Int.html:17
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Int.html:17
	qw422016.N().S(`<label for="`)
//line views/components/edit/Int.html:18
	qw422016.E().S(id)
//line views/components/edit/Int.html:18
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/Int.html:18
	qw422016.E().S(title)
//line views/components/edit/Int.html:18
	qw422016.N().S(`</em></label>`)
//line views/components/edit/Int.html:19
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Int.html:19
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Int.html:20
	StreamInt(qw422016, key, id, value, help...)
//line views/components/edit/Int.html:20
	qw422016.N().S(`</div>`)
//line views/components/edit/Int.html:21
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Int.html:21
	qw422016.N().S(`</div>`)
//line views/components/edit/Int.html:23
}

//line views/components/edit/Int.html:23
func WriteIntVertical(qq422016 qtio422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/edit/Int.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Int.html:23
	StreamIntVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Int.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Int.html:23
}

//line views/components/edit/Int.html:23
func IntVertical(key string, id string, title string, value int, indent int, help ...string) string {
//line views/components/edit/Int.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Int.html:23
	WriteIntVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Int.html:23
	qs422016 := string(qb422016.B)
//line views/components/edit/Int.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Int.html:23
	return qs422016
//line views/components/edit/Int.html:23
}

//line views/components/edit/Int.html:25
func StreamIntTable(qw422016 *qt422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/edit/Int.html:26
	id = cutil.CleanID(key, id)

//line views/components/edit/Int.html:26
	qw422016.N().S(`<tr>`)
//line views/components/edit/Int.html:28
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Int.html:28
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Int.html:29
	qw422016.E().S(id)
//line views/components/edit/Int.html:29
	qw422016.N().S(`"`)
//line views/components/edit/Int.html:29
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Int.html:29
	qw422016.N().S(`>`)
//line views/components/edit/Int.html:29
	qw422016.E().S(title)
//line views/components/edit/Int.html:29
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Int.html:30
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Int.html:30
	qw422016.N().S(`<td>`)
//line views/components/edit/Int.html:31
	StreamInt(qw422016, key, id, value, help...)
//line views/components/edit/Int.html:31
	qw422016.N().S(`</td>`)
//line views/components/edit/Int.html:32
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Int.html:32
	qw422016.N().S(`</tr>`)
//line views/components/edit/Int.html:34
}

//line views/components/edit/Int.html:34
func WriteIntTable(qq422016 qtio422016.Writer, key string, id string, title string, value int, indent int, help ...string) {
//line views/components/edit/Int.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Int.html:34
	StreamIntTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Int.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Int.html:34
}

//line views/components/edit/Int.html:34
func IntTable(key string, id string, title string, value int, indent int, help ...string) string {
//line views/components/edit/Int.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Int.html:34
	WriteIntTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Int.html:34
	qs422016 := string(qb422016.B)
//line views/components/edit/Int.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Int.html:34
	return qs422016
//line views/components/edit/Int.html:34
}
