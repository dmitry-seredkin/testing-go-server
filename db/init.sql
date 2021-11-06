CREATE EXTENSION pgcrypto;

CREATE EXTENSION "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4 (),
    password VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    email VARCHAR,
    phone VARCHAR,
    PRIMARY KEY (id)
);

