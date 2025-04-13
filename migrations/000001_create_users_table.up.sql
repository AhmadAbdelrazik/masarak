CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	email text UNIQUE NOT NULL,
	username text UNIQUE NOT NULL,
	name text NOT NULL,
	password bytea NOT NULL,
	role text NOT NULL,
	version int DEFAULT 1
);
