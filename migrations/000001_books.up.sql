CREATE TABLE IF NOT EXISTS books (
    id         UUID PRIMARY KEY,
    title      VARCHAR(200) NOT NULL,
    author     VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL default now()
);
