-- +goose Up
create table feed_follow (
    id uuid primary key ,
    feed_id uuid not null references feeds(id) on delete cascade ,
    user_id uuid not null references users(id) on delete  cascade ,
    unique(user_id,feed_id) ,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
drop table feed_follow;