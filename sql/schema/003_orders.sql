-- +goose up

CREATE TYPE status AS ENUM ('pending', 'in_progress', 'complete', 'failed');

CREATE TABLE orders(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    recipe_id UUID NOT NULL REFERENCES cocktail_recipes(id),
    status status NOT NULL DEFAULT 'pending',
    started_at TIMESTAMP,
    completed_at TIMESTAMP,
    error_message TEXT
);


-- +goose down
DROP TABLE orders;
DROP TYPE status;