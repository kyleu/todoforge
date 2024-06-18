-- {% func ItemDrop() %}
drop table if exists "item";
-- {% endfunc %}

-- {% func ItemCreate() %}
create table if not exists "item" (
  "id" uuid not null,
  "collection_id" uuid not null,
  "name" text not null,
  "created" timestamp not null default current_timestamp,
  foreign key ("collection_id") references "collection" ("id"),
  primary key ("id")
);

create index if not exists "item__collection_id_idx" on "item" ("collection_id");
-- {% endfunc %}
