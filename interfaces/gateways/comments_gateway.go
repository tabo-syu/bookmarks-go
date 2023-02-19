package gateways

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
	"github.com/tabo-syu/bookmarks/sqlc"
	"github.com/tabo-syu/bookmarks/usecases"
)

type commentsGateway struct {
	db sqlc.Querier
}

func NewCommentsGateway(sqlc sqlc.Querier) usecases.CommentsRepository {
	return &commentsGateway{sqlc}
}

func (g *commentsGateway) Get(ctx context.Context, id *uuid.UUID) (*domain.Comment, error) {
	record, err := g.db.GetComment(ctx, *id)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	comment := &domain.Comment{
		ID:         record.ID,
		BookmarkID: record.BookmarkID,
		Body:       record.Body,
		CreatedAt:  record.CreatedAt,
		UpdatedAt:  record.UpdatedAt,
	}

	return comment, err
}

func (g *commentsGateway) List(ctx context.Context, id *uuid.UUID) ([]*domain.Comment, error) {
	records, err := g.db.ListComments(ctx, *id)
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	comments := make([]*domain.Comment, len(records))
	for i, record := range records {
		comment := &domain.Comment{
			ID:         record.ID,
			BookmarkID: record.BookmarkID,
			Body:       record.Body,
			CreatedAt:  record.CreatedAt,
			UpdatedAt:  record.UpdatedAt,
		}
		comments[i] = comment
	}

	return comments, nil
}

func (g *commentsGateway) Create(ctx context.Context, comment *domain.Comment) (*domain.Comment, error) {
	record, err := g.db.CreateComment(ctx, sqlc.CreateCommentParams{
		BookmarkID: comment.BookmarkID,
		Body:       comment.Body,
	})
	if err != nil {
		return nil, NewMissingEntityError(err)
	}

	comment.ID = record.ID
	comment.CreatedAt = record.CreatedAt
	comment.UpdatedAt = record.UpdatedAt

	return comment, err
}

func (g *commentsGateway) Delete(ctx context.Context, comment *domain.Comment) error {
	err := g.db.DeleteComment(ctx, comment.ID)
	if err != nil {
		return NewMissingEntityError(err)
	}

	return nil
}
