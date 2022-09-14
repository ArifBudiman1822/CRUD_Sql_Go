package repository

import (
	"context"
	"testcrud/entyty"
)

type CommentRepository interface {
	Insert(ctx context.Context, comment entyty.Comment) (entyty.Comment, error)
	FindById(ctx context.Context, id int32) (entyty.Comment, error)
	FindAll(ctx context.Context) ([]entyty.Comment, error)
}
