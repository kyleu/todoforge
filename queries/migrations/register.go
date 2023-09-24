package migrations

import (
	"github.com/kyleu/todoforge/app/lib/database/migrate"
)

func LoadMigrations(debug bool) {
	migrate.RegisterMigration("create initial database", Migration1InitialDatabase(debug))
}
