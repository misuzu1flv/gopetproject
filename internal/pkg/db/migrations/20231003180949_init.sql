-- +goose Up
-- +goose StatementBegin
create table posts(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    body text not null default '',
    created_at timestamp with time zone default now() not null
);

create table comments(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    post_id BIGINT NOT NULL,
    body text not null default '',
    created_at timestamp with time zone default now() not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table posts;
drop table comments;
-- +goose StatementEnd
