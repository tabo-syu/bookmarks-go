package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
	"github.com/tabo-syu/bookmarks/sqlc"
	"github.com/tabo-syu/bookmarks/usecases"
)

type tagBookmarksGateway struct {
	db sqlc.Querier
}

func NewTagBookmarksGateway(sqlc sqlc.Querier) usecases.TagBookmarksRepository {
	return &tagBookmarksGateway{sqlc}
}

func (g *tagBookmarksGateway) List(ctx context.Context, id *uuid.UUID) ([]*domain.Bookmark, error) {
	records, err := g.db.FindBookmarksByTag(ctx, *id)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	bookmarks := make([]*domain.Bookmark, len(records))
	for i, record := range records {
		tag := &domain.Bookmark{
			ID:          record.ID,
			Url:         record.Url,
			Title:       record.Title,
			Description: record.Description,
			CreatedAt:   record.CreatedAt,
			UpdatedAt:   record.UpdatedAt,
		}
		bookmarks[i] = tag
	}

	return bookmarks, nil
}
