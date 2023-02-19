package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
	"github.com/tabo-syu/bookmarks/sqlc"
	"github.com/tabo-syu/bookmarks/usecases"
)

type bookmarkTagsGateway struct {
	db sqlc.Querier
}

func NewBookmarkTagsGateway(sqlc sqlc.Querier) usecases.BookmarkTagsRepository {
	return &bookmarkTagsGateway{sqlc}
}

func (g *bookmarkTagsGateway) List(ctx context.Context, id *uuid.UUID) ([]*domain.Tag, error) {
	records, err := g.db.FindTagsByBookmark(ctx, *id)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	tags := make([]*domain.Tag, len(records))
	for i, record := range records {
		tag := &domain.Tag{
			ID:        record.ID,
			Name:      record.Name,
			Color:     record.Color,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		}
		tags[i] = tag
	}

	return tags, nil
}

func (g *bookmarkTagsGateway) Add(ctx context.Context, bookmarkID *uuid.UUID, tagID *uuid.UUID) error {
	err := g.db.AddTagToBookmark(ctx, sqlc.AddTagToBookmarkParams{
		BookmarkID: *bookmarkID,
		TagID:      *tagID,
	})
	if err != nil {
		return NewPersistenceError(err)
	}

	return nil
}

func (g *bookmarkTagsGateway) Remove(ctx context.Context, bookmarkID *uuid.UUID, tagID *uuid.UUID) error {
	err := g.db.RemoveTagFromBookmark(ctx, sqlc.RemoveTagFromBookmarkParams{
		BookmarkID: *bookmarkID,
		TagID:      *tagID,
	})
	if err != nil {
		return NewPersistenceError(err)
	}

	return nil
}
