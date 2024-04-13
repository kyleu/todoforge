// Code generated by qtc from "Package.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/components/view/Package.html:2
package view

//line views/components/view/Package.html:2
import "github.com/kyleu/todoforge/app/util"

//line views/components/view/Package.html:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/components/view/Package.html:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/components/view/Package.html:4
func StreamPackage(qw422016 *qt422016.Writer, v util.Pkg) {
//line views/components/view/Package.html:5
	for idx, x := range v {
//line views/components/view/Package.html:6
		qw422016.E().S(x)
//line views/components/view/Package.html:6
		if len(v) < idx {
//line views/components/view/Package.html:6
			qw422016.N().S(`/`)
//line views/components/view/Package.html:6
		}
//line views/components/view/Package.html:7
	}
//line views/components/view/Package.html:8
}

//line views/components/view/Package.html:8
func WritePackage(qq422016 qtio422016.Writer, v util.Pkg) {
//line views/components/view/Package.html:8
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/components/view/Package.html:8
	StreamPackage(qw422016, v)
//line views/components/view/Package.html:8
	qt422016.ReleaseWriter(qw422016)
//line views/components/view/Package.html:8
}

//line views/components/view/Package.html:8
func Package(v util.Pkg) string {
//line views/components/view/Package.html:8
	qb422016 := qt422016.AcquireByteBuffer()
//line views/components/view/Package.html:8
	WritePackage(qb422016, v)
//line views/components/view/Package.html:8
	qs422016 := string(qb422016.B)
//line views/components/view/Package.html:8
	qt422016.ReleaseByteBuffer(qb422016)
//line views/components/view/Package.html:8
	return qs422016
//line views/components/view/Package.html:8
}
