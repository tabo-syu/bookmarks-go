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

func (s *TagsUsecase) Get(ctx context.Context, req *TagGetRequest) (*domain.Tag, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	return s.tags.Get(ctx, &uuid)
}

func (s *TagsUsecase) List(ctx context.Context) ([]*domain.Tag, error) {
	return s.tags.List(ctx)
}

type TagCreateRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required,hexcolor"`
}

func (s *TagsUsecase) Create(ctx context.Context, req *TagCreateRequest) (*domain.Tag, error) {
	tag, err := domain.NewTag(req.Name, req.Color)
	if err != nil {
		return nil, err
	}

	return s.tags.Create(ctx, tag)
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

func (s *TagsUsecase) Update(ctx context.Context, req *TagUpdateRequest) (*domain.Tag, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	before, err := s.tags.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	after, err := domain.NewTag(req.Name, req.Color)
	if err != nil {
		return nil, err
	}

	after.ID = before.ID

	return s.tags.Update(ctx, after)
}

type TagDeleteRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (s *TagsUsecase) Delete(ctx context.Context, req *TagDeleteRequest) (*domain.Tag, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	tag, err := s.tags.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	if err := s.tags.Delete(ctx, tag); err != nil {
		return nil, err
	}

	return tag, nil
}
