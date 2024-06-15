// Code generated by qtc from "Menu.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// <!-- Content managed by Project Forge, see [projectforge.md] for details. -->

//line views/layout/Menu.html:2
package layout

//line views/layout/Menu.html:2
import (
	"fmt"
	"strings"

	"github.com/kyleu/todoforge/app/controller/cmenu"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/menu"
	"github.com/kyleu/todoforge/views/components"
)

//line views/layout/Menu.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/layout/Menu.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/layout/Menu.html:12
func StreamMenu(qw422016 *qt422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:13
	if len(ps.Menu.Visible()) > 0 {
//line views/layout/Menu.html:13
		qw422016.N().S(`<div class="menu-container">`)
//line views/layout/Menu.html:15
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:15
		qw422016.N().S(`<div class="menu">`)
//line views/layout/Menu.html:17
		components.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:17
		qw422016.N().S(`<menu class="level-0">`)
//line views/layout/Menu.html:19
		for _, i := range ps.Menu {
//line views/layout/Menu.html:20
			StreamMenuItem(qw422016, i, []string{}, ps.Breadcrumbs, 3, ps)
//line views/layout/Menu.html:21
		}
//line views/layout/Menu.html:22
		components.StreamIndent(qw422016, true, 3)
//line views/layout/Menu.html:22
		qw422016.N().S(`</menu>`)
//line views/layout/Menu.html:24
		components.StreamIndent(qw422016, true, 2)
//line views/layout/Menu.html:24
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:26
		components.StreamIndent(qw422016, true, 1)
//line views/layout/Menu.html:26
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:28
	}
//line views/layout/Menu.html:29
}

//line views/layout/Menu.html:29
func WriteMenu(qq422016 qtio422016.Writer, ps *cutil.PageState) {
//line views/layout/Menu.html:29
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:29
	StreamMenu(qw422016, ps)
//line views/layout/Menu.html:29
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:29
}

//line views/layout/Menu.html:29
func Menu(ps *cutil.PageState) string {
//line views/layout/Menu.html:29
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:29
	WriteMenu(qb422016, ps)
//line views/layout/Menu.html:29
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:29
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:29
	return qs422016
//line views/layout/Menu.html:29
}

