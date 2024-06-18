// Code generated by qtc from "Timestamp.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/components/edit/Timestamp.html:1
package edit

//line views/components/edit/Timestamp.html:1
import (
	"time"

	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/components"
)

//line views/components/edit/Timestamp.html:9
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/edit/Timestamp.html:9
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/edit/Timestamp.html:9
func StreamTimestamp(qw422016 *qt422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/edit/Timestamp.html:10
	if id == "" {
//line views/components/edit/Timestamp.html:10
		qw422016.N().S(`<input name="`)
//line views/components/edit/Timestamp.html:11
		qw422016.E().S(key)
//line views/components/edit/Timestamp.html:11
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/edit/Timestamp.html:11
		qw422016.E().S(util.TimeToFull(value))
//line views/components/edit/Timestamp.html:11
		qw422016.N().S(`"`)
//line views/components/edit/Timestamp.html:11
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Timestamp.html:11
		qw422016.N().S(`/>`)
//line views/components/edit/Timestamp.html:12
	} else {
//line views/components/edit/Timestamp.html:12
		qw422016.N().S(`<input id="`)
//line views/components/edit/Timestamp.html:13
		qw422016.E().S(id)
//line views/components/edit/Timestamp.html:13
		qw422016.N().S(`" name="`)
//line views/components/edit/Timestamp.html:13
		qw422016.E().S(key)
//line views/components/edit/Timestamp.html:13
		qw422016.N().S(`" type="datetime-local" value="`)
//line views/components/edit/Timestamp.html:13
		qw422016.E().S(util.TimeToFull(value))
//line views/components/edit/Timestamp.html:13
		qw422016.N().S(`"`)
//line views/components/edit/Timestamp.html:13
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Timestamp.html:13
		qw422016.N().S(`/>`)
//line views/components/edit/Timestamp.html:14
	}
//line views/components/edit/Timestamp.html:15
}

//line views/components/edit/Timestamp.html:15
func WriteTimestamp(qq422016 qtio422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/edit/Timestamp.html:15
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Timestamp.html:15
	StreamTimestamp(qw422016, key, id, value, placeholder...)
//line views/components/edit/Timestamp.html:15
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Timestamp.html:15
}

//line views/components/edit/Timestamp.html:15
func Timestamp(key string, id string, value *time.Time, placeholder ...string) string {
//line views/components/edit/Timestamp.html:15
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Timestamp.html:15
	WriteTimestamp(qb422016, key, id, value, placeholder...)
//line views/components/edit/Timestamp.html:15
	qs422016 := string(qb422016.B)
//line views/components/edit/Timestamp.html:15
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Timestamp.html:15
	return qs422016
//line views/components/edit/Timestamp.html:15
}

//line views/components/edit/Timestamp.html:17
func StreamTimestampVertical(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:18
	id = cutil.CleanID(key, id)

//line views/components/edit/Timestamp.html:18
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Timestamp.html:20
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:20
	qw422016.N().S(`<label for="`)
//line views/components/edit/Timestamp.html:21
	qw422016.E().S(id)
//line views/components/edit/Timestamp.html:21
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/Timestamp.html:21
	qw422016.E().S(title)
//line views/components/edit/Timestamp.html:21
	qw422016.N().S(`</em></label>`)
//line views/components/edit/Timestamp.html:22
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:22
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Timestamp.html:23
	StreamTimestamp(qw422016, key, id, value, help...)
//line views/components/edit/Timestamp.html:23
	qw422016.N().S(`</div>`)
//line views/components/edit/Timestamp.html:24
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Timestamp.html:24
	qw422016.N().S(`</div>`)
//line views/components/edit/Timestamp.html:26
}

//line views/components/edit/Timestamp.html:26
func WriteTimestampVertical(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:26
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Timestamp.html:26
	StreamTimestampVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:26
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Timestamp.html:26
}

//line views/components/edit/Timestamp.html:26
func TimestampVertical(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/edit/Timestamp.html:26
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Timestamp.html:26
	WriteTimestampVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:26
	qs422016 := string(qb422016.B)
//line views/components/edit/Timestamp.html:26
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Timestamp.html:26
	return qs422016
//line views/components/edit/Timestamp.html:26
}

