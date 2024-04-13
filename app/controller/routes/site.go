// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/clib"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
)

func SiteRoutes(logger util.Logger) (http.Handler, error) {
	r := mux.NewRouter()

	makeRoute(r, http.MethodGet, cutil.DefaultProfilePath, clib.ProfileSite)
	makeRoute(r, http.MethodPost, cutil.DefaultProfilePath, clib.ProfileSave)

	makeRoute(r, http.MethodGet, "/favicon.ico", clib.Favicon)
	makeRoute(r, http.MethodGet, "/assets/{path:.*}", clib.Static)

	makeRoute(r, http.MethodGet, "/", controller.Site)
	makeRoute(r, http.MethodGet, "/{path:.*}", controller.Site)

	makeRoute(r, http.MethodOptions, "/", controller.Options)

	return cutil.WireRouter(r, controller.NotFoundAction, logger)
}
