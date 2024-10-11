-- +goose Up
CREATE TABLE Users (
    user_id INT PRIMARY KEY ,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE Notes (
    note_id INT PRIMARY KEY ,
    user_id INT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY (user_id) REFERENCES Users(user_id)
);

-- +goose Down
DROP TABLE users;
DROP TABLE notes;