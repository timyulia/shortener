-- +goose Up
-- +goose StatementBegin
CREATE TABLE link
(
    id serial not null unique,
    short varchar(10) not null unique,
    long varchar(255) not null unique
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE link;
-- +goose StatementEnd
