CREATE TABLE IF NOT EXISTS users (
   id bigserial PRIMARY KEY,
   email VARCHAR (254) NOT NULL UNIQUE,
   password VARCHAR (300) NOT NULL
);