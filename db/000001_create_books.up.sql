CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    name varchar(128) NOT NULL,
    author varchar(32) NOT NULL,
    genre varchar(32)
);