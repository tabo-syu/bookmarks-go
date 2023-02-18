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

func (u *BookmarksUsecase) Get(ctx context.Context, req *BookmarkGetRequest) (*domain.Bookmark, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	return u.bookmarks.Get(ctx, &uuid)
}

func (u *BookmarksUsecase) List(ctx context.Context) ([]*domain.Bookmark, error) {
	return u.bookmarks.List(ctx)
}

type BookmarkCreateRequest struct {
	Url         string `json:"url" binding:"required,url"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (u *BookmarksUsecase) Create(ctx context.Context, req *BookmarkCreateRequest) (*domain.Bookmark, error) {
	bookmark, err := domain.NewBookmark(req.Url, req.Title, req.Description)
	if err != nil {
		return nil, err
	}

	return u.bookmarks.Create(ctx, bookmark)
}

type BookmarkUpdateURIRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

type BookmarkUpdateJSONRequest struct {
	Url         string `json:"url" binding:"required,url"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type BookmarkUpdateRequest struct {
	BookmarkUpdateURIRequest
	BookmarkUpdateJSONRequest
}

func (u *BookmarksUsecase) Update(ctx context.Context, req *BookmarkUpdateRequest) (*domain.Bookmark, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	before, err := u.bookmarks.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	after, err := domain.NewBookmark(req.Url, req.Title, req.Description)
	if err != nil {
		return nil, err
	}

	after.ID = before.ID

	return u.bookmarks.Update(ctx, after)
}

type BookmarkDeleteRequest struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (u *BookmarksUsecase) Delete(ctx context.Context, req *BookmarkDeleteRequest) (*domain.Bookmark, error) {
	uuid, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, err
	}

	bookmark, err := u.bookmarks.Get(ctx, &uuid)
	if err != nil {
		return nil, err
	}

	if err := u.bookmarks.Delete(ctx, bookmark); err != nil {
		return nil, err
	}

	return bookmark, nil
}
