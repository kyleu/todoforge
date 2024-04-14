// Package database - Content managed by Project Forge, see [projectforge.md] for details.
package database

import (
	"context"

	"github.com/kyleu/todoforge/app/util"
)

type SQLiteParams struct {
	File   string `json:"file"`
	Schema string `json:"schema,omitempty"`
	Debug  bool   `json:"debug,omitempty"`
}

func OpenSQLite(ctx context.Context, key string, prefix string, logger util.Logger) (*Service, error) {
	if key == "" {
		key = util.AppKey
	}
	envParams := SQLiteParamsFromEnv(key, prefix)
	return OpenSQLiteDatabase(ctx, key, envParams, logger)
}

func OpenDefaultSQLite(ctx context.Context, logger util.Logger) (*Service, error) {
	return OpenSQLite(ctx, "", "", logger)
}

func SQLiteParamsFromEnv(key string, prefix string) *SQLiteParams {
	f := key + ".sqlite"
	if x := util.GetEnv(prefix + cfgFile); x != "" {
		f = x
	}
	s := "public"
	if x := util.GetEnv(prefix + cfgSchema); x != "" {
		s = x
	}
	debug := false
	if x := util.GetEnv(prefix + cfgDebug); x != "" {
		debug = x != util.BoolFalse
	}
	return &SQLiteParams{File: f, Schema: s, Debug: debug}
}

func SQLiteParamsFromMap(m util.ValueMap) *SQLiteParams {
	return &SQLiteParams{File: m.GetStringOpt("file"), Schema: m.GetStringOpt("schema"), Debug: m.GetBoolOpt("debug")}
}
