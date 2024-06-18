// Code generated by qtc from "String.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/edit/String.html:1
package edit

//line views/components/edit/String.html:1
import (
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/views/components"
)

//line views/components/edit/String.html:6
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/String.html:6
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/String.html:6
func StreamString(qw422016 *qt422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:7
	if id == "" {
//line views/components/edit/String.html:7
		qw422016.N().S(`<input name="`)
//line views/components/edit/String.html:8
		qw422016.E().S(key)
//line views/components/edit/String.html:8
		qw422016.N().S(`" value="`)
//line views/components/edit/String.html:8
		qw422016.E().S(value)
//line views/components/edit/String.html:8
		qw422016.N().S(`"`)
//line views/components/edit/String.html:8
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:8
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:9
	} else {
//line views/components/edit/String.html:9
		qw422016.N().S(`<input id="`)
//line views/components/edit/String.html:10
		qw422016.E().S(id)
//line views/components/edit/String.html:10
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:10
		qw422016.E().S(key)
//line views/components/edit/String.html:10
		qw422016.N().S(`" value="`)
//line views/components/edit/String.html:10
		qw422016.E().S(value)
//line views/components/edit/String.html:10
		qw422016.N().S(`"`)
//line views/components/edit/String.html:10
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:10
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:11
	}
//line views/components/edit/String.html:12
}

//line views/components/edit/String.html:12
func WriteString(qq422016 qtio422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:12
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:12
	StreamString(qw422016, key, id, value, placeholder...)
//line views/components/edit/String.html:12
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:12
}

//line views/components/edit/String.html:12
func String(key string, id string, value string, placeholder ...string) string {
//line views/components/edit/String.html:12
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:12
	WriteString(qb422016, key, id, value, placeholder...)
//line views/components/edit/String.html:12
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:12
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:12
	return qs422016
//line views/components/edit/String.html:12
}

//line views/components/edit/String.html:14
func StreamStringVertical(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:15
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:15
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/String.html:17
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:17
	qw422016.N().S(`<label for="`)
//line views/components/edit/String.html:18
	qw422016.E().S(id)
//line views/components/edit/String.html:18
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/String.html:18
	qw422016.E().S(title)
//line views/components/edit/String.html:18
	qw422016.N().S(`</em></label>`)
//line views/components/edit/String.html:19
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:19
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/String.html:20
	StreamString(qw422016, key, id, value, help...)
//line views/components/edit/String.html:20
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:21
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:21
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:23
}

//line views/components/edit/String.html:23
func WriteStringVertical(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:23
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:23
	StreamStringVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:23
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:23
}

//line views/components/edit/String.html:23
func StringVertical(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:23
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:23
	WriteStringVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:23
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:23
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:23
	return qs422016
//line views/components/edit/String.html:23
}

//line views/components/edit/String.html:25
func StreamStringTable(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:26
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:26
	qw422016.N().S(`<tr>`)
//line views/components/edit/String.html:28
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:28
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/String.html:29
	qw422016.E().S(id)
//line views/components/edit/String.html:29
	qw422016.N().S(`"`)
//line views/components/edit/String.html:29
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/String.html:29
	qw422016.N().S(`>`)
//line views/components/edit/String.html:29
	qw422016.E().S(title)
//line views/components/edit/String.html:29
	qw422016.N().S(`</label></th>`)
//line views/components/edit/String.html:30
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:30
	qw422016.N().S(`<td>`)
//line views/components/edit/String.html:31
	StreamString(qw422016, key, id, value, help...)
//line views/components/edit/String.html:31
	qw422016.N().S(`</td>`)
//line views/components/edit/String.html:32
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:32
	qw422016.N().S(`</tr>`)
//line views/components/edit/String.html:34
}

//line views/components/edit/String.html:34
func WriteStringTable(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:34
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:34
	StreamStringTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:34
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:34
}

//line views/components/edit/String.html:34
func StringTable(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:34
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:34
	WriteStringTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:34
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:34
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:34
	return qs422016
//line views/components/edit/String.html:34
}

