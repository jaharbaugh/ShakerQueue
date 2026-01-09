-- +goose Up
CREATE TYPE build_type AS ENUM ('stirred', 'shaken', 'sour', 'collins');

CREATE TABLE cocktail_recipes(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT UNIQUE NOT NULL,
    ingredients JSONB  NOT NULL,
    build build_type NOT NULL DEFAULT 'stirred' 
);


-- +goose Down
DROP TABLE cocktail_recipes;
DROP TYPE build_type;