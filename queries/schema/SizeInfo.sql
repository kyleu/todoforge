-- {% func SizeInfo(dbType string) %}
-- {% switch dbType %}
-- {% case "sqlite" %}
with
  btree_sizes as (
    select "name", "ncell", "pgsize" as "bytes"
    from "dbstat"
    where "aggregate" = true
  ),
  index_sizes as (
    select m."tbl_name", coalesce(sum(b."bytes"), 0) as "bytes"
    from "sqlite_schema" m
      left join btree_sizes b on b."name" = m."name"
    where m."type" = 'index'
    group by m."tbl_name"
  ),
  table_sizes as (
    select
      m."name",
      case
        when p."wr" = 1 then coalesce(b."ncell", 0)
        else coalesce((select sum(d."ncell") from "dbstat" d where d."name" = m."name" and d."pagetype" = 'leaf'), 0)
      end as "row_estimate",
      coalesce(b."bytes", 0) as "table_bytes",
      coalesce(i."bytes", 0) as "index_bytes"
    from "sqlite_schema" m
      left join "pragma_table_list" p on p."schema" = 'main' and p."name" = m."name"
      left join btree_sizes b on b."name" = m."name"
      left join index_sizes i on i."tbl_name" = m."name"
    where m."type" = 'table' and m."name" not like 'sqlite_%'
  ),
  size_info as (
    select
      "name",
      cast("row_estimate" as text) as "row_estimate",
      "table_bytes" + "index_bytes" as "total",
      "index_bytes" as "index",
      "table_bytes" as "table"
    from table_sizes
  )
select
  'default' as "table_schema",
  "name" as "table_name",
  "row_estimate",
  "total",
  case
    when "total" >= 1099511627776 then printf('%.2f TB', "total" / 1099511627776.0)
    when "total" >= 1073741824 then printf('%.2f GB', "total" / 1073741824.0)
    when "total" >= 1048576 then printf('%.2f MB', "total" / 1048576.0)
    when "total" >= 1024 then printf('%.2f KB', "total" / 1024.0)
    else printf('%d B', "total")
  end as "total_pretty",
  "index",
  case
    when "index" >= 1099511627776 then printf('%.2f TB', "index" / 1099511627776.0)
    when "index" >= 1073741824 then printf('%.2f GB', "index" / 1073741824.0)
    when "index" >= 1048576 then printf('%.2f MB', "index" / 1048576.0)
    when "index" >= 1024 then printf('%.2f KB', "index" / 1024.0)
    else printf('%d B', "index")
  end as "index_pretty",
  0 as "toast",
  '' as "toast_pretty",
  "table",
  case
    when "table" >= 1099511627776 then printf('%.2f TB', "table" / 1099511627776.0)
    when "table" >= 1073741824 then printf('%.2f GB', "table" / 1073741824.0)
    when "table" >= 1048576 then printf('%.2f MB', "table" / 1048576.0)
    when "table" >= 1024 then printf('%.2f KB', "table" / 1024.0)
    else printf('%d B', "table")
  end as "table_pretty"
from size_info
order by "total" desc, "table_name";
-- {% default %}
select 'unhandled database type [{%s dbType %}]';
-- {% endswitch %}
-- {% endfunc %}
