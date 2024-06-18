package clib

import (
	"net/http"

	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
)

func Healthcheck(w http.ResponseWriter, _ *http.Request) {
	x := util.ValueMap{"status": "OK"}
	_, _ = cutil.RespondJSON(cutil.NewWriteCounter(w), "", x)
}
