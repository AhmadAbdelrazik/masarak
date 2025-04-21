CREATE TABLE IF NOT EXISTS businesses(
	id SERIAL PRIMARY KEY,
	name TEXT NOT NULL,
	business_email TEXT NOT NULL,
	description TEXT NOT NULL,
	image_url TEXT NOT NULL,
	owner_email TEXT NOT NULL REFERENCES users(email),
	version INT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS employees(
	id SERIAL PRIMARY KEY,
	email TEXT NOT NULL UNIQUE,
	business_id INT REFERENCES businesses(id),
	version INT DEFAULT 1
);

CREATE TABLE IF NOT EXISTS jobs(
	id SERIAL PRIMARY KEY,
	business_id INT NOT NULL REFERENCES businesses(id),
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	work_location TEXT NOT NULL,
	work_time TEXT NOT NULL,
	skills TEXT[10] NOT NULL,
	min_years_of_experience INT NOT NULL DEFAULT 0,
	max_years_of_experience INT NOT NULL DEFAULT 40,
	min_expected_salary NUMERIC NOT NULL DEFAULT 0,
	max_expected_salary NUMERIC NOT NULL DEFAULT 1000000,
	expected_salary_currency TEXT NOT NULL DEFAULT 'EGP',
	status TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	version int DEFAULT 1, 
	CHECK (min_years_of_experience >= 0),
	CHECK (max_years_of_experience >= min_years_of_experience),
	CHECK (min_expected_salary >= 0),
	CHECK (max_expected_salary >= min_expected_salary)
);

CREATE TABLE IF NOT EXISTS applications(
	id SERIAL PRIMARY KEY,
	business_id INT NOT NULL REFERENCES businesses(id),
	job_id INT NOT NULL REFERENCES jobs(id),
	status TEXT NOT NULL,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	title TEXT NOT NULL,
	years_of_experience INT NOT NULL DEFAULT 0,
	hourly_rate_amount NUMERIC NOT NULL DEFAULT 0,
	hourly_rate_currency TEXT NOT NULL DEFAULT 'EGP',
	profile_url TEXT NOT NULL,
	resume_url TEXT NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
	version int DEFAULT 1, 
	CHECK (years_of_experience >= 0)
);
