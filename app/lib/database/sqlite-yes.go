//go:build darwin || (!android && linux && 386) || (!android && linux && amd64) || (!android && linux && arm) || (!android && linux && arm64) || (!android && linux && riscv64) || (windows && amd64)

// Content managed by Project Forge, see [projectforge.md] for details.
package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	_ "modernc.org/sqlite" // load sqlite driver.

	"github.com/kyleu/todoforge/app/lib/telemetry"
	"github.com/kyleu/todoforge/app/util"
)

const SQLiteEnabled = true

var typeSQLite = &DBType{Key: "sqlite", Title: "SQLite", Quote: `"`, Placeholder: "$", SupportsReturning: true}

func OpenSQLiteDatabase(ctx context.Context, key string, params *SQLiteParams, logger util.Logger) (*Service, error) {
	_, span, logger := telemetry.StartSpan(ctx, "database:open", logger)
	defer span.Complete()
	if params.File == "" {
		return nil, errors.New("need filename for SQLite database")
	}
	db, err := sqlx.Open("sqlite", params.File)
	if err != nil {
		return nil, errors.Wrap(err, "error opening database")
	}
	return NewService(typeSQLite, key, key, params.Schema, "sqlite", params.Debug, db, logger)
}
