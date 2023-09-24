-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func CollectionDrop() %}
drop table if exists "collection";
-- {% endfunc %}

-- {% func CollectionCreate() %}
create table if not exists "collection" (
  "id" uuid not null,
  "name" text not null,
  "created" timestamp not null default current_timestamp,
  primary key ("id")
);
-- {% endfunc %}
