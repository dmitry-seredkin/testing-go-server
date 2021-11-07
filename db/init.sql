CREATE EXTENSION pgcrypto;

CREATE EXTENSION "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4 (),
    login VARCHAR UNIQUE,
    password VARCHAR NOT NULL,
    name VARCHAR,
    email VARCHAR,
    phone VARCHAR,
    PRIMARY KEY (id)
);

