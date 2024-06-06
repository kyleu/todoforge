// Code generated by qtc from "Bool.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Bool.html:2
package view

//line views/components/view/Bool.html:2
import (
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/views/components"
)

//line views/components/view/Bool.html:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Bool.html:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Bool.html:7
func StreamBool(qw422016 *qt422016.Writer, b bool) {
//line views/components/view/Bool.html:8
	if b {
//line views/components/view/Bool.html:8
		qw422016.N().S(`true`)
//line views/components/view/Bool.html:8
	} else {
//line views/components/view/Bool.html:8
		qw422016.N().S(`false`)
//line views/components/view/Bool.html:8
	}
//line views/components/view/Bool.html:9
}

//line views/components/view/Bool.html:9
func WriteBool(qq422016 qtio422016.Writer, b bool) {
//line views/components/view/Bool.html:9
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Bool.html:9
	StreamBool(qw422016, b)
//line views/components/view/Bool.html:9
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Bool.html:9
}

//line views/components/view/Bool.html:9
func Bool(b bool) string {
//line views/components/view/Bool.html:9
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Bool.html:9
	WriteBool(qb422016, b)
//line views/components/view/Bool.html:9
	qs422016 := string(qb422016.B)
//line views/components/view/Bool.html:9
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Bool.html:9
	return qs422016
//line views/components/view/Bool.html:9
}

//line views/components/view/Bool.html:11
func StreamBoolIcon(qw422016 *qt422016.Writer, b bool, size int, cls string, ps *cutil.PageState) {
//line views/components/view/Bool.html:12
	if cls == "" {
		cls = "inline"
	}

//line views/components/view/Bool.html:13
	if b {
//line views/components/view/Bool.html:14
		components.StreamSVGRef(qw422016, "check", size, size, cls, ps)
//line views/components/view/Bool.html:15
	} else {
//line views/components/view/Bool.html:16
		components.StreamSVGRef(qw422016, "times", size, size, cls, ps)
//line views/components/view/Bool.html:17
	}
//line views/components/view/Bool.html:18
}

//line views/components/view/Bool.html:18
func WriteBoolIcon(qq422016 qtio422016.Writer, b bool, size int, cls string, ps *cutil.PageState) {
//line views/components/view/Bool.html:18
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Bool.html:18
	StreamBoolIcon(qw422016, b, size, cls, ps)
//line views/components/view/Bool.html:18
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Bool.html:18
}

//line views/components/view/Bool.html:18
func BoolIcon(b bool, size int, cls string, ps *cutil.PageState) string {
//line views/components/view/Bool.html:18
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Bool.html:18
	WriteBoolIcon(qb422016, b, size, cls, ps)
//line views/components/view/Bool.html:18
	qs422016 := string(qb422016.B)
//line views/components/view/Bool.html:18
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Bool.html:18
	return qs422016
//line views/components/view/Bool.html:18
}
