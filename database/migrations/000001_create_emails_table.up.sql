CREATE TABLE IF NOT EXISTS emails (
    id SERIAL PRIMARY KEY,
    owner_email VARCHAR(255) NOT NULL,
    email_html TEXT NOT NULL
);