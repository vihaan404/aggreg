-- +goose Up
alter table users
add column API_KEY varchar(64) unique not null default encode(sha256(random()::text::bytea), 'hex') ;



-- +goose Down
alter table users
drop column API_KEY