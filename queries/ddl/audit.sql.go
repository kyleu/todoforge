// Code generated by qtc from "audit.sql". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// -- Content managed by Project Forge, see [projectforge.md] for details.
// --

//line queries/ddl/audit.sql:2
package ddl

//line queries/ddl/audit.sql:2
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line queries/ddl/audit.sql:2
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line queries/ddl/audit.sql:2
func StreamAuditDrop(qw422016 *qt422016.Writer) {
//line queries/ddl/audit.sql:2
	qw422016.N().S(`
drop table if exists "audit_record";
drop table if exists "audit";
-- `)
//line queries/ddl/audit.sql:5
}

//line queries/ddl/audit.sql:5
func WriteAuditDrop(qq422016 qtio422016.Writer) {
//line queries/ddl/audit.sql:5
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/audit.sql:5
	StreamAuditDrop(qw422016)
//line queries/ddl/audit.sql:5
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/audit.sql:5
}

//line queries/ddl/audit.sql:5
func AuditDrop() string {
//line queries/ddl/audit.sql:5
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/audit.sql:5
	WriteAuditDrop(qb422016)
//line queries/ddl/audit.sql:5
	qs422016 := string(qb422016.B)
//line queries/ddl/audit.sql:5
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/audit.sql:5
	return qs422016
//line queries/ddl/audit.sql:5
}

// --

//line queries/ddl/audit.sql:7
func StreamAuditCreate(qw422016 *qt422016.Writer) {
//line queries/ddl/audit.sql:7
	qw422016.N().S(`
create table if not exists "audit" (
  "id" uuid not null,
  "app" text not null,
  "act" text not null,
  "client" text not null,
  "server" text not null,
  "user" text not null,
  "metadata" jsonb not null,
  "message" text not null,
  "started" timestamp not null default now(),
  "completed" timestamp not null default now(),
  primary key ("id")
);

create index if not exists "audit__act" on "audit" using btree ("act" asc nulls last);
create index if not exists "audit__app" on "audit" using btree ("app" asc nulls last);
create index if not exists "audit__client" on "audit" using btree ("client" asc nulls last);
create index if not exists "audit__server" on "audit" using btree ("server" asc nulls last);
create index if not exists "audit__user_id" on "audit" using btree ("user" asc nulls last);

create table if not exists "audit_record" (
  "id" uuid not null,
  "audit_id" uuid not null,
  "t" text not null,
  "pk" text not null,
  "changes" jsonb not null,
  "metadata" jsonb not null,
  "occurred" timestamp not null default now(),
  foreign key ("audit_id") references "audit" ("id"),
  primary key ("id")
);

create index if not exists "audit_record__t" on "audit_record" using btree ("t" asc nulls last);
create index if not exists "audit_record__pk" on "audit_record" using btree ("pk" asc nulls last);
create index if not exists "audit_record__changes" on "audit_record" using gin ("changes");
create index if not exists "audit_record__metadata" on "audit_record" using gin ("metadata");

create index if not exists "audit_record__audit_id_idx" on "audit_record" ("audit_id");
-- `)
//line queries/ddl/audit.sql:46
}

//line queries/ddl/audit.sql:46
func WriteAuditCreate(qq422016 qtio422016.Writer) {
//line queries/ddl/audit.sql:46
	qw422016 := qt422016.AcquireWriter(qq422016)
//line queries/ddl/audit.sql:46
	StreamAuditCreate(qw422016)
//line queries/ddl/audit.sql:46
	qt422016.ReleaseWriter(qw422016)
//line queries/ddl/audit.sql:46
}

//line queries/ddl/audit.sql:46
func AuditCreate() string {
//line queries/ddl/audit.sql:46
	qb422016 := qt422016.AcquireByteBuffer()
//line queries/ddl/audit.sql:46
	WriteAuditCreate(qb422016)
//line queries/ddl/audit.sql:46
	qs422016 := string(qb422016.B)
//line queries/ddl/audit.sql:46
	qt422016.ReleaseByteBuffer(qb422016)
//line queries/ddl/audit.sql:46
	return qs422016
//line queries/ddl/audit.sql:46
}
