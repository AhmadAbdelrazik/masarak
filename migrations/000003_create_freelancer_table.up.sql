CREATE TABLE IF NOT EXISTS freelancer_profiles(
	email text PRIMARY KEY REFERENCES users(email),
	name text NOT NULL,
	title text NOT NULL,
	picture_url text NOT NULL,
	skills text[] NOT NULL,
	years_of_experience int NOT NULL,
	hourly_rate_currency VARCHAR(3) NOT NULL,
	hourly_rate_amount NUMERIC NOT NULL,
	resume_url text NOT NULL,
	version int DEFAULT 1
);
