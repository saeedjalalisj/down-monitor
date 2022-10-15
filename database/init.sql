-- CREATE USER docker WITH PASSWORD '123456';
-- CREATE DATABASE docker
--  WITH 
--     OWNER = docker
--     ENCODING = 'UTF8'
--     LC_COLLATE = 'en_US.utf8'
--     LC_CTYPE = 'en_US.utf8'
--     TABLESPACE = pg_default
--     CONNECTION LIMIT = -1;;
-- GRANT ALL PRIVILEGES ON DATABASE docker TO docker;

create extension "uuid-ossp";

CREATE TABLE if not exists service
(
    id uuid default uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    url VARCHAR(250) NOT NULL,
    created_at timestamp with time zone not null default current_timestamp,
    token char(200) NOT NULL
);