//line views/components/edit/Timestamp.html:28
func StreamTimestampTable(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:29
	id = cutil.CleanID(key, id)

//line views/components/edit/Timestamp.html:29
	qw422016.N().S(`<tr>`)
//line views/components/edit/Timestamp.html:31
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:31
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Timestamp.html:32
	qw422016.E().S(id)
//line views/components/edit/Timestamp.html:32
	qw422016.N().S(`"`)
//line views/components/edit/Timestamp.html:32
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Timestamp.html:32
	qw422016.N().S(`>`)
//line views/components/edit/Timestamp.html:32
	qw422016.E().S(title)
//line views/components/edit/Timestamp.html:32
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Timestamp.html:33
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:33
	qw422016.N().S(`<td>`)
//line views/components/edit/Timestamp.html:34
	StreamTimestamp(qw422016, key, id, value, help...)
//line views/components/edit/Timestamp.html:34
	qw422016.N().S(`</td>`)
//line views/components/edit/Timestamp.html:35
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Timestamp.html:35
	qw422016.N().S(`</tr>`)
//line views/components/edit/Timestamp.html:37
}

//line views/components/edit/Timestamp.html:37
func WriteTimestampTable(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:37
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Timestamp.html:37
	StreamTimestampTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:37
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Timestamp.html:37
}

//line views/components/edit/Timestamp.html:37
func TimestampTable(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/edit/Timestamp.html:37
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Timestamp.html:37
	WriteTimestampTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:37
	qs422016 := string(qb422016.B)
//line views/components/edit/Timestamp.html:37
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Timestamp.html:37
	return qs422016
//line views/components/edit/Timestamp.html:37
}

//line views/components/edit/Timestamp.html:39
func StreamTimestampDay(qw422016 *qt422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/edit/Timestamp.html:40
	if id == "" {
//line views/components/edit/Timestamp.html:40
		qw422016.N().S(`<input name="`)
//line views/components/edit/Timestamp.html:41
		qw422016.E().S(key)
//line views/components/edit/Timestamp.html:41
		qw422016.N().S(`" type="date" value="`)
//line views/components/edit/Timestamp.html:41
		qw422016.E().S(util.TimeToYMD(value))
//line views/components/edit/Timestamp.html:41
		qw422016.N().S(`"`)
//line views/components/edit/Timestamp.html:41
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Timestamp.html:41
		qw422016.N().S(`/>`)
//line views/components/edit/Timestamp.html:42
	} else {
//line views/components/edit/Timestamp.html:42
		qw422016.N().S(`<input id="`)
//line views/components/edit/Timestamp.html:43
		qw422016.E().S(id)
//line views/components/edit/Timestamp.html:43
		qw422016.N().S(`" name="`)
//line views/components/edit/Timestamp.html:43
		qw422016.E().S(key)
//line views/components/edit/Timestamp.html:43
		qw422016.N().S(`" type="date" value="`)
//line views/components/edit/Timestamp.html:43
		qw422016.E().S(util.TimeToYMD(value))
//line views/components/edit/Timestamp.html:43
		qw422016.N().S(`"`)
//line views/components/edit/Timestamp.html:43
		components.StreamPlaceholderFor(qw422016, placeholder)
//line views/components/edit/Timestamp.html:43
		qw422016.N().S(`/>`)
//line views/components/edit/Timestamp.html:44
	}
//line views/components/edit/Timestamp.html:45
}

//line views/components/edit/Timestamp.html:45
func WriteTimestampDay(qq422016 qtio422016.Writer, key string, id string, value *time.Time, placeholder ...string) {
//line views/components/edit/Timestamp.html:45
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Timestamp.html:45
	StreamTimestampDay(qw422016, key, id, value, placeholder...)
//line views/components/edit/Timestamp.html:45
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Timestamp.html:45
}

//line views/components/edit/Timestamp.html:45
func TimestampDay(key string, id string, value *time.Time, placeholder ...string) string {
//line views/components/edit/Timestamp.html:45
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Timestamp.html:45
	WriteTimestampDay(qb422016, key, id, value, placeholder...)
//line views/components/edit/Timestamp.html:45
	qs422016 := string(qb422016.B)
//line views/components/edit/Timestamp.html:45
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Timestamp.html:45
	return qs422016
//line views/components/edit/Timestamp.html:45
}

