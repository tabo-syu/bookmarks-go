package gateways

import (
	"context"

	"github.com/tabo-syu/bookmarks/domain"
	"github.com/tabo-syu/bookmarks/sqlc"
)

type BookmarksRepository interface {
	List(context.Context) ([]*domain.Bookmark, error)
	Create(context.Context, *domain.BookmarkInput) (*domain.Bookmark, error)
}

type bookmarksRepository struct {
	db sqlc.Querier
}

func NewBookmarksRepository(sqlc sqlc.Querier) BookmarksRepository {
	return &bookmarksRepository{sqlc}
}

func (r *bookmarksRepository) List(ctx context.Context) ([]*domain.Bookmark, error) {
	records, err := r.db.ListBookmarks(ctx)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	bookmarks := []*domain.Bookmark{}
	for _, record := range records {
		bookmark := &domain.Bookmark{
			ID:          record.ID,
			Url:         record.Url,
			Title:       record.Title,
			Description: record.Description,
			CreatedAt:   record.CreatedAt,
			UpdatedAt:   record.UpdatedAt,
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

func (r *bookmarksRepository) Create(ctx context.Context, param *domain.BookmarkInput) (*domain.Bookmark, error) {
	bookmark, err := domain.NewBookmark(param)
	if err != nil {
		return nil, NewValidationError("Bookmark", err)
	}

	record, err := r.db.CreateBookmark(ctx, sqlc.CreateBookmarkParams{
		Url:         param.Url,
		Title:       param.Title,
		Description: param.Description,
	})
	if err != nil {
		return nil, NewPersistenceError(err)
	}

	bookmark.ID = record.ID
	bookmark.CreatedAt = record.CreatedAt
	bookmark.UpdatedAt = record.UpdatedAt

	return bookmark, nil
}

var _ BookmarksRepository = (*bookmarksRepository)(nil)
