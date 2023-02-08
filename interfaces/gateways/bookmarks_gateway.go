package gateways

import (
	"context"

	"github.com/tabo-syu/bookmarks/domain"
	"github.com/tabo-syu/bookmarks/sqlc"
	"github.com/tabo-syu/bookmarks/usecases"
)

type bookmarksGateway struct {
	db sqlc.Querier
}

func NewBookmarksGateway(sqlc sqlc.Querier) usecases.BookmarksRepository {
	return &bookmarksGateway{sqlc}
}

func (r *bookmarksGateway) List(ctx context.Context) ([]*domain.Bookmark, error) {
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

func (r *bookmarksGateway) Create(ctx context.Context, req *domain.BookmarkCreateRequest) (*domain.Bookmark, error) {
	bookmark, err := domain.NewBookmark(req.Url, req.Title, req.Description)
	if err != nil {
		return nil, NewValidationError("Bookmark", err)
	}

	record, err := r.db.CreateBookmark(ctx, sqlc.CreateBookmarkParams{
		Url:         req.Url,
		Title:       req.Title,
		Description: req.Description,
	})
	if err != nil {
		return nil, NewPersistenceError(err)
	}

	bookmark.ID = record.ID
	bookmark.CreatedAt = record.CreatedAt
	bookmark.UpdatedAt = record.UpdatedAt

	return bookmark, nil
}
