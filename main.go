package main // import github.com/kyleu/todoforge

import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/cmd"
)

var (
	version = "0.2.11" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(&app.BuildInfo{Version: version, Commit: commit, Date: date})
}
