-- +goose Up
CREATE TYPE user_role AS ENUM ('customer', 'employee', 'admin');

CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    username TEXT UNIQUE NOT NULL,
    email TEXT  UNIQUE NOT NULL,
    role user_role NOT NULL DEFAULT 'customer',
    hashed_password TEXT NOT NULL DEFAULT 'unset'
);


-- +goose Down
DROP TABLE users;
DROP TYPE user_role;