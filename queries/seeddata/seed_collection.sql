-- {% func CollectionSeedData() %}
insert into "collection" (
  "id", "name", "created"
) values (
  '10000000-0000-0000-0000-000000000000', 'Test Collection', current_timestamp
) on conflict do nothing;
-- {% endfunc %}
