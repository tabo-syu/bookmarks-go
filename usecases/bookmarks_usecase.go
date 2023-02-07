package usecases

import (
	"context"
)

type BookmarksUsecase struct {
	bookmarks BookmarksRepository
}

func NewBookmarksUsecase(bookmarks BookmarksRepository) *BookmarksUsecase {
	return &BookmarksUsecase{bookmarks}
}

func (s *BookmarksUsecase) List(ctx context.Context) {
	_, err := s.bookmarks.List(ctx)
	if err != nil {
		return
	}
}
