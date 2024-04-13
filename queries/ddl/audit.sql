-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func AuditDrop() %}
drop table if exists "audit_record";
drop table if exists "audit";
-- {% endfunc %}

-- {% func AuditCreate() %}
create table if not exists "audit" (
  "id" uuid not null,
  "app" text not null,
  "act" text not null,
  "client" text not null,
  "server" text not null,
  "user" text not null,
  "metadata" jsonb not null,
  "message" text not null,
  "started" timestamp not null default current_timestamp,
  "completed" timestamp not null default current_timestamp,
  primary key ("id")
);

create index if not exists "audit__act" on "audit" ("act");
create index if not exists "audit__app" on "audit" ("app");
create index if not exists "audit__client" on "audit" ("client");
create index if not exists "audit__server" on "audit" ("server");
create index if not exists "audit__user_id" on "audit" ("user");

create table if not exists "audit_record" (
  "id" uuid not null,
  "audit_id" uuid not null,
  "t" text not null,
  "pk" text not null,
  "changes" jsonb not null,
  "metadata" jsonb not null,
  "occurred" timestamp not null default current_timestamp,
  foreign key ("audit_id") references "audit" ("id"),
  primary key ("id")
);

create index if not exists "audit_record__t" on "audit_record"("t");
create index if not exists "audit_record__pk" on "audit_record"("pk");
create index if not exists "audit_record__changes" on "audit_record"("changes");
create index if not exists "audit_record__metadata" on "audit_record"("metadata");

create index if not exists "audit_record__audit_id_idx" on "audit_record" ("audit_id");
-- {% endfunc %}
