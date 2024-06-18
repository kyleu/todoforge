-- {% func SizeInfo(dbType string) %}
-- {% switch dbType %}
-- {% case "sqlite" %}
select
  'default' as "table_schema",
  "name" as "table_name",
  0 as "row_estimate",
  0 as "total",
  '' as "total_pretty",
  0 as "index",
  '' as "index_pretty",
  0 as "toast",
  '' as "toast_pretty",
  0 as "table",
  '' as "table_pretty"
from "sqlite_master"
where "type" = 'table'
order by "table_name";
-- {% default %}
select 'unhandled database type [{%s dbType %}]';
-- {% endswitch %}
-- {% endfunc %}
