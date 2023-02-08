package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
)

type BookmarksRepository interface {
	Get(context.Context, *uuid.UUID) (*domain.Bookmark, error)
	List(context.Context) ([]*domain.Bookmark, error)
	Create(context.Context, *domain.Bookmark) (*domain.Bookmark, error)
	Update(context.Context, *domain.Bookmark) (*domain.Bookmark, error)
	Delete(context.Context, *domain.Bookmark) error
}
