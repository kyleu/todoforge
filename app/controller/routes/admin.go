// Package routes - Content managed by Project Forge, see [projectforge.md] for details.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/kyleu/todoforge/app/controller/clib"
)

func adminRoutes(r *mux.Router) {
	makeRoute(r, http.MethodGet, "/admin", clib.Admin)
	makeRoute(r, http.MethodGet, "/admin/audit", clib.AuditList)
	makeRoute(r, http.MethodGet, "/admin/audit/random", clib.AuditCreateFormRandom)
	makeRoute(r, http.MethodGet, "/admin/audit/new", clib.AuditCreateForm)
	makeRoute(r, http.MethodPost, "/admin/audit/new", clib.AuditCreate)
	makeRoute(r, http.MethodGet, "/admin/audit/record/{id}/view", clib.RecordDetail)
	makeRoute(r, http.MethodGet, "/admin/audit/{id}", clib.AuditDetail)
	makeRoute(r, http.MethodGet, "/admin/audit/{id}/edit", clib.AuditEditForm)
	makeRoute(r, http.MethodPost, "/admin/audit/{id}/edit", clib.AuditEdit)
	makeRoute(r, http.MethodGet, "/admin/audit/{id}/delete", clib.AuditDelete)
	makeRoute(r, http.MethodGet, "/admin/database", clib.DatabaseList)
	makeRoute(r, http.MethodGet, "/admin/database/{key}", clib.DatabaseDetail)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/{act}", clib.DatabaseAction)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/tables/{schema}/{table}", clib.DatabaseTableView)
	makeRoute(r, http.MethodGet, "/admin/database/{key}/tables/{schema}/{table}/stats", clib.DatabaseTableStats)
	makeRoute(r, http.MethodPost, "/admin/database/{key}/sql", clib.DatabaseSQLRun)
	makeRoute(r, http.MethodGet, "/admin/{path:.*}", clib.Admin)
	makeRoute(r, http.MethodPost, "/admin/{path:.*}", clib.Admin)
}
