package gateways

import (
	"context"
	"time"

	"github.com/google/uuid"
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

func (g *bookmarksGateway) Get(ctx context.Context, id *uuid.UUID) (*domain.Bookmark, error) {
	record, err := g.db.GetBookmark(ctx, *id)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	bookmark := &domain.Bookmark{
		ID:          record.ID,
		Url:         record.Url,
		Title:       record.Title,
		Description: record.Description,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}

	return bookmark, err
}

func (g *bookmarksGateway) List(ctx context.Context) ([]*domain.Bookmark, error) {
	records, err := g.db.ListBookmarks(ctx)
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

func (g *bookmarksGateway) Create(ctx context.Context, bookmark *domain.Bookmark) (*domain.Bookmark, error) {
	record, err := g.db.CreateBookmark(ctx, sqlc.CreateBookmarkParams{
		Url:         bookmark.Url,
		Title:       bookmark.Title,
		Description: bookmark.Description,
	})
	if err != nil {
		return nil, NewPersistenceError(err)
	}

	bookmark.ID = record.ID
	bookmark.CreatedAt = record.CreatedAt
	bookmark.UpdatedAt = record.UpdatedAt

	return bookmark, nil
}

func (g *bookmarksGateway) Update(ctx context.Context, bookmark *domain.Bookmark) (*domain.Bookmark, error) {
	record, err := g.db.UpdateBookmark(ctx, sqlc.UpdateBookmarkParams{
		ID:          bookmark.ID,
		Url:         bookmark.Url,
		Title:       bookmark.Title,
		Description: bookmark.Description,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		return nil, NewPersistenceError(err)
	}

	bookmark.CreatedAt = record.CreatedAt
	bookmark.UpdatedAt = record.UpdatedAt

	return bookmark, err
}

func (g *bookmarksGateway) Delete(ctx context.Context, bookmark *domain.Bookmark) error {
	err := g.db.DeleteBookmark(ctx, bookmark.ID)
	if err != nil {
		return NewPersistenceError(err)
	}

	return nil
}
