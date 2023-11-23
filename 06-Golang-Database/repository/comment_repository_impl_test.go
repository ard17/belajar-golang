package repository

import (
	"context"
	"fmt"
	golangdatabase "golang-database"
	"golang-database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "test@repo.com",
		Comment: "Test insert comment",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)
}

func TestFindCommentById(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	ctx := context.Background()
	id := 35

	result, err := commentRepository.FindById(ctx, int32(id))
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindAllComment(t *testing.T) {
	commentRepository := NewCommentRepository(golangdatabase.GetConnection())

	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
