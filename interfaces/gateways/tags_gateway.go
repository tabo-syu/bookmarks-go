package gateways

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
	"github.com/tabo-syu/bookmarks/sqlc"
	"github.com/tabo-syu/bookmarks/usecases"
)

type tagsGateway struct {
	db sqlc.Querier
}

func NewTagsGateway(sqlc sqlc.Querier) usecases.TagsRepository {
	return &tagsGateway{sqlc}
}

func (r *tagsGateway) Get(ctx context.Context, id *uuid.UUID) (*domain.Tag, error) {
	record, err := r.db.GetTag(ctx, *id)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	tag := &domain.Tag{
		ID:        record.ID,
		Name:      record.Name,
		Color:     record.Color,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}

	return tag, err
}

func (r *tagsGateway) List(ctx context.Context) ([]*domain.Tag, error) {
	records, err := r.db.ListTags(ctx)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	tags := []*domain.Tag{}
	for _, record := range records {
		tag := &domain.Tag{
			ID:        record.ID,
			Name:      record.Name,
			Color:     record.Color,
			CreatedAt: record.CreatedAt,
			UpdatedAt: record.UpdatedAt,
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *tagsGateway) Create(ctx context.Context, tag *domain.Tag) (*domain.Tag, error) {
	record, err := r.db.CreateTag(ctx, sqlc.CreateTagParams{
		Name:  tag.Name,
		Color: tag.Color,
	})
	if err != nil {
		return nil, NewPersistenceError(err)
	}

	tag.ID = record.ID
	tag.CreatedAt = record.CreatedAt
	tag.UpdatedAt = record.UpdatedAt

	return tag, nil
}

func (r *tagsGateway) Update(ctx context.Context, tag *domain.Tag) (*domain.Tag, error) {
	record, err := r.db.UpdateTag(ctx, sqlc.UpdateTagParams{
		ID:        tag.ID,
		Name:      tag.Name,
		Color:     tag.Color,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, NewPersistenceError(err)
	}

	tag.CreatedAt = record.CreatedAt
	tag.UpdatedAt = record.UpdatedAt

	return tag, err
}

func (r *tagsGateway) Delete(ctx context.Context, tag *domain.Tag) error {
	err := r.db.DeleteTag(ctx, tag.ID)
	if err != nil {
		return NewPersistenceError(err)
	}

	return nil
}
