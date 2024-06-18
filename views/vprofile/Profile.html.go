// Code generated by qtc from "Profile.html". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/vprofile/Profile.html:1
package vprofile

//line views/vprofile/Profile.html:1
import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/theme"
	"github.com/kyleu/todoforge/app/lib/user"
	"github.com/kyleu/todoforge/views/components"
	"github.com/kyleu/todoforge/views/components/edit"
	"github.com/kyleu/todoforge/views/layout"
	"github.com/kyleu/todoforge/views/vtheme"
)

//line views/vprofile/Profile.html:12
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/vprofile/Profile.html:12
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/vprofile/Profile.html:12
type Profile struct {
	layout.Basic
	Profile  *user.Profile
	Theme    *theme.Theme
	Referrer string
}

//line views/vprofile/Profile.html:19
func (p *Profile) StreamBody(qw422016 *qt422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vprofile/Profile.html:19
	qw422016.N().S(`
  <form action="" method="post">
    <input name="referrer" type="hidden" value="`)
//line views/vprofile/Profile.html:21
	qw422016.E().S(p.Referrer)
//line views/vprofile/Profile.html:21
	qw422016.N().S(`" />
    <div class="card">
      <div class="right">
        <a href="#modal-profile"><button type="button">Profile</button></a>
        <a href="#modal-theme"><button type="button">Theme</button></a>
      </div>
      <h3>Profile</h3>
      <div class="overflow full-width">
        <table class="mt expanded">
          <tbody>
            `)
//line views/vprofile/Profile.html:31
	edit.StreamStringTable(qw422016, "name", "", "Name", ps.Username(), 5)
//line views/vprofile/Profile.html:31
	qw422016.N().S(`
            <tr>
              <th class="shrink"><label>Mode</label></th>
              <td>
                <label>
`)
//line views/vprofile/Profile.html:36
	if p.Profile.Mode == "" {
//line views/vprofile/Profile.html:36
		qw422016.N().S(`                  <input type="radio" class="mode-input" name="mode" value="" checked="checked" />
`)
//line views/vprofile/Profile.html:38
	} else {
//line views/vprofile/Profile.html:38
		qw422016.N().S(`                  <input type="radio" class="mode-input" name="mode" value="" />
`)
//line views/vprofile/Profile.html:40
	}
//line views/vprofile/Profile.html:40
	qw422016.N().S(`                  System Default
                </label>
                <label>
`)
//line views/vprofile/Profile.html:44
	if p.Profile.Mode == "light" {
//line views/vprofile/Profile.html:44
		qw422016.N().S(`                  <input type="radio" class="mode-input" name="mode" value="light" checked="checked" />
`)
//line views/vprofile/Profile.html:46
	} else {
//line views/vprofile/Profile.html:46
		qw422016.N().S(`                  <input type="radio" class="mode-input" name="mode" value="light" />
`)
//line views/vprofile/Profile.html:48
	}
//line views/vprofile/Profile.html:48
	qw422016.N().S(`                  Light
                </label>
                <label>
`)
//line views/vprofile/Profile.html:52
	if p.Profile.Mode == "dark" {
//line views/vprofile/Profile.html:52
		qw422016.N().S(`                  <input type="radio" class="mode-input" name="mode" value="dark" checked="checked" />
`)
//line views/vprofile/Profile.html:54
	} else {
//line views/vprofile/Profile.html:54
		qw422016.N().S(`                  <input type="radio" class="mode-input" name="mode" value="dark" />
`)
//line views/vprofile/Profile.html:56
	}
//line views/vprofile/Profile.html:56
	qw422016.N().S(`                  Dark
                </label>
              </td>
            </tr>
            `)
//line views/vprofile/Profile.html:61
	vtheme.StreamChoicePanel(qw422016, as.Themes.All(ps.Logger), "app", 5, ps)
//line views/vprofile/Profile.html:61
	qw422016.N().S(`
          </tbody>
        </table>
      </div>
      <div class="mt">
        <button type="submit">Save All Changes</button>
        <button type="reset">Reset</button>
      </div>
    </div>
  </form>

  `)
//line views/vprofile/Profile.html:72
	components.StreamJSONModal(qw422016, "profile", "Profile JSON", p.Profile, 1)
//line views/vprofile/Profile.html:72
	qw422016.N().S(`
  `)
//line views/vprofile/Profile.html:73
	components.StreamJSONModal(qw422016, "theme", "Theme JSON", p.Theme, 1)
//line views/vprofile/Profile.html:73
	qw422016.N().S(`
`)
//line views/vprofile/Profile.html:74
}

//line views/vprofile/Profile.html:74
func (p *Profile) WriteBody(qq422016 qtio422016.Writer, as *app.State, ps *cutil.PageState) {
//line views/vprofile/Profile.html:74
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/vprofile/Profile.html:74
	p.StreamBody(qw422016, as, ps)
//line views/vprofile/Profile.html:74
	qt422016.ReleaseWriter(qw422016)
//line views/vprofile/Profile.html:74
}

//line views/vprofile/Profile.html:74
func (p *Profile) Body(as *app.State, ps *cutil.PageState) string {
//line views/vprofile/Profile.html:74
	qb422016 := qt422016.AcquireByteBuffer()
//line views/vprofile/Profile.html:74
	p.WriteBody(qb422016, as, ps)
//line views/vprofile/Profile.html:74
	qs422016 := string(qb422016.B)
//line views/vprofile/Profile.html:74
	qt422016.ReleaseByteBuffer(qb422016)
//line views/vprofile/Profile.html:74
	return qs422016
//line views/vprofile/Profile.html:74
}
