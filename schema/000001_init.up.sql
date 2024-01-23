CREATE TABLE users 
(
    id serial not null unique,
    username varchar(255) not null
);