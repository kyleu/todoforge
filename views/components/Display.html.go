// Code generated by qtc from "Display.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/Display.html:2
package components

//line views/components/Display.html:2
import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/filter"
	"github.com/kyleu/todoforge/app/util"
)

//line views/components/Display.html:14
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/Display.html:14
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/Display.html:14
func StreamDisplayTimestamp(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:14
	qw422016.N().S(`<span class="nowrap">`)
//line views/components/Display.html:15
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:15
	qw422016.N().S(`</span>`)
//line views/components/Display.html:16
}

//line views/components/Display.html:16
func WriteDisplayTimestamp(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:16
	StreamDisplayTimestamp(qw422016, value)
//line views/components/Display.html:16
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:16
}

//line views/components/Display.html:16
func DisplayTimestamp(value *time.Time) string {
//line views/components/Display.html:16
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:16
	WriteDisplayTimestamp(qb422016, value)
//line views/components/Display.html:16
	qs422016 := string(qb422016.B)
//line views/components/Display.html:16
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:16
	return qs422016
//line views/components/Display.html:16
}

//line views/components/Display.html:18
func StreamDisplayTimestampRelative(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:18
	qw422016.N().S(`<span class="nowrap reltime" data-time="`)
//line views/components/Display.html:19
	qw422016.E().S(util.TimeToFull(value))
//line views/components/Display.html:19
	qw422016.N().S(`">`)
//line views/components/Display.html:19
	qw422016.E().S(util.TimeRelative(value))
//line views/components/Display.html:19
	qw422016.N().S(`</span>`)
//line views/components/Display.html:20
}

//line views/components/Display.html:20
func WriteDisplayTimestampRelative(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:20
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:20
	StreamDisplayTimestampRelative(qw422016, value)
//line views/components/Display.html:20
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:20
}

//line views/components/Display.html:20
func DisplayTimestampRelative(value *time.Time) string {
//line views/components/Display.html:20
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:20
	WriteDisplayTimestampRelative(qb422016, value)
//line views/components/Display.html:20
	qs422016 := string(qb422016.B)
//line views/components/Display.html:20
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:20
	return qs422016
//line views/components/Display.html:20
}

//line views/components/Display.html:22
func StreamDisplayTimestampDay(qw422016 *qt422016.Writer, value *time.Time) {
//line views/components/Display.html:23
	qw422016.E().S(util.TimeToYMD(value))
//line views/components/Display.html:24
}

//line views/components/Display.html:24
func WriteDisplayTimestampDay(qq422016 qtio422016.Writer, value *time.Time) {
//line views/components/Display.html:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:24
	StreamDisplayTimestampDay(qw422016, value)
//line views/components/Display.html:24
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:24
}

//line views/components/Display.html:24
func DisplayTimestampDay(value *time.Time) string {
//line views/components/Display.html:24
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:24
	WriteDisplayTimestampDay(qb422016, value)
//line views/components/Display.html:24
	qs422016 := string(qb422016.B)
//line views/components/Display.html:24
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:24
	return qs422016
//line views/components/Display.html:24
}

//line views/components/Display.html:26
func StreamDisplayUUID(qw422016 *qt422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:27
	if value != nil {
//line views/components/Display.html:28
		qw422016.E().S(value.String())
//line views/components/Display.html:29
	}
//line views/components/Display.html:30
}

//line views/components/Display.html:30
func WriteDisplayUUID(qq422016 qtio422016.Writer, value *uuid.UUID) {
//line views/components/Display.html:30
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:30
	StreamDisplayUUID(qw422016, value)
//line views/components/Display.html:30
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:30
}

//line views/components/Display.html:30
func DisplayUUID(value *uuid.UUID) string {
//line views/components/Display.html:30
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:30
	WriteDisplayUUID(qb422016, value)
//line views/components/Display.html:30
	qs422016 := string(qb422016.B)
//line views/components/Display.html:30
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:30
	return qs422016
//line views/components/Display.html:30
}

//line views/components/Display.html:32
func StreamDisplayStringArray(qw422016 *qt422016.Writer, value []string) {
//line views/components/Display.html:33
	if len(value) == 0 {
//line views/components/Display.html:33
		qw422016.N().S(`<em>empty</em>`)
//line views/components/Display.html:35
	}
//line views/components/Display.html:37
	maxCount := 3
	display := value
	var extra int
	if len(value) > maxCount {
		extra = len(value) - maxCount
		display = display[:maxCount]
	}

//line views/components/Display.html:45
	if extra > 0 {
//line views/components/Display.html:45
		qw422016.N().S(`<span title="`)
//line views/components/Display.html:45
		qw422016.E().S(strings.Join(value, `, `))
//line views/components/Display.html:45
		qw422016.N().S(`">`)
//line views/components/Display.html:45
	}
//line views/components/Display.html:46
	for idx, v := range display {
//line views/components/Display.html:47
		if idx > 0 {
//line views/components/Display.html:47
			qw422016.N().S(`,`)
//line views/components/Display.html:47
			qw422016.N().S(` `)
//line views/components/Display.html:47
		}
//line views/components/Display.html:48
		qw422016.E().S(v)
//line views/components/Display.html:49
	}
//line views/components/Display.html:50
	if extra > 0 {
//line views/components/Display.html:50
		qw422016.N().S(`, <em>and`)
//line views/components/Display.html:50
		qw422016.N().S(` `)
//line views/components/Display.html:50
		qw422016.N().D(extra)
//line views/components/Display.html:50
		qw422016.N().S(` `)
//line views/components/Display.html:50
		qw422016.N().S(`more...</em>`)
//line views/components/Display.html:50
	}
//line views/components/Display.html:51
	if extra > 0 {
//line views/components/Display.html:51
		qw422016.N().S(`</span>`)
//line views/components/Display.html:51
	}
//line views/components/Display.html:52
}

