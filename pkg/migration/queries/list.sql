-- List user defined schemas, excluding
--  Extension created schemas
--  Powerbase managed schemas
select pn.nspname
from pg_namespace pn
left join pg_depend pd
  on pd.objid = pn.oid
join pg_roles r 
  on pn.nspowner = r.oid
where pd.deptype is null
  and not pn.nspname like any($1)
  and r.rolname != 'powerbase_admin'
order by pn.nspname