//line views/components/edit/String.html:36
func StreamPassword(qw422016 *qt422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:37
	if id == "" {
//line views/components/edit/String.html:37
		qw422016.N().S(`<input name="`)
//line views/components/edit/String.html:38
		qw422016.E().S(key)
//line views/components/edit/String.html:38
		qw422016.N().S(`" type="password" value="`)
//line views/components/edit/String.html:38
		qw422016.E().S(value)
//line views/components/edit/String.html:38
		qw422016.N().S(`"`)
//line views/components/edit/String.html:38
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:38
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:39
	} else {
//line views/components/edit/String.html:39
		qw422016.N().S(`<input id="`)
//line views/components/edit/String.html:40
		qw422016.E().S(id)
//line views/components/edit/String.html:40
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:40
		qw422016.E().S(key)
//line views/components/edit/String.html:40
		qw422016.N().S(`" type="password" value="`)
//line views/components/edit/String.html:40
		qw422016.E().S(value)
//line views/components/edit/String.html:40
		qw422016.N().S(`"`)
//line views/components/edit/String.html:40
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:40
		qw422016.N().S(`/>`)
//line views/components/edit/String.html:41
	}
//line views/components/edit/String.html:42
}

//line views/components/edit/String.html:42
func WritePassword(qq422016 qtio422016.Writer, key string, id string, value string, placeholder ...string) {
//line views/components/edit/String.html:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:42
	StreamPassword(qw422016, key, id, value, placeholder...)
//line views/components/edit/String.html:42
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:42
}

//line views/components/edit/String.html:42
func Password(key string, id string, value string, placeholder ...string) string {
//line views/components/edit/String.html:42
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:42
	WritePassword(qb422016, key, id, value, placeholder...)
//line views/components/edit/String.html:42
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:42
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:42
	return qs422016
//line views/components/edit/String.html:42
}

//line views/components/edit/String.html:44
func StreamPasswordVertical(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:45
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:45
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/String.html:47
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:47
	qw422016.N().S(`<label for="`)
//line views/components/edit/String.html:48
	qw422016.E().S(id)
//line views/components/edit/String.html:48
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/String.html:48
	qw422016.E().S(title)
//line views/components/edit/String.html:48
	qw422016.N().S(`</em></label>`)
//line views/components/edit/String.html:49
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:49
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/String.html:50
	StreamPassword(qw422016, key, id, value, help...)
//line views/components/edit/String.html:50
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:51
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:51
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:53
}

//line views/components/edit/String.html:53
func WritePasswordVertical(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:53
	StreamPasswordVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:53
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:53
}

//line views/components/edit/String.html:53
func PasswordVertical(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:53
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:53
	WritePasswordVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:53
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:53
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:53
	return qs422016
//line views/components/edit/String.html:53
}

//line views/components/edit/String.html:55
func StreamPasswordTable(qw422016 *qt422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:56
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:56
	qw422016.N().S(`<tr>`)
//line views/components/edit/String.html:58
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:58
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/String.html:59
	qw422016.E().S(id)
//line views/components/edit/String.html:59
	qw422016.N().S(`"`)
//line views/components/edit/String.html:59
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/String.html:59
	qw422016.N().S(`>`)
//line views/components/edit/String.html:59
	qw422016.E().S(title)
//line views/components/edit/String.html:59
	qw422016.N().S(`</label></th>`)
//line views/components/edit/String.html:60
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:60
	qw422016.N().S(`<td>`)
//line views/components/edit/String.html:61
	StreamPassword(qw422016, key, id, value, help...)
//line views/components/edit/String.html:61
	qw422016.N().S(`</td>`)
//line views/components/edit/String.html:62
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:62
	qw422016.N().S(`</tr>`)
//line views/components/edit/String.html:64
}

//line views/components/edit/String.html:64
func WritePasswordTable(qq422016 qtio422016.Writer, key string, id string, title string, value string, indent int, help ...string) {
//line views/components/edit/String.html:64
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:64
	StreamPasswordTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:64
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:64
}

//line views/components/edit/String.html:64
func PasswordTable(key string, id string, title string, value string, indent int, help ...string) string {
//line views/components/edit/String.html:64
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:64
	WritePasswordTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/String.html:64
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:64
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:64
	return qs422016
//line views/components/edit/String.html:64
}