//line views/layout/Menu.html:31
func StreamMenuItem(qw422016 *qt422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:33
	path = append(path, i.Key)
	active, final := breadcrumbs.Active(i, path)

//line views/layout/Menu.html:36
	if i.Key == "" {
//line views/layout/Menu.html:37
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:37
		qw422016.N().S(`<li class="separator"></li>`)
//line views/layout/Menu.html:39
	} else if len(i.Children.Visible()) > 0 {
//line views/layout/Menu.html:40
		itemID := strings.Join(path, "--")

//line views/layout/Menu.html:41
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:42
		if active {
//line views/layout/Menu.html:42
			qw422016.N().S(`<li class="active" data-menu-key="`)
//line views/layout/Menu.html:42
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:42
			qw422016.N().S(`">`)
//line views/layout/Menu.html:42
		} else {
//line views/layout/Menu.html:42
			qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:42
			qw422016.E().S(i.Key)
//line views/layout/Menu.html:42
			qw422016.N().S(`">`)
//line views/layout/Menu.html:42
		}
//line views/layout/Menu.html:43
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:43
		qw422016.N().S(`<input id="`)
//line views/layout/Menu.html:44
		qw422016.E().S(itemID)
//line views/layout/Menu.html:44
		qw422016.N().S(`-input" type="checkbox"`)
//line views/layout/Menu.html:44
		if active {
//line views/layout/Menu.html:44
			qw422016.N().S(` `)
//line views/layout/Menu.html:44
			qw422016.N().S(`checked="checked"`)
//line views/layout/Menu.html:44
		}
//line views/layout/Menu.html:44
		qw422016.N().S(` `)
//line views/layout/Menu.html:44
		qw422016.N().S(`hidden="hidden" />`)
//line views/layout/Menu.html:45
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:46
		if final {
//line views/layout/Menu.html:46
			qw422016.N().S(`<label class="final" for="`)
//line views/layout/Menu.html:46
			qw422016.E().S(itemID)
//line views/layout/Menu.html:46
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:46
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:46
			qw422016.N().S(`">`)
//line views/layout/Menu.html:46
		} else {
//line views/layout/Menu.html:46
			qw422016.N().S(`<label for="`)
//line views/layout/Menu.html:46
			qw422016.E().S(itemID)
//line views/layout/Menu.html:46
			qw422016.N().S(`-input" title="`)
//line views/layout/Menu.html:46
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:46
			qw422016.N().S(`">`)
//line views/layout/Menu.html:46
		}
//line views/layout/Menu.html:47
		if i.Route != "" && i.Badge == "" {
//line views/layout/Menu.html:48
			components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:48
			qw422016.N().S(`<a class="label-link" href="`)
//line views/layout/Menu.html:49
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:49
			qw422016.N().S(`">`)
//line views/layout/Menu.html:49
			components.StreamSVGSimple(qw422016, `link`, 15, ps)
//line views/layout/Menu.html:49
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:50
		}
//line views/layout/Menu.html:51
		components.StreamExpandCollapse(qw422016, indent+3, ps)
//line views/layout/Menu.html:52
		streammenuBadge(qw422016, i, indent+3, ps)
//line views/layout/Menu.html:53
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:54
		if i.Icon != "" {
//line views/layout/Menu.html:55
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:55
			qw422016.N().S(` `)
//line views/layout/Menu.html:56
		}
//line views/layout/Menu.html:57
		if i.Route != "" {
//line views/layout/Menu.html:58
			if i.Warning != "" {
//line views/layout/Menu.html:58
				qw422016.N().S(`<a class="link-confirm" data-message="`)
//line views/layout/Menu.html:59
				qw422016.E().S(i.Warning)
//line views/layout/Menu.html:59
				qw422016.N().S(`" href="`)
//line views/layout/Menu.html:59
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:59
				qw422016.N().S(`">`)
//line views/layout/Menu.html:59
				qw422016.E().S(i.Title)
//line views/layout/Menu.html:59
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:60
			} else {
//line views/layout/Menu.html:60
				qw422016.N().S(`<a href="`)
//line views/layout/Menu.html:61
				qw422016.E().S(i.Route)
//line views/layout/Menu.html:61
				qw422016.N().S(`">`)
//line views/layout/Menu.html:61
				qw422016.E().S(i.Title)
//line views/layout/Menu.html:61
				qw422016.N().S(`</a>`)
//line views/layout/Menu.html:62
			}
//line views/layout/Menu.html:63
		} else {
//line views/layout/Menu.html:64
			qw422016.E().S(i.Title)
//line views/layout/Menu.html:65
		}
//line views/layout/Menu.html:66
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:67
		if final {
//line views/layout/Menu.html:67
			qw422016.N().S(`</label>`)
//line views/layout/Menu.html:67
		} else {
//line views/layout/Menu.html:67
			qw422016.N().S(`</label>`)
//line views/layout/Menu.html:67
		}
//line views/layout/Menu.html:68
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:68
		qw422016.N().S(`<div class="menu-content level-`)
//line views/layout/Menu.html:69
		qw422016.N().D(len(path))
//line views/layout/Menu.html:69
		qw422016.N().S(`">`)
//line views/layout/Menu.html:70
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:70
		qw422016.N().S(`<menu>`)
//line views/layout/Menu.html:72
		for _, i := range i.Children.Visible() {
//line views/layout/Menu.html:73
			StreamMenuItem(qw422016, i, path, breadcrumbs, indent+3, ps)
//line views/layout/Menu.html:74
		}
//line views/layout/Menu.html:75
		components.StreamIndent(qw422016, true, indent+3)
//line views/layout/Menu.html:75
		qw422016.N().S(`</menu>`)
//line views/layout/Menu.html:77
		components.StreamIndent(qw422016, true, indent+2)
//line views/layout/Menu.html:77
		qw422016.N().S(`</div>`)
//line views/layout/Menu.html:79
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:79
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:81
	} else {
//line views/layout/Menu.html:83
		finalClass := "item"
		if active {
			finalClass += " active"
		}
		if final {
			finalClass += " final"
		}
		if i.Warning != "" {
			finalClass += " link-confirm"
		}

//line views/layout/Menu.html:94
		components.StreamIndent(qw422016, true, indent+1)
//line views/layout/Menu.html:94
		qw422016.N().S(`<li data-menu-key="`)
//line views/layout/Menu.html:95
		qw422016.E().S(i.Key)
//line views/layout/Menu.html:95
		qw422016.N().S(`">`)
//line views/layout/Menu.html:96
		if i.Warning != "" {
//line views/layout/Menu.html:96
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:97
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:97
			qw422016.N().S(`" data-message="`)
//line views/layout/Menu.html:97
			qw422016.E().S(i.Warning)
//line views/layout/Menu.html:97
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:97
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:97
			qw422016.N().S(`" title="`)
//line views/layout/Menu.html:97
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:97
			qw422016.N().S(`">`)
//line views/layout/Menu.html:98
		} else {
//line views/layout/Menu.html:98
			qw422016.N().S(`<a class="`)
//line views/layout/Menu.html:99
			qw422016.E().S(finalClass)
//line views/layout/Menu.html:99
			qw422016.N().S(`" href="`)
//line views/layout/Menu.html:99
			qw422016.E().S(i.Route)
//line views/layout/Menu.html:99
			qw422016.N().S(`" title="`)
//line views/layout/Menu.html:99
			qw422016.E().S(i.Desc())
//line views/layout/Menu.html:99
			qw422016.N().S(`">`)
//line views/layout/Menu.html:100
		}
//line views/layout/Menu.html:101
		streammenuBadge(qw422016, i, indent+3, ps)
//line views/layout/Menu.html:102
		if i.Icon != "" {
//line views/layout/Menu.html:103
			components.StreamSVGRef(qw422016, i.Icon, 16, 16, "icon", ps)
//line views/layout/Menu.html:103
			qw422016.N().S(` `)
//line views/layout/Menu.html:104
		}
//line views/layout/Menu.html:104
		qw422016.N().S(`<span>`)
//line views/layout/Menu.html:105
		qw422016.E().S(i.Title)
//line views/layout/Menu.html:105
		qw422016.N().S(`</span>`)
//line views/layout/Menu.html:106
		if i.Warning != "" {
//line views/layout/Menu.html:106
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:108
		} else {
//line views/layout/Menu.html:108
			qw422016.N().S(`</a>`)
//line views/layout/Menu.html:110
		}
//line views/layout/Menu.html:110
		qw422016.N().S(`</li>`)
//line views/layout/Menu.html:112
	}
//line views/layout/Menu.html:113
}

