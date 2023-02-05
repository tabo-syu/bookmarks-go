package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/infrastructures"
	"github.com/tabo-syu/bookmarks/sqlc"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf(err.Error())
	}
}

func run() error {
	db, err := infrastructures.NewSQLClient()
	if err != nil {
		log.Fatalf(err.Error())
	}

	q := sqlc.New(db)
	ctx := context.Background()
	fmt.Println("--- ListBookmarks() ---")
	bookmarks, err := q.ListBookmarks(ctx)
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, bookmark := range bookmarks {
		fmt.Println("=======")
		fmt.Println(bookmark.ID)
		fmt.Println(bookmark.Url)
		fmt.Println(bookmark.Description)
	}

	id, err := uuid.Parse("92449b60-90ac-4161-9be1-200b07854a46")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println("")
	fmt.Println("--- ListBookmarks() ---")
	res, err := q.FindBookmarksByTags(ctx, []uuid.UUID{id})
	if err != nil {
		log.Fatalf(err.Error())
	}
	for _, bookmark := range res {
		fmt.Println("=======")
		fmt.Println(bookmark.ID)
		fmt.Println(bookmark.Title)
		fmt.Println(bookmark.Description)
		fmt.Println(bookmark.Url)
	}

	return nil
}