//line views/components/edit/String.html:66
func StreamTextarea(qw422016 *qt422016.Writer, key string, id string, rows int, value string, placeholder ...string) {
//line views/components/edit/String.html:67
	if id == "" {
//line views/components/edit/String.html:67
		qw422016.N().S(`<textarea rows="`)
//line views/components/edit/String.html:68
		qw422016.N().D(rows)
//line views/components/edit/String.html:68
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:68
		qw422016.E().S(key)
//line views/components/edit/String.html:68
		qw422016.N().S(`"`)
//line views/components/edit/String.html:68
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:68
		qw422016.N().S(`>`)
//line views/components/edit/String.html:68
		qw422016.E().S(value)
//line views/components/edit/String.html:68
		qw422016.N().S(`</textarea>`)
//line views/components/edit/String.html:69
	} else {
//line views/components/edit/String.html:69
		qw422016.N().S(`<textarea rows="`)
//line views/components/edit/String.html:70
		qw422016.N().D(rows)
//line views/components/edit/String.html:70
		qw422016.N().S(`" id="`)
//line views/components/edit/String.html:70
		qw422016.E().S(id)
//line views/components/edit/String.html:70
		qw422016.N().S(`" name="`)
//line views/components/edit/String.html:70
		qw422016.E().S(key)
//line views/components/edit/String.html:70
		qw422016.N().S(`"`)
//line views/components/edit/String.html:70
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/String.html:70
		qw422016.N().S(`>`)
//line views/components/edit/String.html:70
		qw422016.E().S(value)
//line views/components/edit/String.html:70
		qw422016.N().S(`</textarea>`)
//line views/components/edit/String.html:71
	}
//line views/components/edit/String.html:72
}

//line views/components/edit/String.html:72
func WriteTextarea(qq422016 qtio422016.Writer, key string, id string, rows int, value string, placeholder ...string) {
//line views/components/edit/String.html:72
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:72
	StreamTextarea(qw422016, key, id, rows, value, placeholder...)
//line views/components/edit/String.html:72
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:72
}

//line views/components/edit/String.html:72
func Textarea(key string, id string, rows int, value string, placeholder ...string) string {
//line views/components/edit/String.html:72
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:72
	WriteTextarea(qb422016, key, id, rows, value, placeholder...)
//line views/components/edit/String.html:72
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:72
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:72
	return qs422016
//line views/components/edit/String.html:72
}

//line views/components/edit/String.html:74
func StreamTextareaVertical(qw422016 *qt422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:75
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:75
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/String.html:77
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:77
	qw422016.N().S(`<label for="`)
//line views/components/edit/String.html:78
	qw422016.E().S(id)
//line views/components/edit/String.html:78
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/String.html:78
	qw422016.E().S(title)
//line views/components/edit/String.html:78
	qw422016.N().S(`</em></label>`)
//line views/components/edit/String.html:79
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:79
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/String.html:80
	StreamTextarea(qw422016, key, id, rows, value, help...)
//line views/components/edit/String.html:80
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:81
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:81
	qw422016.N().S(`</div>`)
//line views/components/edit/String.html:83
}

//line views/components/edit/String.html:83
func WriteTextareaVertical(qq422016 qtio422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:83
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:83
	StreamTextareaVertical(qw422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:83
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:83
}

//line views/components/edit/String.html:83
func TextareaVertical(key string, id string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/edit/String.html:83
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:83
	WriteTextareaVertical(qb422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:83
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:83
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:83
	return qs422016
//line views/components/edit/String.html:83
}

//line views/components/edit/String.html:85
func StreamTextareaTable(qw422016 *qt422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:86
	id = cutil.CleanID(key, id)

//line views/components/edit/String.html:86
	qw422016.N().S(`<tr>`)
//line views/components/edit/String.html:88
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:88
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/String.html:89
	qw422016.E().S(id)
//line views/components/edit/String.html:89
	qw422016.N().S(`"`)
//line views/components/edit/String.html:89
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/String.html:89
	qw422016.N().S(`>`)
//line views/components/edit/String.html:89
	qw422016.E().S(title)
//line views/components/edit/String.html:89
	qw422016.N().S(`</label></th>`)
//line views/components/edit/String.html:90
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/String.html:90
	qw422016.N().S(`<td>`)
//line views/components/edit/String.html:91
	StreamTextarea(qw422016, key, id, rows, value, help...)
//line views/components/edit/String.html:91
	qw422016.N().S(`</td>`)
//line views/components/edit/String.html:92
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/String.html:92
	qw422016.N().S(`</tr>`)
//line views/components/edit/String.html:94
}

//line views/components/edit/String.html:94
func WriteTextareaTable(qq422016 qtio422016.Writer, key string, id string, title string, rows int, value string, indent int, help ...string) {
//line views/components/edit/String.html:94
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/String.html:94
	StreamTextareaTable(qw422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:94
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/String.html:94
}

//line views/components/edit/String.html:94
func TextareaTable(key string, id string, title string, rows int, value string, indent int, help ...string) string {
//line views/components/edit/String.html:94
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/String.html:94
	WriteTextareaTable(qb422016, key, id, title, rows, value, indent, help...)
//line views/components/edit/String.html:94
	qs422016 := string(qb422016.B)
//line views/components/edit/String.html:94
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/String.html:94
	return qs422016
//line views/components/edit/String.html:94
}
