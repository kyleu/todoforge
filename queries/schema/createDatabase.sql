-- Content managed by Project Forge, see [projectforge.md] for details.
-- {% func CreateDatabase() %}
create role "todoforge" with login password 'todoforge';

create database "todoforge";
alter database "todoforge" set timezone to 'utc';
grant all privileges on database "todoforge" to "todoforge";
-- {% endfunc %}
