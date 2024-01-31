CREATE TABLE users 
(
    id serial not null unique,
    username varchar(255) not null,
    address varchar(255) not null,
    pubkey varchar(255) not null
);

CREATE TABLE pair 
(
    id serial not null unique,
    id_user_initiator int references users(id) not null,
    id_user_partner int references users(id) not null
);

CREATE TABLE message 
(
    id serial not null unique,
    id_pair int references pair(id),
    encrypted_message varchar(255) not null,
    sent_at timestamp
);