//line views/components/Display.html:52
func WriteDisplayStringArray(qq422016 qtio422016.Writer, value []string) {
//line views/components/Display.html:52
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:52
	StreamDisplayStringArray(qw422016, value)
//line views/components/Display.html:52
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:52
}

//line views/components/Display.html:52
func DisplayStringArray(value []string) string {
//line views/components/Display.html:52
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:52
	WriteDisplayStringArray(qb422016, value)
//line views/components/Display.html:52
	qs422016 := string(qb422016.B)
//line views/components/Display.html:52
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:52
	return qs422016
//line views/components/Display.html:52
}

//line views/components/Display.html:54
func StreamDisplayDiffs(qw422016 *qt422016.Writer, value util.Diffs) {
//line views/components/Display.html:55
	if len(value) == 0 {
//line views/components/Display.html:55
		qw422016.N().S(`<em>no changes</em>`)
//line views/components/Display.html:57
	} else {
//line views/components/Display.html:57
		qw422016.N().S(`<div class="overflow full-width"><table><tbody>`)
//line views/components/Display.html:61
		for _, d := range value {
//line views/components/Display.html:61
			qw422016.N().S(`<tr><td><code>`)
//line views/components/Display.html:63
			qw422016.E().S(d.Path)
//line views/components/Display.html:63
			qw422016.N().S(`</code></td><td><code class="error">`)
//line views/components/Display.html:64
			qw422016.E().S(d.Old)
//line views/components/Display.html:64
			qw422016.N().S(`</code></td><td>→</td><td><code class="success">`)
//line views/components/Display.html:66
			qw422016.E().S(d.New)
//line views/components/Display.html:66
			qw422016.N().S(`</code></td></tr>`)
//line views/components/Display.html:68
		}
//line views/components/Display.html:68
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/Display.html:72
	}
//line views/components/Display.html:73
}

//line views/components/Display.html:73
func WriteDisplayDiffs(qq422016 qtio422016.Writer, value util.Diffs) {
//line views/components/Display.html:73
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:73
	StreamDisplayDiffs(qw422016, value)
//line views/components/Display.html:73
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:73
}

//line views/components/Display.html:73
func DisplayDiffs(value util.Diffs) string {
//line views/components/Display.html:73
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:73
	WriteDisplayDiffs(qb422016, value)
//line views/components/Display.html:73
	qs422016 := string(qb422016.B)
//line views/components/Display.html:73
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:73
	return qs422016
//line views/components/Display.html:73
}

//line views/components/Display.html:75
func StreamDisplayMaps(qw422016 *qt422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:76
	if len(maps) == 0 {
//line views/components/Display.html:76
		qw422016.N().S(`<em>no results</em>`)
//line views/components/Display.html:78
	} else {
//line views/components/Display.html:78
		qw422016.N().S(`<div class="overflow full-width"><table><thead><tr>`)
//line views/components/Display.html:83
		for _, k := range maps[0].Keys() {
//line views/components/Display.html:84
			StreamTableHeaderSimple(qw422016, "map", k, k, "", params, nil, ps)
//line views/components/Display.html:85
		}
//line views/components/Display.html:85
		qw422016.N().S(`</tr></thead><tbody>`)
//line views/components/Display.html:89
		for _, m := range maps {
//line views/components/Display.html:89
			qw422016.N().S(`<tr>`)
//line views/components/Display.html:91
			for _, k := range m.Keys() {
//line views/components/Display.html:93
				res := ""
				switch t := m[k].(type) {
				case string:
					res = t
				case []byte:
					res = string(t)
				default:
					res = fmt.Sprint(m[k])
				}

//line views/components/Display.html:103
				if preserveWhitespace {
//line views/components/Display.html:103
					qw422016.N().S(`<td class="prews">`)
//line views/components/Display.html:104
					qw422016.E().S(res)
//line views/components/Display.html:104
					qw422016.N().S(`</td>`)
//line views/components/Display.html:105
				} else {
//line views/components/Display.html:105
					qw422016.N().S(`<td>`)
//line views/components/Display.html:106
					qw422016.E().S(res)
//line views/components/Display.html:106
					qw422016.N().S(`</td>`)
//line views/components/Display.html:107
				}
//line views/components/Display.html:108
			}
//line views/components/Display.html:108
			qw422016.N().S(`</tr>`)
//line views/components/Display.html:110
		}
//line views/components/Display.html:110
		qw422016.N().S(`</tbody></table></div>`)
//line views/components/Display.html:114
	}
//line views/components/Display.html:115
}

//line views/components/Display.html:115
func WriteDisplayMaps(qq422016 qtio422016.Writer, maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) {
//line views/components/Display.html:115
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/Display.html:115
	StreamDisplayMaps(qw422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:115
	qt422016.ReleaseWriter(qw422016)
//line views/components/Display.html:115
}

//line views/components/Display.html:115
func DisplayMaps(maps []util.ValueMap, params *filter.Params, preserveWhitespace bool, ps *cutil.PageState) string {
//line views/components/Display.html:115
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/Display.html:115
	WriteDisplayMaps(qb422016, maps, params, preserveWhitespace, ps)
//line views/components/Display.html:115
	qs422016 := string(qb422016.B)
//line views/components/Display.html:115
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/Display.html:115
	return qs422016
//line views/components/Display.html:115
}
