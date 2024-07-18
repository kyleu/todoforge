package cmd

import (
	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/util"
)

func Run(bi *app.BuildInfo) (util.Logger, error) {
	_buildInfo = bi

	if err := rootCmd().Execute(); err != nil {
		return util.RootLogger, err
	}
	return util.RootLogger, nil
}
