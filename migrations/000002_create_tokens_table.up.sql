CREATE TABLE IF NOT EXISTS tokens(
	token bytea PRIMARY KEY,
	email text REFERENCES users(email) ON DELETE CASCADE,
	created_at TIMESTAMP DEFAULT NOW()
);
