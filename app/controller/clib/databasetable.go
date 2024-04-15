// Package clib - Content managed by Project Forge, see [projectforge.md] for details.
package clib

import (
	"fmt"
	"net/http"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/lib/database"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views"
	"github.com/kyleu/todoforge/views/vdatabase"
)

func DatabaseTableView(w http.ResponseWriter, r *http.Request) {
	controller.Act("database.table.view", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		svc, err := getDatabaseService(r)
		if err != nil {
			return "", err
		}
		prms := ps.Params.Get("table", []string{"*"}, ps.Logger).Sanitize("table")
		sch, tbl, key := loadTable(r)
		q := database.SQLSelect("*", key, "", prms.OrderByString(), prms.Limit, prms.Offset, svc.Type)
		res, err := svc.QueryRows(ps.Context, q, nil, ps.Logger)
		bc := databaseTableBC(svc.Key)
		ps.SetTitleAndData(tbl+" data", res)
		return controller.Render(r, as, &vdatabase.Results{Svc: svc, Schema: sch, Table: tbl, Results: res, Params: prms, Error: err}, ps, bc...)
	})
}

func DatabaseTableStats(w http.ResponseWriter, r *http.Request) {
	controller.Act("database.table.stats", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		svc, err := getDatabaseService(r)
		if err != nil {
			return "", err
		}
		_, tbl, _ := loadTable(r)
		ret := util.ValueMap{"status": "todo"}
		bc := databaseTableBC(svc.Key, "Stats")
		ps.SetTitleAndData(tbl+" stats", ret)
		return controller.Render(r, as, &views.Debug{}, ps, bc...)
	})
}

func loadTable(r *http.Request) (string, string, string) {
	schema, _ := cutil.PathString(r, "schema", true)
	table, _ := cutil.PathString(r, "table", true)

	tbl := fmt.Sprintf("%q", table)
	if schema != "default" {
		tbl = fmt.Sprintf("%q.%q", schema, table)
	}
	return schema, table, tbl
}

func databaseTableBC(key string, args ...string) []string {
	tbls := fmt.Sprintf("Tables||%s%s/tables", dbRoute, key)
	return databaseBC(key, append([]string{tbls}, args...)...)
}
