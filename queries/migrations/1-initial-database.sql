-- {% import "github.com/kyleu/todoforge/queries/ddl" %}
-- {% import "github.com/kyleu/todoforge/queries/seeddata" %}
-- {% func Migration1InitialDatabase(debug bool) %}

-- {%- if debug -%}
-- {%= ddl.DropAll() %}
-- {%- endif -%}

-- {%= ddl.CreateAll() %}
-- {%= seeddata.SeedDataAll() %}
-- {% endfunc %}
