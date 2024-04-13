package controller

import (
	"fmt"
	"net/http"

	"github.com/samber/lo"

	"github.com/kyleu/todoforge/app"
	"github.com/kyleu/todoforge/app/controller/cutil"
	"github.com/kyleu/todoforge/app/util"
	"github.com/kyleu/todoforge/views/vgraphql"
)

func GraphQLIndex(w http.ResponseWriter, r *http.Request) {
	Act("graphql.index", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		keys := as.GraphQL.Keys()
		if len(keys) == 1 {
			return "/graphql/" + keys[0], nil
		}
		counts := lo.Map(keys, func(key string, _ int) int {
			return as.GraphQL.ExecCount(key)
		})
		ps.SetTitleAndData("GraphQL List", keys)
		return Render(w, r, as, &vgraphql.List{Keys: keys, Counts: counts}, ps, "graphql")
	})
}

func GraphQLDetail(w http.ResponseWriter, r *http.Request) {
	Act("graphql.detail", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}
		titles := as.GraphQL.Titles()
		if len(titles) > 1 {
			ps.Title = fmt.Sprintf("[%s] GraphQL Workspace", titles[key])
		} else {
			ps.Title = "GraphQL Workspace"
		}
		ps.Data = key
		bc := []string{"graphql"}
		if len(titles) > 1 {
			bc = append(bc, key)
		}
		return Render(w, r, as, &vgraphql.Detail{Key: key}, ps, bc...)
	})
}

func GraphQLRun(w http.ResponseWriter, r *http.Request) {
	Act("graphql.run", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		key, err := cutil.PathString(r, "key", false)
		if err != nil {
			return "", err
		}
		frm, err := cutil.ParseForm(r, ps.RequestBody)
		if err != nil {
			return "", err
		}
		q := frm.GetStringOpt("query")
		op := frm.GetStringOpt("operationName")
		v := frm.GetStringOpt("variables")
		vars := util.ValueMap{}
		if v != "" {
			_ = util.FromJSON([]byte(v), &vars)
		}
		rsp, err := as.GraphQL.Exec(ps.Context, key, q, op, vars, ps.Logger)
		if err != nil {
			return "", err
		}
		return cutil.RespondMIME("", "application/json", "json", util.ToJSONBytes(rsp, true), w)
	})
}
