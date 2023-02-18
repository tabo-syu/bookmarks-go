package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
)

type TagsUsecase struct {
	tags TagsRepository
}

func NewTagsUsecase(tags TagsRepository) *TagsUsecase {
	return &TagsUsecase{tags}
}

type TagGetRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (u *TagsUsecase) Get(ctx context.Context, req *TagGetRequest) (*domain.Tag, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	return u.tags.Get(ctx, &uuid)
}

func (u *TagsUsecase) List(ctx context.Context) ([]*domain.Tag, error) {
	return u.tags.List(ctx)
}

type TagCreateRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required,hexcolor"`
}

func (u *TagsUsecase) Create(ctx context.Context, req *TagCreateRequest) (*domain.Tag, error) {
	tag, err := domain.NewTag(req.Name, req.Color)
	if err != nil {
		return nil, err
	}

	return u.tags.Create(ctx, tag)
}

type TagUpdateURIRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type TagUpdateJSONRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required,hexcolor"`
}

type TagUpdateRequest struct {
	TagUpdateURIRequest
	TagUpdateJSONRequest
}

func (u *TagsUsecase) Update(ctx context.Context, req *TagUpdateRequest) (*domain.Tag, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	before, err := u.tags.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	after, err := domain.NewTag(req.Name, req.Color)
	if err != nil {
		return nil, err
	}

	after.ID = before.ID

	return u.tags.Update(ctx, after)
}

type TagDeleteRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (u *TagsUsecase) Delete(ctx context.Context, req *TagDeleteRequest) (*domain.Tag, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	tag, err := u.tags.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	if err := u.tags.Delete(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}
