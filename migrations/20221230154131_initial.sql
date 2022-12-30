-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       email text NOT NULL,
                       username character varying(30) NOT NULL,
                       firstname text NOT NULL,
                       lastname text NOT NULL,
                       age integer NOT NULL,
                       birthdate date NOT NULL,
                       password text NOT NULL,
                       created_at date NOT NULL,
                       updated_at date NOT NULL,
                       deleted_at date
);

CREATE UNIQUE INDEX users_email_unique ON users(email text_ops);
CREATE UNIQUE INDEX users_username_unique ON users(username text_ops);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX users_email_unique;
DROP INDEX users_username_unique;
DROP TABLE users;
-- +goose StatementEnd
