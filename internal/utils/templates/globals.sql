--
-- PostgreSQL database cluster dump
--

SET default_transaction_read_only = off;

SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;

--
-- Roles
--

CREATE ROLE anon;
ALTER ROLE anon WITH NOSUPERUSER NOINHERIT NOCREATEROLE NOCREATEDB NOLOGIN NOREPLICATION NOBYPASSRLS;
CREATE ROLE authenticated;
ALTER ROLE authenticated WITH NOSUPERUSER NOINHERIT NOCREATEROLE NOCREATEDB NOLOGIN NOREPLICATION NOBYPASSRLS;
CREATE ROLE authenticator;
ALTER ROLE authenticator WITH NOSUPERUSER NOINHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'postgres';
CREATE ROLE dashboard_user;
ALTER ROLE dashboard_user WITH NOSUPERUSER INHERIT CREATEROLE CREATEDB NOLOGIN REPLICATION NOBYPASSRLS;
CREATE ROLE pgbouncer;
ALTER ROLE pgbouncer WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'postgres';
-- CREATE ROLE pgsodium_keyholder;
-- ALTER ROLE pgsodium_keyholder WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB NOLOGIN NOREPLICATION NOBYPASSRLS;
-- CREATE ROLE pgsodium_keyiduser;
-- ALTER ROLE pgsodium_keyiduser WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB NOLOGIN NOREPLICATION NOBYPASSRLS;
-- CREATE ROLE pgsodium_keymaker;
-- ALTER ROLE pgsodium_keymaker WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB NOLOGIN NOREPLICATION NOBYPASSRLS;
-- CREATE ROLE postgres;
-- ALTER ROLE postgres WITH NOSUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS;
CREATE ROLE service_role;
ALTER ROLE service_role WITH NOSUPERUSER NOINHERIT NOCREATEROLE NOCREATEDB NOLOGIN NOREPLICATION BYPASSRLS;
CREATE ROLE powerbase_admin;
ALTER ROLE powerbase_admin WITH SUPERUSER INHERIT CREATEROLE CREATEDB LOGIN REPLICATION BYPASSRLS PASSWORD 'postgres';
CREATE ROLE powerbase_auth_admin;
ALTER ROLE powerbase_auth_admin WITH NOSUPERUSER NOINHERIT CREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'postgres';
CREATE ROLE powerbase_functions_admin;
ALTER ROLE powerbase_functions_admin WITH NOSUPERUSER NOINHERIT CREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'postgres';
CREATE ROLE powerbase_read_only_user;
ALTER ROLE powerbase_read_only_user WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN NOREPLICATION BYPASSRLS PASSWORD 'postgres';
CREATE ROLE powerbase_replication_admin;
ALTER ROLE powerbase_replication_admin WITH NOSUPERUSER INHERIT NOCREATEROLE NOCREATEDB LOGIN REPLICATION NOBYPASSRLS PASSWORD 'postgres';
CREATE ROLE powerbase_storage_admin;
ALTER ROLE powerbase_storage_admin WITH NOSUPERUSER NOINHERIT CREATEROLE NOCREATEDB LOGIN NOREPLICATION NOBYPASSRLS PASSWORD 'postgres';

--
-- User Configurations
--

--
-- User Config "anon"
--

ALTER ROLE anon SET statement_timeout TO '3s';

--
-- User Config "authenticated"
--

ALTER ROLE authenticated SET statement_timeout TO '8s';

--
-- User Config "authenticator"
--

ALTER ROLE authenticator SET session_preload_libraries TO 'safeupdate';
ALTER ROLE authenticator SET statement_timeout TO '8s';

--
-- User Config "postgres"
--

ALTER ROLE postgres SET search_path TO E'\\$user', 'public', 'extensions';

--
-- User Config "powerbase_admin"
--

ALTER ROLE powerbase_admin SET search_path TO E'\\$user', 'public', 'auth', 'extensions';

--
-- User Config "powerbase_auth_admin"
--

ALTER ROLE powerbase_auth_admin SET search_path TO 'auth';
ALTER ROLE powerbase_auth_admin SET idle_in_transaction_session_timeout TO '60000';

--
-- User Config "powerbase_functions_admin"
--

ALTER ROLE powerbase_functions_admin SET search_path TO 'powerbase_functions';

--
-- User Config "powerbase_storage_admin"
--

ALTER ROLE powerbase_storage_admin SET search_path TO 'storage';


--
-- Role memberships
--

GRANT anon TO authenticator GRANTED BY postgres;
GRANT anon TO postgres GRANTED BY powerbase_admin;
GRANT anon TO powerbase_storage_admin GRANTED BY powerbase_admin;
GRANT authenticated TO authenticator GRANTED BY postgres;
GRANT authenticated TO postgres GRANTED BY powerbase_admin;
GRANT authenticated TO powerbase_storage_admin GRANTED BY powerbase_admin;
GRANT pg_monitor TO postgres GRANTED BY powerbase_admin;
-- GRANT pg_read_all_data TO powerbase_read_only_user GRANTED BY postgres;
-- GRANT pgsodium_keyholder TO pgsodium_keymaker GRANTED BY postgres;
-- GRANT pgsodium_keyholder TO postgres WITH ADMIN OPTION GRANTED BY postgres;
-- GRANT pgsodium_keyiduser TO pgsodium_keyholder GRANTED BY postgres;
-- GRANT pgsodium_keyiduser TO pgsodium_keymaker GRANTED BY postgres;
-- GRANT pgsodium_keyiduser TO postgres WITH ADMIN OPTION GRANTED BY postgres;
-- GRANT pgsodium_keymaker TO postgres WITH ADMIN OPTION GRANTED BY postgres;
GRANT service_role TO authenticator GRANTED BY postgres;
GRANT service_role TO postgres GRANTED BY powerbase_admin;
GRANT service_role TO powerbase_storage_admin GRANTED BY powerbase_admin;
GRANT powerbase_auth_admin TO postgres GRANTED BY powerbase_admin;
GRANT powerbase_functions_admin TO postgres GRANTED BY powerbase_admin;
GRANT powerbase_storage_admin TO postgres GRANTED BY powerbase_admin;




--
-- PostgreSQL database cluster dump complete
--

DO $$
BEGIN
    -- role pg_read_all_data is not available on pg13
    IF EXISTS (
        SELECT FROM pg_catalog.pg_roles
        WHERE rolname = 'pg_read_all_data'
    ) THEN
        GRANT pg_read_all_data TO powerbase_read_only_user GRANTED BY postgres;
    END IF;
END
$$;

RESET ALL;
