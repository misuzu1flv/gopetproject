package postgresql

import (
	"context"
	"errors"
	"homework-8/internal/pkg/db"
	"homework-8/internal/pkg/repository"
)

type PostRepo struct {
	db db.DBops
}

func NewPostRepo(db db.DBops) *PostRepo {
	return &PostRepo{db: db}
}

func (pr *PostRepo) GetById(ctx context.Context, id int64) (*repository.Post, error) {
	res := &repository.Post{}
	err := pr.db.Get(ctx, res, "SELECT id, body, created_at FROM posts WHERE id=$1", id)
	if err != nil {
		return nil, repository.ErrUknownId
	}
	return res, nil
}

func (pr *PostRepo) Add(ctx context.Context, post *repository.Post) error {
	_, err := pr.GetById(ctx, post.Id)
	if err == nil {
		return repository.ErrObjectExists
	}
	if !errors.Is(err, repository.ErrUknownId) {
		return err
	}
	_, err = pr.db.Exec(ctx, "INSERT INTO POSTS (id, body) VALUES ($1, $2)", post.Id, post.Body)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRepo) Delete(ctx context.Context, id int64) error {
	_, err := pr.GetById(ctx, id)
	if err != nil {
		return err
	}
	_, err = pr.db.Exec(ctx, "DELETE FROM POSTS WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRepo) Update(ctx context.Context, post *repository.Post) error {
	_, err := pr.GetById(ctx, post.Id)
	if err != nil {
		return err
	}
	_, err = pr.db.Exec(ctx, "UPDATE POSTS SET body=$1 WHERE id=$2", post.Body, post.Id)
	if err != nil {
		return err
	}
	return nil
}
