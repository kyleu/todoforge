-- {% func ItemSeedData() %}
insert into "item" (
  "id", "collection_id", "name", "created"
) values (
  '10000000-0000-0000-0000-000000000001', '10000000-0000-0000-0000-000000000000', 'Test Item', current_timestamp
) on conflict do nothing;
-- {% endfunc %}
