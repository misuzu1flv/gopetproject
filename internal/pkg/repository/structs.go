package repository

import "time"

type Post struct {
	Id        int64     `db:"id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}

type Comment struct {
	Id        int64     `db:"id"`
	PostId    int64     `db:"post_id"`
	Body      string    `db:"body"`
	CreatedAt time.Time `db:"created_at"`
}
