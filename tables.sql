CREATE TABLE users
(
    id       serial primary key,
    name     varchar,
    username varchar not null unique,
    password varchar not null,
    role     varchar not null
);

CREATE TABLE items
(
    id    serial primary key,
    name  varchar not null,
    info  varchar,
    price bigint
);