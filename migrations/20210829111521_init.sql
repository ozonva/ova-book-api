-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS books (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL CHECK (user_id >= 1),
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    isbn10 CHAR(10) NOT NULL,
    isbn13 CHAR(13) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS books;
-- +goose StatementEnd
