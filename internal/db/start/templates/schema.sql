\set pgpass `echo "$PGPASSWORD"`
\set jwt_secret `echo "$JWT_SECRET"`
\set jwt_exp `echo "$JWT_EXP"`

ALTER DATABASE postgres SET "app.settings.jwt_secret" TO :'jwt_secret';
ALTER DATABASE postgres SET "app.settings.jwt_exp" TO :'jwt_exp';

ALTER USER postgres WITH PASSWORD :'pgpass';
ALTER USER authenticator WITH PASSWORD :'pgpass';
ALTER USER pgbouncer WITH PASSWORD :'pgpass';
ALTER USER powerbase_auth_admin WITH PASSWORD :'pgpass';
ALTER USER powerbase_storage_admin WITH PASSWORD :'pgpass';
ALTER USER powerbase_replication_admin WITH PASSWORD :'pgpass';
ALTER USER powerbase_read_only_user WITH PASSWORD :'pgpass';

create schema if not exists _realtime;
alter schema _realtime owner to postgres;
