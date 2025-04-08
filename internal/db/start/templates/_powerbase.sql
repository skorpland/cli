CREATE DATABASE _powerbase WITH OWNER postgres;

-- Switch to the newly created _powerbase database
\c _powerbase
-- Create schemas in _powerbase database for
-- internals tools and reports to not overload user database
-- with non-user activity
CREATE SCHEMA IF NOT EXISTS _analytics;
ALTER SCHEMA _analytics OWNER TO postgres;

CREATE SCHEMA IF NOT EXISTS _powerpooler;
ALTER SCHEMA _powerpooler OWNER TO postgres;
\c postgres
