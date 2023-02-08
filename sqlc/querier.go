// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateBookmark(ctx context.Context, arg CreateBookmarkParams) (Bookmark, error)
	DeleteBookmark(ctx context.Context, id uuid.UUID) error
	FindBookmarks(ctx context.Context, ids []uuid.UUID) ([]Bookmark, error)
	FindBookmarksByTags(ctx context.Context, ids []uuid.UUID) ([]FindBookmarksByTagsRow, error)
	FindCommentsByBookmark(ctx context.Context, ids []uuid.UUID) ([]Comment, error)
	GetBookmark(ctx context.Context, id uuid.UUID) (Bookmark, error)
	ListBookmarks(ctx context.Context) ([]Bookmark, error)
	UpdateBookmark(ctx context.Context, arg UpdateBookmarkParams) (Bookmark, error)
}

var _ Querier = (*Queries)(nil)
