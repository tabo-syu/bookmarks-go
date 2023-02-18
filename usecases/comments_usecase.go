package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
)

type CommentsUsecase struct {
	comments CommentsRepository
}

func NewCommentsUsecase(comments CommentsRepository) *CommentsUsecase {
	return &CommentsUsecase{comments}
}

type CommentGetRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (u *CommentsUsecase) Get(ctx context.Context, req *CommentGetRequest) (*domain.Comment, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	return u.comments.Get(ctx, &uuid)
}

type CommentsListRequest struct {
	BookmarkId string `uri:"bookmark_id" binding:"required,uuid"`
}

func (u *CommentsUsecase) List(ctx context.Context, req *CommentsListRequest) ([]*domain.Comment, error) {
	uuid, err := uuid.Parse(req.BookmarkId)
	if err != nil {
		return nil, err
	}

	return u.comments.List(ctx, &uuid)
}

type CommentCreateRequest struct {
	CommentCreateURIRequest
	CommentCreateJSONRequest
}

type CommentCreateURIRequest struct {
	BookmarkID string `uri:"bookmark_id" binding:"required,uuid"`
}

type CommentCreateJSONRequest struct {
	Body string `json:"body" binding:"required"`
}

func (u *CommentsUsecase) Create(ctx context.Context, req *CommentCreateRequest) (*domain.Comment, error) {
	uuid, err := uuid.Parse(req.BookmarkID)
	if err != nil {
		return nil, err
	}

	comment, err := domain.NewComment(&uuid, req.Body)
	if err != nil {
		return nil, err
	}

	return u.comments.Create(ctx, comment)
}

type CommentDeleteRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (u *CommentsUsecase) Delete(ctx context.Context, req *CommentDeleteRequest) (*domain.Comment, error) {
	uuid, err := uuid.Parse(req.ID)
	if err != nil {
		return nil, err
	}

	comment, err := u.comments.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	if err := u.comments.Delete(ctx, comment); err != nil {
		return nil, err
	}

	return comment, nil
}
