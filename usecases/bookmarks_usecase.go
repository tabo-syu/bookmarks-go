package usecases

import (
	"context"

	"github.com/tabo-syu/bookmarks/domain"
)

type BookmarksUsecase struct {
	bookmarks BookmarksRepository
}

func NewBookmarksUsecase(bookmarks BookmarksRepository) *BookmarksUsecase {
	return &BookmarksUsecase{bookmarks}
}

func (s *BookmarksUsecase) List(ctx context.Context) ([]*domain.Bookmark, error) {
	return s.bookmarks.List(ctx)
}

type BookmarkCreateRequest struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (s *BookmarksUsecase) Create(ctx context.Context, req *BookmarkCreateRequest) (*domain.Bookmark, error) {
	bookmark, err := domain.NewBookmark(req.Url, req.Title, req.Description)
	if err != nil {
		return nil, err
	}

	return s.bookmarks.Create(ctx, bookmark)
}
