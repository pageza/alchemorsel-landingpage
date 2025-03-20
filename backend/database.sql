-- cursor--Initial commit: SQL script to create the subscribers table in PostgreSQL.
CREATE TABLE IF NOT EXISTS subscribers (
    id SERIAL PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
); 