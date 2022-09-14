package main

import (
	"context"
	"fmt"
	"testcrud/database"
	"testcrud/entyty"
	"testcrud/repository"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsert(t *testing.T) {
	commentrepo := repository.NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment := entyty.Comment{
		Name:    "Repo Insert 01",
		Email:   "repo@gmail.com",
		Comment: "Repository Test",
	}

	result, err := commentrepo.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByID(t *testing.T) {
	commentrepo := repository.NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment, err := commentrepo.FindById(ctx, 16)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentrepo := repository.NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comments, err := commentrepo.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
