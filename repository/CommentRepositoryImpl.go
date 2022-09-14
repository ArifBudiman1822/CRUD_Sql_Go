package repository

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"testcrud/entyty"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &CommentRepositoryImpl{DB: db}
}

func (repo *CommentRepositoryImpl) Insert(ctx context.Context, comment entyty.Comment) (entyty.Comment, error) {
	query := "INSERT INTO comment(name,email,comment)VALUES(?,?,?)"

	result, err := repo.DB.ExecContext(ctx, query, comment.Name, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repo *CommentRepositoryImpl) FindById(ctx context.Context, id int32) (entyty.Comment, error) {
	query := "SELECT id,name,email,comment from comment where id = ?"
	comment := entyty.Comment{}
	rows, err := repo.DB.QueryContext(ctx, query, id)
	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Name, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("Id :" + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repo *CommentRepositoryImpl) FindAll(ctx context.Context) ([]entyty.Comment, error) {
	query := "SELECT id,name,email,comment from comment"

	rows, err := repo.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []entyty.Comment
	for rows.Next() {
		comment := entyty.Comment{}
		rows.Scan(&comment.Id, &comment.Name, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil

}
