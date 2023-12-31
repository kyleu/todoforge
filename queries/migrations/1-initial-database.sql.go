// Code generated by qtc from "1-initial-database.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// --

//line queries/migrations/1-initial-database.sql:1
package migrations

//line queries/migrations/1-initial-database.sql:1
import "github.com/kyleu/todoforge/queries/ddl"

// --

//line queries/migrations/1-initial-database.sql:2
import "github.com/kyleu/todoforge/queries/seeddata"

// --

//line queries/migrations/1-initial-database.sql:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/migrations/1-initial-database.sql:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/migrations/1-initial-database.sql:3
func StreamMigration1InitialDatabase(qw422016 *qt422016.Writer, debug bool) {
//line queries/migrations/1-initial-database.sql:3
	qw422016.N().S(`

--`)
//line queries/migrations/1-initial-database.sql:5
	if debug {
//line queries/migrations/1-initial-database.sql:5
		qw422016.N().S(`-- `)
//line queries/migrations/1-initial-database.sql:6
		ddl.StreamDropAll(qw422016)
//line queries/migrations/1-initial-database.sql:6
		qw422016.N().S(`
--`)
//line queries/migrations/1-initial-database.sql:7
	}
//line queries/migrations/1-initial-database.sql:7
	qw422016.N().S(`
-- `)
//line queries/migrations/1-initial-database.sql:9
	ddl.StreamCreateAll(qw422016)
//line queries/migrations/1-initial-database.sql:9
	qw422016.N().S(`
-- `)
//line queries/migrations/1-initial-database.sql:10
	seeddata.StreamSeedDataAll(qw422016)
//line queries/migrations/1-initial-database.sql:10
	qw422016.N().S(`
-- `)
//line queries/migrations/1-initial-database.sql:11
}

//line queries/migrations/1-initial-database.sql:11
func WriteMigration1InitialDatabase(qq422016 qtio422016.Writer, debug bool) {
//line queries/migrations/1-initial-database.sql:11
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/migrations/1-initial-database.sql:11
	StreamMigration1InitialDatabase(qw422016, debug)
//line queries/migrations/1-initial-database.sql:11
	qt422016.ReleaseWriter(qw422016)
//line queries/migrations/1-initial-database.sql:11
}

//line queries/migrations/1-initial-database.sql:11
func Migration1InitialDatabase(debug bool) string {
//line queries/migrations/1-initial-database.sql:11
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/migrations/1-initial-database.sql:11
	WriteMigration1InitialDatabase(qb422016, debug)
//line queries/migrations/1-initial-database.sql:11
	qs422016 := string(qb422016.B)
//line queries/migrations/1-initial-database.sql:11
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/migrations/1-initial-database.sql:11
	return qs422016
//line queries/migrations/1-initial-database.sql:11
}
