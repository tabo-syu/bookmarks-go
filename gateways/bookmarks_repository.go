package gateways

import (
	"context"

	"github.com/tabo-syu/bookmarks/infrastructures/sqlc"
)

type BookmarksRepository interface {
	List(context.Context) ([]sqlc.Bookmark, error)
}

type bookmarksRepository struct {
	db *sqlc.Queries
}

func NewBookmarksRepository(sqlc *sqlc.Queries) BookmarksRepository {
	return &bookmarksRepository{sqlc}
}

func (r *bookmarksRepository) List(ctx context.Context) ([]sqlc.Bookmark, error) {
	return r.db.ListBookmarks(ctx)
}
