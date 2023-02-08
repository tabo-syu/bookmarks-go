package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/domain"
)

type BookmarksUsecase struct {
	bookmarks      BookmarksRepository
	responseWriter WebAPIOutput
}

func NewBookmarksUsecase(bookmarks BookmarksRepository, writer WebAPIOutput) *BookmarksUsecase {
	return &BookmarksUsecase{bookmarks, writer}
}

func (s *BookmarksUsecase) List(g *gin.Context) {
	bookmarks, err := s.bookmarks.List(g.Request.Context())
	s.responseWriter.Read(g, bookmarks, err)
}

func (s *BookmarksUsecase) Create(g *gin.Context, req *domain.BookmarkCreateRequest) {
	bookmark, err := s.bookmarks.Create(g.Request.Context(), req)
	s.responseWriter.Create(g, bookmark, err)
}