//line views/components/edit/Timestamp.html:47
func StreamTimestampDayVertical(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:48
	id = cutil.CleanID(key, id)

//line views/components/edit/Timestamp.html:48
	qw422016.N().S(`<div class="mb expanded">`)
//line views/components/edit/Timestamp.html:50
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:50
	qw422016.N().S(`<label for="`)
//line views/components/edit/Timestamp.html:51
	qw422016.E().S(id)
//line views/components/edit/Timestamp.html:51
	qw422016.N().S(`"><em class="title">`)
//line views/components/edit/Timestamp.html:51
	qw422016.E().S(title)
//line views/components/edit/Timestamp.html:51
	qw422016.N().S(`</em></label>`)
//line views/components/edit/Timestamp.html:52
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:52
	qw422016.N().S(`<div class="mt">`)
//line views/components/edit/Timestamp.html:53
	StreamTimestampDay(qw422016, key, id, value, help...)
//line views/components/edit/Timestamp.html:53
	qw422016.N().S(`</div>`)
//line views/components/edit/Timestamp.html:54
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Timestamp.html:54
	qw422016.N().S(`</div>`)
//line views/components/edit/Timestamp.html:56
}

//line views/components/edit/Timestamp.html:56
func WriteTimestampDayVertical(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:56
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Timestamp.html:56
	StreamTimestampDayVertical(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:56
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Timestamp.html:56
}

//line views/components/edit/Timestamp.html:56
func TimestampDayVertical(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/edit/Timestamp.html:56
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Timestamp.html:56
	WriteTimestampDayVertical(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:56
	qs422016 := string(qb422016.B)
//line views/components/edit/Timestamp.html:56
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Timestamp.html:56
	return qs422016
//line views/components/edit/Timestamp.html:56
}

//line views/components/edit/Timestamp.html:58
func StreamTimestampDayTable(qw422016 *qt422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:59
	id = cutil.CleanID(key, id)

//line views/components/edit/Timestamp.html:59
	qw422016.N().S(`<tr>`)
//line views/components/edit/Timestamp.html:61
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:61
	qw422016.N().S(`<th class="shrink"><label for="`)
//line views/components/edit/Timestamp.html:62
	qw422016.E().S(id)
//line views/components/edit/Timestamp.html:62
	qw422016.N().S(`"`)
//line views/components/edit/Timestamp.html:62
	components.StreamTitleFor(qw422016, help)
//line views/components/edit/Timestamp.html:62
	qw422016.N().S(`>`)
//line views/components/edit/Timestamp.html:62
	qw422016.E().S(title)
//line views/components/edit/Timestamp.html:62
	qw422016.N().S(`</label></th>`)
//line views/components/edit/Timestamp.html:63
	components.StreamIndent(qw422016, true, indent+1)
//line views/components/edit/Timestamp.html:63
	qw422016.N().S(`<td>`)
//line views/components/edit/Timestamp.html:64
	StreamTimestampDay(qw422016, key, id, value, help...)
//line views/components/edit/Timestamp.html:64
	qw422016.N().S(`</td>`)
//line views/components/edit/Timestamp.html:65
	components.StreamIndent(qw422016, true, indent)
//line views/components/edit/Timestamp.html:65
	qw422016.N().S(`</tr>`)
//line views/components/edit/Timestamp.html:67
}

//line views/components/edit/Timestamp.html:67
func WriteTimestampDayTable(qq422016 qtio422016.Writer, key string, id string, title string, value *time.Time, indent int, help ...string) {
//line views/components/edit/Timestamp.html:67
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/edit/Timestamp.html:67
	StreamTimestampDayTable(qw422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:67
	qt422016.ReleaseWriter(qw422016)
//line views/components/edit/Timestamp.html:67
}

//line views/components/edit/Timestamp.html:67
func TimestampDayTable(key string, id string, title string, value *time.Time, indent int, help ...string) string {
//line views/components/edit/Timestamp.html:67
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/edit/Timestamp.html:67
	WriteTimestampDayTable(qb422016, key, id, title, value, indent, help...)
//line views/components/edit/Timestamp.html:67
	qs422016 := string(qb422016.B)
//line views/components/edit/Timestamp.html:67
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/edit/Timestamp.html:67
	return qs422016
//line views/components/edit/Timestamp.html:67
}
