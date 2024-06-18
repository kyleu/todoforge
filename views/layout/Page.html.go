// Code generated by qtc from "Page.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// This is the base page interface. All pages implement this interface.
//

//line views/layout/Page.html:3
package layout

//line views/layout/Page.html:3
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
)

//line views/layout/Page.html:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Page.html:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Page.html:9
type Page interface {
//line views/layout/Page.html:9
	Head(as *app.State, ps *cutil.PageState) string
//line views/layout/Page.html:9
	StreamHead(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:9
	WriteHead(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:9
	Nav(as *app.State, ps *cutil.PageState) string
//line views/layout/Page.html:9
	StreamNav(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:9
	WriteNav(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:9
	Menu(ps *cutil.PageState) string
//line views/layout/Page.html:9
	StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState)
//line views/layout/Page.html:9
	WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState)
//line views/layout/Page.html:9
	Body(as *app.State, ps *cutil.PageState) string
//line views/layout/Page.html:9
	StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:9
	WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState)
//line views/layout/Page.html:9
}
