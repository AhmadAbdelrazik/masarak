CREATE TABLE IF NOT EXISTS users(
	email text PRIMARY KEY,
	name text NOT NULL,
	password bytea NOT NULL,
	role text NOT NULL,
	version int DEFAULT 1
);
