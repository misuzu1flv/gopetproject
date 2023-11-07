package postgresql

import (
	"context"
	"homework-8/internal/pkg/db"
	"homework-8/internal/pkg/repository"
	"log"
)

type CommentRepo struct {
	db db.DBops
}

func NewCommentRepo(db db.DBops) *CommentRepo {
	return &CommentRepo{db: db}
}

func (cr *CommentRepo) Add(ctx context.Context, comment *repository.Comment) error {
	_, err := cr.GetById(ctx, comment.Id)
	if err != repository.ErrUknownId {
		return repository.ErrObjectExists
	}
	_, err = cr.db.Exec(ctx, "INSERT INTO COMMENTS (id, post_id, body) VALUES ($1, $2, $3)", comment.Id, comment.PostId, comment.Body)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepo) DeleteByPostId(ctx context.Context, post_id int64) error {
	_, err := cr.db.Exec(ctx, "DELETE FROM COMMENTS WHERE post_id=$1", post_id)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepo) GetByPostId(ctx context.Context, post_id int64) ([]*repository.Comment, error) {
	var ammount int64
	err := cr.db.Get(ctx, &ammount, "SELECT COUNT(*) FROM comments WHERE post_id=$1", post_id)
	if err != nil {
		return nil, err
	}
	res := make([]*repository.Comment, ammount)
	err = cr.db.Select(ctx, &res, "SELECT id, post_id, body, created_at FROM comments WHERE post_id=$1", post_id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cr *CommentRepo) GetById(ctx context.Context, id int64) (*repository.Comment, error) {
	res := &repository.Comment{}
	err := cr.db.Get(ctx, res, "SELECT id, post_id, body, created_at FROM comments WHERE id=$1", id)
	log.Println(err)
	if err != nil {
		return nil, repository.ErrUknownId
	}
	return res, nil
}
