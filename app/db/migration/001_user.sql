-- +goose Up
CREATE TABLE Users (
    id SERIAL PRIMARY KEY ,
    username VARCHAR UNIQUE NOT NULL,
    password_hash VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE Notes (
    note_id BIGSERIAL PRIMARY KEY ,
    user_id SERIAL,
    title VARCHAR NOT NULL,
    content TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- +goose Down
DROP TABLE users;
DROP TABLE notes;