//line views/layout/Menu.html:113
func WriteMenuItem(qq422016 qtio422016.Writer, i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:113
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:113
	StreamMenuItem(qw422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:113
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:113
}

//line views/layout/Menu.html:113
func MenuItem(i *menu.Item, path []string, breadcrumbs cmenu.Breadcrumbs, indent int, ps *cutil.PageState) string {
//line views/layout/Menu.html:113
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:113
	WriteMenuItem(qb422016, i, path, breadcrumbs, indent, ps)
//line views/layout/Menu.html:113
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:113
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:113
	return qs422016
//line views/layout/Menu.html:113
}

//line views/layout/Menu.html:115
func streammenuBadge(qw422016 *qt422016.Writer, i *menu.Item, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:116
	if i.Badge != "" {
//line views/layout/Menu.html:118
		badgeTitle := ""
		if idx := strings.Index(i.Badge, "**"); idx > -1 {
			badgeTitle = fmt.Sprintf(" title=%q", i.Badge[idx+2:])
			i.Badge = i.Badge[:idx]
		}

//line views/layout/Menu.html:124
		if strings.HasPrefix(i.Badge, ":") {
//line views/layout/Menu.html:125
			components.StreamIndent(qw422016, true, indent)
//line views/layout/Menu.html:125
			qw422016.N().S(`<div class="badge"`)
//line views/layout/Menu.html:126
			qw422016.N().S(badgeTitle)
//line views/layout/Menu.html:126
			qw422016.N().S(`>`)
//line views/layout/Menu.html:126
			components.StreamSVGSimple(qw422016, i.Badge[1:], 18, ps)
//line views/layout/Menu.html:126
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:127
		} else {
//line views/layout/Menu.html:128
			components.StreamIndent(qw422016, true, indent)
//line views/layout/Menu.html:128
			qw422016.N().S(`<div class="badge"`)
//line views/layout/Menu.html:129
			qw422016.N().S(badgeTitle)
//line views/layout/Menu.html:129
			qw422016.N().S(`>`)
//line views/layout/Menu.html:129
			qw422016.E().S(i.Badge)
//line views/layout/Menu.html:129
			qw422016.N().S(`</div>`)
//line views/layout/Menu.html:130
		}
//line views/layout/Menu.html:131
	}
//line views/layout/Menu.html:132
}

//line views/layout/Menu.html:132
func writemenuBadge(qq422016 qtio422016.Writer, i *menu.Item, indent int, ps *cutil.PageState) {
//line views/layout/Menu.html:132
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/layout/Menu.html:132
	streammenuBadge(qw422016, i, indent, ps)
//line views/layout/Menu.html:132
	qt422016.ReleaseWriter(qw422016)
//line views/layout/Menu.html:132
}

//line views/layout/Menu.html:132
func menuBadge(i *menu.Item, indent int, ps *cutil.PageState) string {
//line views/layout/Menu.html:132
	qb422016 := qt422016.AcquireByteBuffer()
//line views/layout/Menu.html:132
	writemenuBadge(qb422016, i, indent, ps)
//line views/layout/Menu.html:132
	qs422016 := string(qb422016.B)
//line views/layout/Menu.html:132
	qt422016.ReleaseByteBuffer(qb422016)
//line views/layout/Menu.html:132
	return qs422016
//line views/layout/Menu.html:132
}
