CREATE TABLE users (
  id bigserial NOT NULL primary key,
  email varchar not null unique,
  encrypted_password varchar not null
);
