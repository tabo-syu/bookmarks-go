package usecases

import (
	"context"

	"github.com/tabo-syu/bookmarks/domain"
)

type BookmarksRepository interface {
	List(context.Context) ([]*domain.Bookmark, error)
	Create(context.Context, *domain.BookmarkCreateRequest) (*domain.Bookmark, error)
}
