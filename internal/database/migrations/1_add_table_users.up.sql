CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
	id uuid PRIMARY KEY DEFAULT public.uuid_generate_v4(),
	firstname  varchar(256) NOT NULL,
	lastname varchar(256) NOT NULL,
	email varchar(256) NOT NULL,
	age smallint NOT NULL,
	created timestamp NOT NULL
);