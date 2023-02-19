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

type BookmarkTagsRepository interface {
	List(ctx context.Context, bookmarkID *uuid.UUID) ([]*domain.Tag, error)
	Add(ctx context.Context, bookmarkID *uuid.UUID, tagID *uuid.UUID) error
	Remove(ctx context.Context, bookmarkID *uuid.UUID, tagID *uuid.UUID) error
}

type TagsRepository interface {
	Get(context.Context, *uuid.UUID) (*domain.Tag, error)
	List(context.Context) ([]*domain.Tag, error)
	Create(context.Context, *domain.Tag) (*domain.Tag, error)
	Update(context.Context, *domain.Tag) (*domain.Tag, error)
	Delete(context.Context, *domain.Tag) error
}

type CommentsRepository interface {
	Get(context.Context, *uuid.UUID) (*domain.Comment, error)
	List(context.Context, *uuid.UUID) ([]*domain.Comment, error)
	Create(context.Context, *domain.Comment) (*domain.Comment, error)
	Delete(context.Context, *domain.Comment) error
}
