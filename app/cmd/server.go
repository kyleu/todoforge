// Content managed by Project Forge, see [projectforge.md] for details.
package cmd

import (
	"context"
	"fmt"
	"runtime"

	"github.com/muesli/coral"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/routes"
	"github.com/kyleu/todoforge/app/util"
)

const keyServer = "server"

func serverCmd() *coral.Command {
	short := fmt.Sprintf("Starts the http server on port %d (by default)", util.AppPort)
	f := func(*coral.Command, []string) error { return startServer(_flags) }
	ret := &coral.Command{Use: keyServer, Short: short, RunE: f}
	return ret
}

func startServer(flags *Flags) error {
	if err := initIfNeeded(); err != nil {
		return errors.Wrap(err, "error initializing application")
	}

	st, r, logger, err := loadServer(flags, _logger)
	if err != nil {
		return err
	}

	_, err = listenandserve(util.AppName, flags.Address, flags.Port, r)
	if err != nil {
		return err
	}

	err = st.Close(context.Background(), logger)
	if err != nil {
		return errors.Wrap(err, "unable to close application")
	}

	return nil
}

func loadServer(flags *Flags, logger util.Logger) (*app.State, fasthttp.RequestHandler, util.Logger, error) {
	st, err := buildDefaultAppState(flags, logger)
	if err != nil {
		return nil, nil, logger, err
	}
	logger.Infof("started %s v%s using address [%s:%d] on %s:%s", util.AppName, _buildInfo.Version, flags.Address, flags.Port, runtime.GOOS, runtime.GOARCH)

	controller.SetAppState(st, logger)
	r := routes.AppRoutes(st, logger)

	return st, r, logger, nil
}
