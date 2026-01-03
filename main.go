package main // import github.com/kyleu/todoforge

import (
	"context"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/cmd"
)

var (
	version = "0.3.7" // updated by bin/tag.sh and ldflags
	commit  = ""
	date    = "unknown"
)

func main() {
	cmd.Entrypoint(context.Background(), &app.BuildInfo{Version: version, Commit: commit, Date: date})
}
