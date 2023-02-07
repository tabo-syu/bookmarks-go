package usecases

import (
	"github.com/gin-gonic/gin"
	"github.com/tabo-syu/bookmarks/domain"
)

type BookmarksUsecase struct {
	Bookmarks      BookmarksRepository
	ResponseWriter WebAPIOutput
}

func NewBookmarksUsecase(bookmarks BookmarksRepository, writer WebAPIOutput) *BookmarksUsecase {
	return &BookmarksUsecase{bookmarks, writer}
}

func (s *BookmarksUsecase) List(g *gin.Context) {
	bookmarks, err := s.Bookmarks.List(g.Request.Context())
	s.ResponseWriter.Read(g, bookmarks, err)
}

func (s *BookmarksUsecase) Create(g *gin.Context, req *domain.BookmarkInput) {
	bookmark, err := s.Bookmarks.Create(g.Request.Context(), req)
	s.ResponseWriter.Create(g, bookmark, err)
}
