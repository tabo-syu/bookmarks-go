package usecases

import (
	"context"

	"github.com/gin-gonic/gin"
)

type BookmarksUsecase struct {
	bookmarks BookmarksRepository
	writer    WebAPIOutput
}

func NewBookmarksUsecase(bookmarks BookmarksRepository, writer WebAPIOutput) *BookmarksUsecase {
	return &BookmarksUsecase{bookmarks, writer}
}

func (s *BookmarksUsecase) List(ctx context.Context, g *gin.Context) {
	bookmarks, err := s.bookmarks.List(ctx)
	s.writer.Read(g, bookmarks, err)
}
