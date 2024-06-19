package app

import (
	"context"

	"github.com/kyleu/todoforge/app/lib/audit"
	"github.com/kyleu/todoforge/app/util"
)

type CoreServices struct {
	Audit *audit.Service
}

func initCoreServices(ctx context.Context, st *State, auditSvc *audit.Service, logger util.Logger) CoreServices {
	return CoreServices{
		Audit: auditSvc,
	}
}
