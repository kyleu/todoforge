// Content managed by Project Forge, see [projectforge.md] for details.
package main // import github.com/kyleu/todoforge

import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/cmd"
)

var (
	version = "0.0.1" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
