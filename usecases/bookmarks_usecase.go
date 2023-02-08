package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/tabo-syu/bookmarks/domain"
)

type BookmarksUsecase struct {
	bookmarks BookmarksRepository
}

func NewBookmarksUsecase(bookmarks BookmarksRepository) *BookmarksUsecase {
	return &BookmarksUsecase{bookmarks}
}

type BookmarkGetRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (s *BookmarksUsecase) Get(ctx context.Context, req *BookmarkGetRequest) (*domain.Bookmark, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	return s.bookmarks.Get(ctx, &uuid)
}

func (s *BookmarksUsecase) List(ctx context.Context) ([]*domain.Bookmark, error) {
	return s.bookmarks.List(ctx)
}

type BookmarkCreateRequest struct {
	Url         string `json:"url" binding:"required,url"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (s *BookmarksUsecase) Create(ctx context.Context, req *BookmarkCreateRequest) (*domain.Bookmark, error) {
	bookmark, err := domain.NewBookmark(req.Url, req.Title, req.Description)
	if err != nil {
		return nil, err
	}

	return s.bookmarks.Create(ctx, bookmark)
}

type BookmarkDeleteRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (s *BookmarksUsecase) Delete(ctx context.Context, req *BookmarkDeleteRequest) (*domain.Bookmark, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	bookmark, err := s.bookmarks.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	if err := s.bookmarks.Delete(ctx, bookmark); err != nil {
		return nil, err
	}

	return bookmark, nil
}
